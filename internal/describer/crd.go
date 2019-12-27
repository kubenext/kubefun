package describer

import (
	"context"
	"fmt"
	"github.com/kubenext/kubefun/internal/config"
	"github.com/kubenext/kubefun/internal/gvk"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/link"
	"github.com/kubenext/kubefun/internal/modules/overview/yamlviewer"
	"github.com/kubenext/kubefun/internal/printer"
	"github.com/kubenext/kubefun/internal/queryer"
	"github.com/kubenext/kubefun/internal/resourceviewer"
	"github.com/kubenext/kubefun/pkg/icon"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type crdPrinter func(ctx context.Context, crd, object *unstructured.Unstructured, options printer.Options) (component.Component, error)
type resourceViewerPrinter func(ctx context.Context, object *unstructured.Unstructured, dashConfig config.Dash, q queryer.Queryer) (component.Component, error)
type yamlPrinter func(runtime.Object) (*component.YAML, error)

type crdOption func(*crd)

type crd struct {
	base

	path                  string
	name                  string
	summaryPrinter        crdPrinter
	resourceViewerPrinter resourceViewerPrinter
	yamlPrinter           yamlPrinter
}

var _ Describer = (*crd)(nil)

func newCRD(name, path string, options ...crdOption) *crd {
	d := &crd{
		path:                  path,
		name:                  name,
		summaryPrinter:        printer.CustomResourceHandler,
		resourceViewerPrinter: createCRDResourceViewer,
		yamlPrinter:           yamlviewer.ToComponent,
	}

	for _, option := range options {
		option(d)
	}

	return d
}

func (c *crd) Describe(ctx context.Context, namespace string, options Options) (component.ContentResponse, error) {
	objectStore := options.ObjectStore()
	crd, err := CustomResourceDefinition(ctx, c.name, objectStore)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	kubefunCRD, err := kubefun.NewCustomResourceDefinition(crd)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	crdVersions, err := kubefunCRD.Versions()
	if err != nil {
		return component.EmptyContentResponse, fmt.Errorf("get versions for crd %s: %w", crd.GetName(), err)
	} else if len(crdVersions) == 0 {
		return component.EmptyContentResponse, fmt.Errorf("crd %s has no no versions", crd.GetName())
	}

	crGVK, err := gvk.CustomResource(crd, crdVersions[0])
	if err != nil {
		return component.EmptyContentResponse, fmt.Errorf("get gvk for custom resource")
	}

	apiVersion, kind := crGVK.ToAPIVersionAndKind()

	key := store.Key{
		Namespace:  namespace,
		APIVersion: apiVersion,
		Kind:       kind,
		Name:       options.Fields["name"],
	}

	object, found, err := objectStore.Get(ctx, key)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	if !found {
		return component.EmptyContentResponse, err
	}

	title := component.Title(
		component.NewText("Custom Resources"),
		component.NewText(crd.GetName()),
		component.NewText(object.GetName()))

	iconName, iconSource := loadIcon(icon.CustomResourceDefinition)
	cr := component.NewContentResponse(title)
	cr.IconName = iconName
	cr.IconSource = iconSource

	linkGenerator, err := link.NewFromDashConfig(options)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	printOptions := printer.Options{
		DashConfig: options,
		Link:       linkGenerator,
	}

	summary, err := c.summaryPrinter(ctx, crd, object, printOptions)
	if err != nil {
		return component.EmptyContentResponse, err
	}
	summary.SetAccessor("summary")

	cr.Add(summary)

	resourceViewerComponent, err := c.resourceViewerPrinter(ctx, object, options, options.Queryer)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	resourceViewerComponent.SetAccessor("resourceViewer")
	cr.Add(resourceViewerComponent)

	yvComponent, err := c.yamlPrinter(object)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	yvComponent.SetAccessor("yaml")
	cr.Add(yvComponent)

	pluginPrinter := options.PluginManager()
	tabs, err := pluginPrinter.Tabs(ctx, object)
	if err != nil {
		return component.EmptyContentResponse, errors.Wrap(err, "getting tabs from plugins")
	}

	for _, tab := range tabs {
		tab.Contents.SetAccessor(tab.Name)
		cr.Add(&tab.Contents)
	}

	return *cr, nil
}

func (c *crd) PathFilters() []PathFilter {
	return []PathFilter{
		*NewPathFilter(c.path, c),
	}
}

func createCRDResourceViewer(ctx context.Context, object *unstructured.Unstructured, dashConfig config.Dash, q queryer.Queryer) (component.Component, error) {
	rv, err := resourceviewer.New(dashConfig, resourceviewer.WithDefaultQueryer(dashConfig, q))
	if err != nil {
		return nil, err
	}

	handler, err := resourceviewer.NewHandler(dashConfig)
	if err != nil {
		return nil, err
	}

	if err := rv.Visit(ctx, object, handler); err != nil {
		return nil, err
	}

	return resourceviewer.GenerateComponent(ctx, handler, "")
}
