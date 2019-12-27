package resourceviewer

import (
	"context"
	"github.com/kubenext/kubefun/internal/config"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/internal/objectvisitor"
	"github.com/kubenext/kubefun/internal/queryer"
	"github.com/kubenext/kubefun/internal/util/kubernetes"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"time"
)

const (
	visitMaxDuration = 3 * time.Second
)

// ViewerOpt is an option for ResourceViewer.
type ViewerOpt func(*ResourceViewer) error

// WithDefaultQueryer configures ResourceViewer with the default visitor.
func WithDefaultQueryer(dashConfig config.Dash, q queryer.Queryer) ViewerOpt {
	return func(rv *ResourceViewer) error {
		visitor, err := objectvisitor.NewDefaultVisitor(dashConfig, q)
		if err != nil {
			return err
		}

		rv.visitor = visitor
		return nil
	}
}

// ResourceViewer visits an object and creates a view component.
type ResourceViewer struct {
	dashConfig config.Dash
	visitor    objectvisitor.Visitor
}

// New creates an instance of ResourceViewer.
func New(dashConfig config.Dash, opts ...ViewerOpt) (*ResourceViewer, error) {
	rv := &ResourceViewer{
		dashConfig: dashConfig,
	}

	for _, opt := range opts {
		if err := opt(rv); err != nil {
			return nil, errors.Wrap(err, "invalid resource viewer option")
		}
	}

	if rv.visitor == nil {
		return nil, errors.New("resource viewer visitor is nil")
	}

	return rv, nil
}

// Visit visits an object and creates a view component.
func (rv *ResourceViewer) Visit(ctx context.Context, object runtime.Object, handler *Handler) error {
	ctx, span := trace.StartSpan(ctx, "resourceViewer")
	defer span.End()

	if handler == nil {
		return errors.New("handler is nil")
	}

	logger := log.From(ctx).With("object", kubernetes.PrintObject(object))

	now := time.Now()
	defer func() {
		elapsed := time.Since(now)
		if elapsed > visitMaxDuration {
			logger.With("elapsed", elapsed).Debugf("ending resource viewer visit")
		}
	}()

	m, err := runtime.DefaultUnstructuredConverter.ToUnstructured(object)
	if err != nil {
		return err
	}

	u := &unstructured.Unstructured{Object: m}

	if err := rv.visitor.Visit(ctx, u, handler, true); err != nil {
		return errors.Wrapf(err, "error unable to visit object %s", kubernetes.PrintObject(object))
	}

	return nil
}
