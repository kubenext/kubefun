package api

import (
	"context"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/pkg/action"
)

const (
	RequestPerformAction = "performAction"
)

// ActionRequestManager manages action requests. Action requests allow a generic interface
// for supporting dynamic requests from clients.
type ActionRequestManager struct {
}

var _ StateManager = (*ActionRequestManager)(nil)

// NewActionRequestManager creates an instance of ActionRequestManager.
func NewActionRequestManager() *ActionRequestManager {
	return &ActionRequestManager{}
}

func (a ActionRequestManager) Start(ctx context.Context, state kubefun.State, s KubefunClient) {
}

// Handlers returns the handlers this manager supports.
func (a *ActionRequestManager) Handlers() []kubefun.ClientRequestHandler {
	return []kubefun.ClientRequestHandler{
		{
			RequestType: RequestPerformAction,
			Handler:     a.PerformAction,
		},
	}
}

// PerformAction is a handler than runs an action.
func (a *ActionRequestManager) PerformAction(state kubefun.State, payload action.Payload) error {
	ctx := context.TODO()

	actionName, err := payload.String("action")
	if err != nil {
		// TODO: alert the user this action doesn't exist (GH#493)
		return nil
	}

	if err := state.Dispatch(ctx, actionName, payload); err != nil {
		return err
	}

	return nil
}
