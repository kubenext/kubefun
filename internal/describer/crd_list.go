package describer

import (
	"context"
	"github.com/kubenext/kubefun/internal/link"
	"github.com/kubenext/kubefun/internal/modules/overview/yamlviewer"
	"github.com/kubenext/kubefun/internal/printer"
	"github.com/kubenext/kubefun/pkg/view/component"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type crdListPrinter func(crdObject *unstructured.Unstructured, resources *unstructured.UnstructuredList, version string, linkGenerator link.Interface) (component.Component, error)

type crdListDescriptionOption func(*crdList)

type crdList struct {
	base

	name    string
	path    string
	printer crdListPrinter
}

var _ Describer = (*crdList)(nil)

func newCRDList(name, path string, options ...crdListDescriptionOption) *crdList {
	d := &crdList{
		name:    name,
		path:    path,
		printer: printer.CustomResourceListHandler,
	}

	for _, option := range options {
		option(d)
	}

	return d
}

func (cld *crdList) Describe(ctx context.Context, namespace string, options Options) (component.ContentResponse, error) {
	objectStore := options.ObjectStore()

	crd, err := CustomResourceDefinition(ctx, cld.name, objectStore)
	if err != nil {
		return component.EmptyContentResponse, err
	}

	printOptions := printer.Options{
		DashConfig: options.Dash,
		Link:       options.Link,
	}

	view, err := printer.CustomResourceDefinitionHandler(ctx, crd, namespace, printOptions)
	if err != nil {
		return component.EmptyContentResponse, err
	}
	view.SetAccessor("summary")

	title := component.Title(
		component.NewText("Custom Resources"),
		component.NewText(crd.GetName()))

	contentResponse := component.NewContentResponse(title)
	contentResponse.Add(view)

	yamlView, err := yamlviewer.ToComponent(crd)
	if err != nil {
		return component.EmptyContentResponse, err
	}
	yamlView.SetAccessor("yaml")

	contentResponse.Add(yamlView)

	return *contentResponse, nil
}

func (cld *crdList) PathFilters() []PathFilter {
	return []PathFilter{
		*NewPathFilter(cld.path, cld),
	}
}
