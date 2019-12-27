package api_test

import (
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/api"
	kubefunFake "github.com/kubenext/kubefun/internal/kubefun/fake"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestActionRequestManager_Handlers(t *testing.T) {
	manager := api.NewActionRequestManager()
	AssertHandlers(t, manager, []string{api.RequestPerformAction})
}

func TestActionRequestManager_PerformAction(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	state := kubefunFake.NewMockState(controller)

	manager := api.NewActionRequestManager()

	payload := action.CreatePayload(api.RequestPerformAction, map[string]interface{}{
		"foo": "bar",
	})

	state.EXPECT().
		Dispatch(gomock.Any(), api.RequestPerformAction, payload).
		Return(nil)

	require.NoError(t, manager.PerformAction(state, payload))
}
