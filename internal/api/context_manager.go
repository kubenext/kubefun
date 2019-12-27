package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/kubenext/kubefun/internal/config"
	"github.com/kubenext/kubefun/internal/event"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/pkg/errors"
)

const (
	RequestSetContext = "setContext"
)

// ContextManagerOption is an option for configuring ContextManager.
type ContextManagerOption func(manager *ContextManager)

// ContextGenerateFunc is a function which generates a context event.
type ContextGenerateFunc func(ctx context.Context, state kubefun.State) (kubefun.Event, error)

// WithContextGenerator sets the context generator.
func WithContextGenerator(fn ContextGenerateFunc) ContextManagerOption {
	return func(manager *ContextManager) {
		manager.contextGenerateFunc = fn
	}
}

// WithContextGeneratorPoll generates the poller.
func WithContextGeneratorPoll(poller Poller) ContextManagerOption {
	return func(manager *ContextManager) {
		manager.poller = poller
	}
}

// ContextManager manages context.
type ContextManager struct {
	dashConfig          config.Dash
	contextGenerateFunc ContextGenerateFunc
	poller              Poller
}

var _ StateManager = (*ContextManager)(nil)

// NewContextManager creates an instances of ContextManager.
func NewContextManager(dashConfig config.Dash, options ...ContextManagerOption) *ContextManager {
	cm := &ContextManager{
		dashConfig: dashConfig,
		poller:     NewInterruptiblePoller("context"),
	}

	cm.contextGenerateFunc = cm.generateContexts

	for _, option := range options {
		option(cm)
	}

	return cm
}

// Handlers returns a slice of handlers.
func (c *ContextManager) Handlers() []kubefun.ClientRequestHandler {
	return []kubefun.ClientRequestHandler{
		{
			RequestType: RequestSetContext,
			Handler:     c.SetContext,
		},
	}
}

// SetContext sets the current context.
func (c *ContextManager) SetContext(state kubefun.State, payload action.Payload) error {
	requestedContext, err := payload.String("requestedContext")
	if err != nil {
		return errors.Wrap(err, "extract requested context from payload")
	}
	state.SetContext(requestedContext)
	return nil
}

// Start starts the manager.
func (c *ContextManager) Start(ctx context.Context, state kubefun.State, s KubefunClient) {
	c.poller.Run(ctx, nil, c.runUpdate(state, s), event.DefaultScheduleDelay)
}

func (c *ContextManager) runUpdate(state kubefun.State, s KubefunClient) PollerFunc {
	var previous []byte

	logger := c.dashConfig.Logger()
	return func(ctx context.Context) bool {
		ev, err := c.contextGenerateFunc(ctx, state)
		if err != nil {
			logger.WithErr(err).Errorf("generate contexts")
		}

		if ctx.Err() == nil {
			cur, err := json.Marshal(ev)
			if err != nil {
				logger.WithErr(err).Errorf("unable to marshal context")
				return false
			}

			if bytes.Compare(previous, cur) != 0 {
				previous = cur
				s.Send(ev)
			}
		}

		return false
	}
}

func (c *ContextManager) generateContexts(ctx context.Context, state kubefun.State) (kubefun.Event, error) {
	generator, err := c.initGenerator(state)
	if err != nil {
		return kubefun.Event{}, err
	}
	return generator.Event(ctx)
}

func (c *ContextManager) initGenerator(state kubefun.State) (*event.ContextsGenerator, error) {
	return event.NewContextsGenerator(c.dashConfig), nil
}
