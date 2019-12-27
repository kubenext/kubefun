package api_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/api/fake"
	configFake "github.com/kubenext/kubefun/internal/config/fake"
	"github.com/kubenext/kubefun/internal/kubefun"
	kubefunFake "github.com/kubenext/kubefun/internal/kubefun/fake"
	"github.com/kubenext/kubefun/internal/log"
	"testing"
)

func TestContextManager_Handlers(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dashConfig := configFake.NewMockDash(controller)

	manager := api.NewContextManager(dashConfig)
	AssertHandlers(t, manager, []string{api.RequestSetContext})
}

func TestContext_GenerateContexts(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	state := kubefunFake.NewMockState(controller)
	kubefunClient := fake.NewMockKubefunClient(controller)

	ev := kubefun.Event{
		Type: "eventType",
	}
	kubefunClient.EXPECT().Send(ev)

	logger := log.NopLogger()

	dashConfig := configFake.NewMockDash(controller)
	dashConfig.EXPECT().Logger().Return(logger).AnyTimes()

	poller := api.NewSingleRunPoller()
	generatorFunc := func(ctx context.Context, state kubefun.State) (kubefun.Event, error) {
		return ev, nil
	}
	manager := api.NewContextManager(dashConfig,
		api.WithContextGenerator(generatorFunc),
		api.WithContextGeneratorPoll(poller))

	ctx := context.Background()
	manager.Start(ctx, state, kubefunClient)
}
