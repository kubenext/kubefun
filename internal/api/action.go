package api

import (
	"context"
	"github.com/kubenext/kubefun/pkg/action"
)

//go:generate mockgen -destination=./fake/mock_action_dispatcher.go -package=fake github.com/kubenext/kubefun/internal/api ActionDispatcher

// ActionDispatcher dispatches actions.
type ActionDispatcher interface {
	Dispatch(ctx context.Context, alerter action.Alerter, actionName string, payload action.Payload) error
}
