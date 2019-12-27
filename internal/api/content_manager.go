package api

import (
	"context"
	"github.com/kubenext/kubefun/internal/event"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/internal/module"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/view/component"
	"strings"
	"time"

	"github.com/pkg/errors"
)

const (
	RequestSetContentPath = "setContentPath"
	RequestSetNamespace   = "setNamespace"
)

// ContentManagerOption is an option for configuring ContentManager.
type ContentManagerOption func(manager *ContentManager)

// ContentGenerateFunc is a function that generates content. It returns `rerun=true`
// if the action should be be immediately rerun.
type ContentGenerateFunc func(ctx context.Context, state kubefun.State) (component.ContentResponse, bool, error)

// WithContentGenerator configures the content generate function.
func WithContentGenerator(fn ContentGenerateFunc) ContentManagerOption {
	return func(manager *ContentManager) {
		manager.contentGenerateFunc = fn
	}
}

// WithContentGeneratorPoller configures the poller.
func WithContentGeneratorPoller(poller Poller) ContentManagerOption {
	return func(manager *ContentManager) {
		manager.poller = poller
	}
}

// ContentManager manages content for websockets.
type ContentManager struct {
	moduleManager       module.ManagerInterface
	logger              log.Logger
	contentGenerateFunc ContentGenerateFunc
	poller              Poller
	updateContentCh     chan struct{}
}

// NewContentManager creates an instance of ContentManager.
func NewContentManager(moduleManager module.ManagerInterface, logger log.Logger, options ...ContentManagerOption) *ContentManager {
	cm := &ContentManager{
		moduleManager:   moduleManager,
		logger:          logger,
		poller:          NewInterruptiblePoller("content"),
		updateContentCh: make(chan struct{}, 1),
	}
	cm.contentGenerateFunc = cm.generateContent

	for _, option := range options {
		option(cm)
	}

	return cm
}

var _ StateManager = (*ContentManager)(nil)

// Start starts the manager.
func (cm *ContentManager) Start(ctx context.Context, state kubefun.State, s KubefunClient) {
	defer func() {
		close(cm.updateContentCh)
	}()

	updateCancel := state.OnContentPathUpdate(func(contentPath string) {
		cm.updateContentCh <- struct{}{}
	})
	defer updateCancel()

	cm.poller.Run(ctx, cm.updateContentCh, cm.runUpdate(state, s), event.DefaultScheduleDelay)
}

func (cm *ContentManager) runUpdate(state kubefun.State, s KubefunClient) PollerFunc {
	return func(ctx context.Context) bool {
		contentPath := state.GetContentPath()
		if contentPath == "" {
			return false
		}

		contentResponse, _, err := cm.contentGenerateFunc(ctx, state)
		if err != nil {
			cm.logger.
				WithErr(err).
				With("content-path", contentPath).
				Errorf("generate content")
			return false
		}

		if ctx.Err() == nil {
			s.Send(CreateContentEvent(contentResponse, state.GetNamespace(), contentPath, state.GetQueryParams()))
		}

		return false
	}
}

func (cm *ContentManager) generateContent(ctx context.Context, state kubefun.State) (component.ContentResponse, bool, error) {
	contentPath := state.GetContentPath()
	logger := cm.logger.With("contentPath", contentPath)

	now := time.Now()
	defer func() {
		logger.With("elapsed", time.Since(now)).Debugf("generating content")
	}()

	m, ok := cm.moduleManager.ModuleForContentPath(contentPath)
	if !ok {
		return component.EmptyContentResponse, false, errors.Errorf("unable to find module for content path %q", contentPath)
	}
	modulePath := strings.TrimPrefix(contentPath, m.Name())
	options := module.ContentOptions{
		LabelSet: FiltersToLabelSet(state.GetFilters()),
	}
	contentResponse, err := m.Content(ctx, modulePath, options)
	if err != nil {
		if nfe, ok := err.(notFound); ok && nfe.NotFound() {
			logger.Debugf("path not found, redirecting to parent")
			state.SetContentPath(notFoundRedirectPath(contentPath))
			return component.EmptyContentResponse, true, nil
		} else {
			return component.EmptyContentResponse, false, errors.Wrap(err, "generate content")
		}
	}

	return contentResponse, false, nil
}

// Handlers returns a slice of client request handlers.
func (cm *ContentManager) Handlers() []kubefun.ClientRequestHandler {
	return []kubefun.ClientRequestHandler{
		{
			RequestType: RequestSetContentPath,
			Handler:     cm.SetContentPath,
		},
		{
			RequestType: RequestSetNamespace,
			Handler:     cm.SetNamespace,
		},
	}
}

// SetQueryParams sets the current query params.
func (cm *ContentManager) SetQueryParams(state kubefun.State, payload action.Payload) error {
	if params, ok := payload["params"].(map[string]interface{}); ok {
		// handle filters
		if filters, ok := params["filters"]; ok {
			list, err := FiltersFromQueryParams(filters)
			if err != nil {
				return errors.Wrap(err, "extract filters from query params")
			}
			state.SetFilters(list)
		}
	}

	return nil
}

// SetNamespace sets the current namespace.
func (cm *ContentManager) SetNamespace(state kubefun.State, payload action.Payload) error {
	namespace, err := payload.String("namespace")
	if err != nil {
		return errors.Wrap(err, "extract namespace from payload")
	}
	state.SetNamespace(namespace)
	return nil
}

// SetContentPath sets the current content path.
func (cm *ContentManager) SetContentPath(state kubefun.State, payload action.Payload) error {
	contentPath, err := payload.String("contentPath")
	if err != nil {
		return errors.Wrap(err, "extract contentPath from payload")
	}
	if err := cm.SetQueryParams(state, payload); err != nil {
		return errors.Wrap(err, "extract query params from payload")
	}

	state.SetContentPath(contentPath)
	return nil
}

type notFound interface {
	NotFound() bool
	Path() string
}

// CreateContentEvent creates a content event.
func CreateContentEvent(contentResponse component.ContentResponse, namespace, contentPath string, queryParams map[string][]string) kubefun.Event {
	return kubefun.Event{
		Type: kubefun.EventTypeContent,
		Data: map[string]interface{}{
			"content":     contentResponse,
			"namespace":   namespace,
			"contentPath": contentPath,
			"queryParams": queryParams,
		},
	}
}
