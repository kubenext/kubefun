package resourceviewer

import (
	"context"
	"github.com/kubenext/kubefun/pkg/view/component"
	"k8s.io/apimachinery/pkg/types"
)

//go:generate mockgen -destination=./fake/mock_details.go -package=fake github.com/kubenext/kubefun/internal/resourceviewer Details

// Details generates details for resource viewer.
type Details interface {
	AdjacencyList() (*component.AdjList, error)
	Nodes(ctx context.Context) (component.Nodes, error)
}

// GenerateComponent generates a resource viewer component given details.
func GenerateComponent(ctx context.Context, details Details, selected types.UID) (*component.ResourceViewer, error) {
	rv := component.NewResourceViewer("Resource Viewer")
	rv.SetAccessor("resourceViewer")

	nodes, err := details.Nodes(ctx)
	if err != nil {
		return nil, err
	}

	for id, node := range nodes {
		rv.AddNode(id, node)
	}

	edges, err := details.AdjacencyList()
	if err != nil {
		return nil, err
	}

	for k, list := range *edges {
		for i := range list {
			item := list[i]
			if err := rv.AddEdge(k, item.Node, item.Type); err != nil {
				return nil, err
			}
		}
	}

	rv.Select(string(selected))

	return rv, nil
}
