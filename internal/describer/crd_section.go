package describer

import (
	"context"
	"github.com/kubenext/kubefun/internal/gvk"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"path"
	"sort"
	"sync"
)

type CRDSection struct {
	describers map[string]Describer
	path       string
	title      string

	mu sync.Mutex
}

var _ Describer = (*CRDSection)(nil)

func NewCRDSection(p, title string) *CRDSection {
	return &CRDSection{
		describers: make(map[string]Describer),
		path:       p,
		title:      title,
	}
}

func (csd *CRDSection) Add(name string, describer Describer) {
	csd.mu.Lock()
	defer csd.mu.Unlock()

	csd.describers[name] = describer
}

func (csd *CRDSection) Remove(name string) {
	csd.mu.Lock()
	defer csd.mu.Unlock()

	delete(csd.describers, name)
}

func (csd *CRDSection) Describe(ctx context.Context, namespace string, options Options) (component.ContentResponse, error) {
	csd.mu.Lock()
	defer csd.mu.Unlock()

	var names []string
	for name := range csd.describers {
		names = append(names, name)
	}

	sort.Strings(names)

	tableCols := component.NewTableCols("Name", "Labels", "Age")
	table := component.NewTable("Custom Resources", "", tableCols)

	for _, name := range names {
		switch d := csd.describers[name].(type) {
		case *crdList:
			key := store.KeyFromGroupVersionKind(gvk.CustomResourceDefinition)
			key.Name = d.name
			crd, _, err := options.ObjectStore().Get(ctx, key)
			if err != nil {
				return component.EmptyContentResponse, err
			}

			crdObject, err := kubefun.NewCustomResourceDefinition(crd)
			if err != nil {
				return component.EmptyContentResponse, err
			}

			versions, err := crdObject.Versions()
			if err != nil {
				return component.EmptyContentResponse, err
			}

			count := 0
			for _, version := range versions {
				crGVK, err := gvk.CustomResource(crd, version)
				if err != nil {
					return component.EmptyContentResponse, err
				}
				key2 := store.KeyFromGroupVersionKind(crGVK)
				key2.Namespace = namespace
				list, _, err := options.ObjectStore().List(ctx, key2)
				if err != nil {
					return component.EmptyContentResponse, err
				}
				count += len(list.Items)
			}

			if count > 0 {
				row := component.TableRow{}

				ref := path.Join("/overview/namespace", namespace, "custom-resources", crd.GetName())
				if namespace == "" {
					ref = path.Join("/cluster-overview/custom-resources", crd.GetName())
				}

				row["Name"] = component.NewLink("", crd.GetName(), ref)
				row["Labels"] = component.NewLabels(crd.GetLabels())
				row["Age"] = component.NewTimestamp(crd.GetCreationTimestamp().Time)

				table.Add(row)
			}

		}
	}

	cr := component.ContentResponse{
		Components: []component.Component{table},
		Title:      component.TitleFromString(csd.title),
	}

	return cr, nil
}

func (csd *CRDSection) PathFilters() []PathFilter {
	return []PathFilter{
		*NewPathFilter(csd.path, csd),
	}
}

func (csd *CRDSection) Reset(ctx context.Context) error {
	csd.mu.Lock()
	defer csd.mu.Unlock()

	logger := log.From(ctx)

	for name := range csd.describers {
		logger.With("describer-name", name, "crd-section-path", csd.path).
			Debugf("removing crd from section")
		delete(csd.describers, name)
	}

	return nil
}
