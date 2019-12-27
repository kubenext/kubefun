package configuration

import (
	"context"
	"fmt"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/internal/log"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/store"
)

type ObjectDeleter struct {
	logger log.Logger
	store  store.Store
}

func NewObjectDeleter(logger log.Logger, clusterClient store.Store) *ObjectDeleter {
	return &ObjectDeleter{
		logger: logger.With("action", kubefun.ActionDeleteObject),
		store:  clusterClient,
	}
}

func (d *ObjectDeleter) ActionName() string {
	return kubefun.ActionDeleteObject
}

func (d *ObjectDeleter) Handle(ctx context.Context, alerter action.Alerter, payload action.Payload) error {
	d.logger.With("payload", payload).Debugf("deleting object")

	key, err := store.KeyFromPayload(payload)
	if err != nil {
		return err
	}

	alertType := action.AlertTypeInfo
	message := fmt.Sprintf("Deleted %s %q", key.Kind, key.Name)
	if err := d.store.Delete(ctx, key); err != nil {
		alertType = action.AlertTypeWarning
		message = fmt.Sprintf("Unable to deleted %s %q: %s", key.Kind, key.Name, err)
	}
	alert := action.CreateAlert(alertType, message, action.DefaultAlertExpiration)
	alerter.SendAlert(alert)

	return nil
}
