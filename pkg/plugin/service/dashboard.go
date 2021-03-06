package service

import (
	"context"
	"github.com/kubenext/kubefun/pkg/plugin/api"
	"github.com/kubenext/kubefun/pkg/store"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

//go:generate mockgen -destination=./fake/mock_dashboard.go -package=fake github.com/kubenext/kubefun/pkg/plugin/service Dashboard

// Dashboard is the client a plugin can use to interact with Kubefun.
type Dashboard interface {
	Close() error
	List(ctx context.Context, key store.Key) (*unstructured.UnstructuredList, error)
	Get(ctx context.Context, key store.Key) (*unstructured.Unstructured, bool, error)
	Update(ctx context.Context, object *unstructured.Unstructured) error
	PortForward(ctx context.Context, req api.PortForwardRequest) (api.PortForwardResponse, error)
	CancelPortForward(ctx context.Context, id string)
	ForceFrontendUpdate(ctx context.Context) error
}

// NewDashboardClient creates a dashboard client.
func NewDashboardClient(dashboardAPIAddress string) (Dashboard, error) {
	client, err := api.NewClient(dashboardAPIAddress)
	if err != nil {
		return nil, err
	}

	return client, nil
}
