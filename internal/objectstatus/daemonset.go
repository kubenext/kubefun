package objectstatus

import (
	"context"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes/scheme"
)

func daemonSet(_ context.Context, object runtime.Object, _ store.Store) (ObjectStatus, error) {
	if object == nil {
		return ObjectStatus{}, errors.Errorf("daemon set is nil")
	}

	ds := &appsv1.DaemonSet{}

	if err := scheme.Scheme.Convert(object, ds, 0); err != nil {
		return ObjectStatus{}, errors.Wrap(err, "convert object to daemon set")
	}

	status := ds.Status

	switch {
	case status.NumberMisscheduled > 0:
		return ObjectStatus{
			nodeStatus: component.NodeStatusWarning,
			Details:    []component.Component{component.NewText("Daemon Set pods are running on nodes that aren't supposed to run Daemon Set pods")},
		}, nil
	case status.NumberAvailable != status.NumberReady:
		return ObjectStatus{
			nodeStatus: component.NodeStatusWarning,
			Details:    []component.Component{component.NewText("Daemon Set pods are not ready")},
		}, nil
	default:
		return ObjectStatus{
			nodeStatus: component.NodeStatusOK,
			Details:    []component.Component{component.NewText("Daemon Set is OK")},
		}, nil
	}
}
