package generator

import (
	"context"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/config"
	"github.com/kubenext/kubefun/internal/describer"
	"github.com/kubenext/kubefun/internal/link"
	"github.com/kubenext/kubefun/internal/printer"
	"github.com/kubenext/kubefun/internal/queryer"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
	kLabels "k8s.io/apimachinery/pkg/labels"
)

// Interface generates content.
type Interface interface {
	Generate(ctx context.Context, contentPath string, opts Options) (component.ContentResponse, error)
}

// Generator is an implementation of Interface that generates content.
type Generator struct {
	pathMatcher *describer.PathMatcher
	printer     printer.Printer
	dashConfig  config.Dash
}

var _ Interface = (*Generator)(nil)

// Options are additional options to pass a Generator
type Options struct {
	LabelSet *kLabels.Set
}

// NewGenerator creates a Generator.
func NewGenerator(pm *describer.PathMatcher, dashConfig config.Dash) (*Generator, error) {
	p := printer.NewResource(dashConfig)

	if err := printer.AddHandlers(p); err != nil {
		return nil, errors.Wrap(err, "add print handlers")
	}

	if pm == nil {
		return nil, errors.New("path matcher is nil")
	}

	return &Generator{
		pathMatcher: pm,
		printer:     p,
		dashConfig:  dashConfig,
	}, nil
}

// Generate generates a content response.
func (g *Generator) Generate(ctx context.Context, contentPath string, opts Options) (component.ContentResponse, error) {
	ctx, span := trace.StartSpan(ctx, "Generate")
	defer span.End()

	pf, err := g.pathMatcher.Find(contentPath)
	if err != nil {
		if err == describer.ErrPathNotFound {
			return component.EmptyContentResponse, api.NewNotFoundError(contentPath)
		}
		return component.EmptyContentResponse, err
	}

	discoveryInterface, err := g.dashConfig.ClusterClient().DiscoveryClient()
	if err != nil {
		return component.EmptyContentResponse, err
	}

	linkGenerator, err := link.NewFromDashConfig(g.dashConfig)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	q := queryer.New(g.dashConfig.ObjectStore(), discoveryInterface)

	loaderFactory := describer.NewObjectLoaderFactory(g.dashConfig)

	fields := pf.Fields(contentPath)
	namespace := ""
	if n, ok := fields["namespace"]; ok {
		namespace = n
	}

	options := describer.Options{
		Queryer:  q,
		Fields:   fields,
		Printer:  g.printer,
		LabelSet: opts.LabelSet,
		Dash:     g.dashConfig,
		Link:     linkGenerator,

		LoadObjects: loaderFactory.LoadObjects,
		LoadObject:  loaderFactory.LoadObject,
	}

	cResponse, err := pf.Describer.Describe(ctx, namespace, options)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	return cResponse, nil
}
