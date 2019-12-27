package portforward

import (
	"context"
	"github.com/kubenext/kubefun/internal/cluster"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/pkg/errors"
	"os"
)

// Default create a port forward instance.
func Default(ctx context.Context, client cluster.ClientInterface, objectStore store.Store) (PortForwarder, error) {
	restClient, err := client.RESTClient()
	if err != nil {
		return nil, errors.Wrap(err, "fetching RESTClient")
	}

	pfOpts := ServiceOptions{
		RESTClient:  restClient,
		Config:      client.RESTConfig(),
		ObjectStore: objectStore,
		PortForwarder: &DefaultPortForwarder{
			IOStreams: IOStreams{
				In:     os.Stdin,
				Out:    os.Stdout,
				ErrOut: os.Stderr,
			},
		},
	}

	svc := New(ctx, pfOpts)

	return svc, nil
}
