package event

import (
	"context"
	"github.com/golang/mock/gomock"
	dashConfigFake "github.com/kubenext/kubefun/internal/config/fake"
	"github.com/kubenext/kubefun/internal/kubeconfig"
	"github.com/kubenext/kubefun/internal/kubeconfig/fake"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_kubeContextGenerator(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	kc := &kubeconfig.KubeConfig{
		CurrentContext: "current-context",
	}

	loader := fake.NewMockLoader(controller)
	loader.EXPECT().
		Load("/path").
		Return(kc, nil)

	configLoaderFuncOpt := func(x *ContextsGenerator) {
		x.ConfigLoader = loader
	}

	dashConfig := dashConfigFake.NewMockDash(controller)
	dashConfig.EXPECT().KubeConfigPath().Return("/path")
	dashConfig.EXPECT().ContextName().Return("")

	kgc := NewContextsGenerator(dashConfig, configLoaderFuncOpt)

	assert.Equal(t, "kubeConfig", kgc.Name())

	ctx := context.Background()
	e, err := kgc.Event(ctx)
	require.NoError(t, err)

	assert.Equal(t, kubefun.EventTypeKubeConfig, e.Type)

	resp := kubeContextsResponse{
		CurrentContext: kc.CurrentContext,
	}

	assert.Equal(t, resp, e.Data)
}
