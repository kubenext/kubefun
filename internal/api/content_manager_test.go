package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/api/fake"
	"github.com/kubenext/kubefun/internal/kubefun"
	kubefunFake "github.com/kubenext/kubefun/internal/kubefun/fake"
	"github.com/kubenext/kubefun/internal/log"
	moduleFake "github.com/kubenext/kubefun/internal/module/fake"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestContentManager_Handlers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	moduleManager := moduleFake.NewMockManagerInterface(controller)

	logger := log.NopLogger()

	manager := api.NewContentManager(moduleManager, logger)
	AssertHandlers(t, manager, []string{
		api.RequestSetContentPath,
		api.RequestSetNamespace,
	})
}

func TestContentManager_GenerateContent(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	params := map[string][]string{}

	moduleManager := moduleFake.NewMockManagerInterface(controller)
	state := kubefunFake.NewMockState(controller)

	state.EXPECT().GetContentPath().Return("/path")
	state.EXPECT().GetNamespace().Return("default")
	state.EXPECT().GetQueryParams().Return(params)
	state.EXPECT().OnContentPathUpdate(gomock.Any()).DoAndReturn(func(fn kubefun.ContentPathUpdateFunc) kubefun.UpdateCancelFunc {
		fn("foo")
		return func() {}
	})
	kubefunClient := fake.NewMockKubefunClient(controller)

	contentResponse := component.ContentResponse{
		IconName: "fake",
	}
	contentEvent := api.CreateContentEvent(contentResponse, "default", "/path", params)
	kubefunClient.EXPECT().Send(contentEvent).AnyTimes()

	logger := log.NopLogger()

	poller := api.NewSingleRunPoller()

	contentGenerator := func(ctx context.Context, state kubefun.State) (component.ContentResponse, bool, error) {
		return contentResponse, false, nil
	}
	manager := api.NewContentManager(moduleManager, logger,
		api.WithContentGenerator(contentGenerator),
		api.WithContentGeneratorPoller(poller))

	ctx := context.Background()
	manager.Start(ctx, state, kubefunClient)
}

func TestContentManager_SetContentPath(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := moduleFake.NewMockModule(controller)
	m.EXPECT().Name().Return("name").AnyTimes()

	moduleManager := moduleFake.NewMockManagerInterface(controller)

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().SetContentPath("/path")

	logger := log.NopLogger()

	manager := api.NewContentManager(moduleManager, logger,
		api.WithContentGeneratorPoller(api.NewSingleRunPoller()))

	payload := action.Payload{
		"contentPath": "/path",
	}

	require.NoError(t, manager.SetContentPath(state, payload))
}

func TestContentManager_SetNamespace(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	m := moduleFake.NewMockModule(controller)
	m.EXPECT().Name().Return("name").AnyTimes()

	moduleManager := moduleFake.NewMockManagerInterface(controller)

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().SetNamespace("kube-system")

	logger := log.NopLogger()

	manager := api.NewContentManager(moduleManager, logger,
		api.WithContentGeneratorPoller(api.NewSingleRunPoller()))

	payload := action.Payload{
		"namespace": "kube-system",
	}

	require.NoError(t, manager.SetNamespace(state, payload))
}

func TestContentManager_SetQueryParams(t *testing.T) {
	tests := []struct {
		name    string
		payload action.Payload
		setup   func(state *kubefunFake.MockState)
	}{
		{
			name: "single filter",
			payload: action.Payload{
				"params": map[string]interface{}{
					"filters": "foo:bar",
				},
			},
			setup: func(state *kubefunFake.MockState) {
				state.EXPECT().SetFilters([]kubefun.Filter{
					{Key: "foo", Value: "bar"},
				})
			},
		},
		{
			name: "multiple filters",
			payload: action.Payload{
				"params": map[string]interface{}{
					"filters": []interface{}{
						"foo:bar",
						"baz:qux",
					},
				},
			},
			setup: func(state *kubefunFake.MockState) {
				state.EXPECT().SetFilters([]kubefun.Filter{
					{Key: "foo", Value: "bar"},
					{Key: "baz", Value: "qux"},
				})
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			m := moduleFake.NewMockModule(controller)
			m.EXPECT().Name().Return("name").AnyTimes()

			moduleManager := moduleFake.NewMockManagerInterface(controller)

			state := kubefunFake.NewMockState(controller)
			require.NotNil(t, test.setup)
			test.setup(state)

			logger := log.NopLogger()

			manager := api.NewContentManager(moduleManager, logger,
				api.WithContentGeneratorPoller(api.NewSingleRunPoller()))
			require.NoError(t, manager.SetQueryParams(state, test.payload))
		})
	}
}
