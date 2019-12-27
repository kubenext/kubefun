package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/api/fake"
	configFake "github.com/kubenext/kubefun/internal/config/fake"
	"github.com/kubenext/kubefun/internal/kubefun"
	kubefunFake "github.com/kubenext/kubefun/internal/kubefun/fake"
	"github.com/kubenext/kubefun/internal/module"
	moduleFake "github.com/kubenext/kubefun/internal/module/fake"
	"github.com/kubenext/kubefun/pkg/navigation"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNavigationManager_GenerateNavigation(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dashConfig := configFake.NewMockDash(controller)

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().GetContentPath().Return("/path")

	kubefunClient := fake.NewMockKubefunClient(controller)

	sections := []navigation.Navigation{{Title: "module"}}

	kubefunClient.EXPECT().
		Send(api.CreateNavigationEvent(sections, "/path"))

	poller := api.NewSingleRunPoller()
	manager := api.NewNavigationManager(dashConfig,
		api.WithNavigationGeneratorPoller(poller),
		api.WithNavigationGenerator(func(ctx context.Context, state kubefun.State, config api.NavigationManagerConfig) ([]navigation.Navigation, error) {
			return sections, nil
		}),
	)

	ctx := context.Background()
	manager.Start(ctx, state, kubefunClient)
}

func TestNavigationGenerator(t *testing.T) {
	tests := []struct {
		name     string
		setup    func(controller *gomock.Controller) (*configFake.MockDash, *kubefunFake.MockState)
		isErr    bool
		expected []navigation.Navigation
	}{
		{
			name: "in general",
			setup: func(controller *gomock.Controller) (*configFake.MockDash, *kubefunFake.MockState) {
				m := moduleFake.NewMockModule(controller)
				m.EXPECT().ContentPath().Return("/module")
				m.EXPECT().Name().Return("module").AnyTimes()
				m.EXPECT().
					Navigation(gomock.Any(), "default", "/module").
					Return([]navigation.Navigation{
						{Title: "module"},
					}, nil)

				moduleManager := moduleFake.NewMockManagerInterface(controller)
				moduleManager.EXPECT().Modules().Return([]module.Module{m})

				dashConfig := configFake.NewMockDash(controller)
				dashConfig.EXPECT().ModuleManager().Return(moduleManager)

				state := kubefunFake.NewMockState(controller)
				state.EXPECT().GetNamespace().Return("default")

				return dashConfig, state
			},
			expected: []navigation.Navigation{
				{Title: "module"},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			require.NotNil(t, test.setup)
			dashConfig, state := test.setup(controller)

			ctx := context.Background()
			got, err := api.NavigationGenerator(ctx, state, dashConfig)
			require.NoError(t, err)

			require.Equal(t, test.expected, got)
		})
	}
}
