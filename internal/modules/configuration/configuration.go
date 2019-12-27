package configuration

import (
	"context"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/config"
	"github.com/kubenext/kubefun/internal/describer"
	"github.com/kubenext/kubefun/internal/event"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/module"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/icon"
	"github.com/kubenext/kubefun/pkg/navigation"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"path"
)

type Options struct {
	DashConfig     config.Dash
	KubeConfigPath string
}

type Configuration struct {
	Options

	pathMatcher          *describer.PathMatcher
	kubeContextGenerator *event.ContextsGenerator
}

var _ module.Module = (*Configuration)(nil)
var _ module.ActionReceiver = (*Configuration)(nil)

func New(ctx context.Context, options Options) *Configuration {
	pm := describer.NewPathMatcher("configuration")
	for _, pf := range rootDescriber.PathFilters() {
		pm.Register(ctx, pf)
	}

	return &Configuration{
		Options:              options,
		pathMatcher:          pm,
		kubeContextGenerator: event.NewContextsGenerator(options.DashConfig),
	}
}

func (Configuration) Name() string {
	return "configuration"
}

func (c Configuration) ClientRequestHandlers() []kubefun.ClientRequestHandler {
	return nil
}

func (c *Configuration) SetContext(ctx context.Context, contextName string) error {
	return nil
}

func (c *Configuration) Content(ctx context.Context, contentPath string, opts module.ContentOptions) (component.ContentResponse, error) {
	pf, err := c.pathMatcher.Find(contentPath)
	if err != nil {
		if err == describer.ErrPathNotFound {
			return component.EmptyContentResponse, api.NewNotFoundError(contentPath)
		}
		return component.EmptyContentResponse, err
	}

	options := describer.Options{
		Fields:   pf.Fields(contentPath),
		LabelSet: opts.LabelSet,
		Dash:     c.DashConfig,
	}

	cResponse, err := pf.Describer.Describe(ctx, "", options)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	return cResponse, nil
}

func (c *Configuration) ContentPath() string {
	return c.Name()
}

func (c *Configuration) Navigation(ctx context.Context, namespace, root string) ([]navigation.Navigation, error) {
	return []navigation.Navigation{
		{
			Title:    "Configuration",
			Path:     path.Join(c.ContentPath(), "/"),
			IconName: icon.Configuration,
			Children: []navigation.Navigation{
				{
					Title:    "Plugins",
					Path:     path.Join(c.ContentPath(), "plugins"),
					IconName: icon.ConfigurationPlugin,
				},
			},
		},
	}, nil
}

func (Configuration) SetNamespace(namespace string) error {
	return nil
}

func (Configuration) Start() error {
	return nil
}

func (Configuration) Stop() {
}

func (c Configuration) SupportedGroupVersionKind() []schema.GroupVersionKind {
	return []schema.GroupVersionKind{}
}

func (c Configuration) GroupVersionKindPath(namespace, apiVersion, kind, name string) (string, error) {
	return "", errors.Errorf("configuration can't create paths for %s %s", apiVersion, kind)
}

func (c Configuration) AddCRD(ctx context.Context, crd *unstructured.Unstructured) error {
	return nil
}

func (c Configuration) RemoveCRD(ctx context.Context, crd *unstructured.Unstructured) error {
	return nil
}

func (c Configuration) ResetCRDs(ctx context.Context) error {
	return nil
}

// Generators allow modules to send events to the frontend.
func (c Configuration) Generators() []kubefun.Generator {
	return []kubefun.Generator{
		c.kubeContextGenerator,
	}
}

func (c *Configuration) ActionPaths() map[string]action.DispatcherFunc {
	objectDeleter := NewObjectDeleter(c.DashConfig.Logger(), c.DashConfig.ObjectStore())

	return map[string]action.DispatcherFunc{
		objectDeleter.ActionName(): objectDeleter.Handle,
	}
}
