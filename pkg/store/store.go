package store

import (
	"context"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

// TODO store for pkg

type Key struct {
	Namespace  string
	APIVersion string
	Kind       string
	Name       string
	Selector   *labels.Set
}

// Store stores Kubernetes objects.
type Store interface {
	List(ctx context.Context, key Key) (list *unstructured.UnstructuredList, loading bool, err error)
	Get(ctx context.Context, key Key) (object *unstructured.Unstructured, found bool, err error)
	Delete(ctx context.Context, key Key) error
	Watch(ctx context.Context, key Key, handler cache.ResourceEventHandler) error
	Unwatch(ctx context.Context, groupVersionKinds ...schema.GroupVersionKind) error
	//UpdateClusterClient(ctx context.Context, client cluster.ClientInterface) error
	RegisterOnUpdate(fn UpdateFn)
	IsLoading(ctx context.Context, key Key) bool
}

type UpdateFn func(store Store)
