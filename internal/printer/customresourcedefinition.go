package printer

import (
	"context"
	"fmt"
	"github.com/kubenext/kubefun/internal/gvk"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func CustomResourceDefinitionHandler(ctx context.Context, crd *unstructured.Unstructured, namespace string, options Options) (component.Component, error) {
	object := NewObject(crd)
	object.EnableEvents()

	config, err := printCustomResourceDefinitionConfig(crd)
	if err != nil {
		return nil, err
	}
	object.RegisterConfig(config)

	kubefunCRD, err := kubefun.NewCustomResourceDefinition(crd)
	if err != nil {
		return nil, err
	}

	objectStore := options.DashConfig.ObjectStore()

	versions, err := kubefunCRD.Versions()
	if err != nil {
		return nil, err
	}

	for _, version := range versions {
		object.RegisterItems(ItemDescriptor{
			Func: func() (c component.Component, err error) {
				crGVK, err := gvk.CustomResource(crd, version)
				if err != nil {
					return nil, err
				}

				key := store.KeyFromGroupVersionKind(crGVK)
				key.Namespace = namespace

				customResources, _, err := objectStore.List(ctx, key)
				if err != nil {
					return nil, err
				}

				return CustomResourceListHandler(crd, customResources, version, options.Link)
			},
			Width: component.WidthFull,
		})

	}

	view, err := object.ToComponent(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("print custom resource definition: %w", err)
	}

	return view, nil
}

func printCustomResourceDefinitionConfig(crd *unstructured.Unstructured) (*component.Summary, error) {
	if crd == nil {
		return nil, fmt.Errorf("custom resource definition is nil")
	}

	summary := component.NewSummary("Config")

	group, err := nestedString(crd, "spec", "group")
	if err != nil {
		return nil, err
	}

	kind, err := nestedString(crd, "spec", "names", "kind")
	if err != nil {
		return nil, err
	}

	summary.AddSection("Group", component.NewText(group))
	summary.AddSection("Kind", component.NewText(kind))

	return summary, nil
}

func nestedString(object *unstructured.Unstructured, fields ...string) (string, error) {
	s, found, err := unstructured.NestedString(object.Object, fields...)
	if err != nil {
		return "", err
	}

	if !found {
		return "", nil
	}

	return s, nil
}
