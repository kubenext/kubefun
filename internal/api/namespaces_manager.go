package api

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/kubenext/kubefun/internal/cluster"
	"github.com/kubenext/kubefun/internal/event"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/pkg/errors"
)

// NamespaceManagerConfig is configuration for NamespacesManager.
type NamespaceManagerConfig interface {
	ClusterClient() cluster.ClientInterface
}

// NamespacesManagerOption is an option for configuring NamespacesManager.
type NamespacesManagerOption func(n *NamespacesManager)

// NamespacesGenerateFunc is a function that generates a list of namespaces.
type NamespacesGenerateFunc func(ctx context.Context, config NamespaceManagerConfig) ([]string, error)

// WithNamespacesGenerator configures the namespaces generator function.
func WithNamespacesGenerator(fn NamespacesGenerateFunc) NamespacesManagerOption {
	return func(n *NamespacesManager) {
		n.namespacesGeneratorFunc = fn
	}
}

// WithNamespacesGeneratorPoller configures the poller.
func WithNamespacesGeneratorPoller(poller Poller) NamespacesManagerOption {
	return func(n *NamespacesManager) {
		n.poller = poller
	}
}

// NamespacesManager manages namespaces.
type NamespacesManager struct {
	config                  NamespaceManagerConfig
	namespacesGeneratorFunc NamespacesGenerateFunc
	poller                  Poller
}

var _ StateManager = (*NamespacesManager)(nil)

// NewNamespacesManager creates an instance of NamespacesManager.
func NewNamespacesManager(config NamespaceManagerConfig, options ...NamespacesManagerOption) *NamespacesManager {
	n := &NamespacesManager{
		config:                  config,
		poller:                  NewInterruptiblePoller("namespaces"),
		namespacesGeneratorFunc: NamespacesGenerator,
	}

	for _, option := range options {
		option(n)
	}

	return n
}

// Handlers returns nil.
func (n NamespacesManager) Handlers() []kubefun.ClientRequestHandler {
	return nil
}

// Start starts the manager. It periodically generates a list of namespaces.
func (n *NamespacesManager) Start(ctx context.Context, state kubefun.State, s KubefunClient) {
	ch := make(chan struct{}, 1)
	defer func() {
		close(ch)
	}()

	n.poller.Run(ctx, ch, n.runUpdate(state, s), event.DefaultScheduleDelay)
}

func (n *NamespacesManager) runUpdate(state kubefun.State, client KubefunClient) PollerFunc {
	var previous []byte

	return func(ctx context.Context) bool {
		logger := log.From(ctx)

		namespaces, err := n.namespacesGeneratorFunc(ctx, n.config)
		if err != nil {
			logger.WithErr(err).Errorf("load namespaces")
			return false
		}

		if ctx.Err() == nil {
			cur, err := json.Marshal(namespaces)
			if err != nil {
				logger.WithErr(err).Errorf("unable to marshal namespaces")
				return false
			}

			if bytes.Compare(previous, cur) != 0 {
				previous = cur
				client.Send(CreateNamespacesEvent(namespaces))
			}
		}

		return false
	}
}

// NamespacesGenerator generates a list of namespaces.
func NamespacesGenerator(_ context.Context, config NamespaceManagerConfig) ([]string, error) {
	if config == nil {
		return nil, errors.New("namespaces manager config is nil")
	}

	clusterClient := config.ClusterClient()
	namespaceClient, err := clusterClient.NamespaceClient()
	if err != nil {
		return nil, errors.Wrap(err, "retrieve namespaces client")
	}

	names, err := namespaceClient.Names()
	if err != nil {
		initialNamespace := namespaceClient.InitialNamespace()
		names = []string{initialNamespace}
	}

	return names, nil
}

// CreateNamespacesEvent creates a namespaces event.
func CreateNamespacesEvent(namespaces []string) kubefun.Event {
	return kubefun.Event{
		Type: kubefun.EventTypeNamespaces,
		Data: map[string]interface{}{
			"namespaces": namespaces,
		},
	}
}
