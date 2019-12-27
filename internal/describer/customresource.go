package describer

import (
	"context"
	"fmt"
	"github.com/kubenext/kubefun/internal/gvk"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/internal/module"
	"github.com/kubenext/kubefun/internal/util/kubernetes"
	"github.com/kubenext/kubefun/pkg/store"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"path"
)

func CustomResourceDefinition(ctx context.Context, name string, o store.Store) (*unstructured.Unstructured, error) {
	key := store.KeyFromGroupVersionKind(gvk.CustomResourceDefinition)
	key.Name = name

	crd, _, err := o.Get(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("get %s: %w", key, err)
	}

	return crd, nil
}

func AddCRD(ctx context.Context, crd *unstructured.Unstructured, pm *PathMatcher, crdSection *CRDSection, m module.Module) {
	name := crd.GetName()

	logger := log.From(ctx).With("crd-name", name, "module", m.Name())
	logger.Debugf("adding CRD")

	cld := newCRDList(name, crdListPath(name))

	// TODO: this should add a list of custom resource definitions (GH#509)
	crdSection.Add(name, cld)

	for _, pf := range cld.PathFilters() {
		pm.Register(ctx, pf)
	}

	cd := newCRD(name, crdObjectPath(name))
	for _, pf := range cd.PathFilters() {
		pm.Register(ctx, pf)
	}

	if err := m.AddCRD(ctx, crd); err != nil {
		logger.WithErr(err).Errorf("unable to add CRD")
	}
}

func DeleteCRD(ctx context.Context, crd *unstructured.Unstructured, pm *PathMatcher, crdSection *CRDSection, m module.Module, s store.Store) {
	name := crd.GetName()

	logger := log.From(ctx).With("crd-name", name, "module", m.Name())
	logger.Debugf("deleting CRD")

	pm.Deregister(ctx, crdListPath(name))
	pm.Deregister(ctx, crdObjectPath(name))

	crdSection.Remove(name)

	if err := m.RemoveCRD(ctx, crd); err != nil {
		logger.WithErr(err).Errorf("unable to remove CRD")
	}

	list, err := kubernetes.CRDResources(crd)
	if err != nil {
		logger.WithErr(err).Errorf("unable to get group/version/kinds for CRD")

	}

	if err := s.Unwatch(ctx, list...); err != nil {
		logger.WithErr(err).Errorf("unable to unwatch CRD")
		return
	}
}

func crdListPath(name string) string {
	return path.Join("/custom-resources", name)
}

func crdObjectPath(name string) string {
	return path.Join(crdListPath(name), ResourceNameRegex)
}
