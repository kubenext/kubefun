package plugin_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/module"
	"github.com/kubenext/kubefun/pkg/navigation"
	"github.com/kubenext/kubefun/pkg/plugin"
	"github.com/kubenext/kubefun/pkg/plugin/fake"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModuleProxy_Name(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := fake.NewMockModuleService(controller)

	metadata := &plugin.Metadata{
		Name: "Test Plugin",
	}

	moduleProxy, err := plugin.NewModuleProxy("plugin-name", metadata, service)
	require.NoError(t, err)

	assert.Equal(t, metadata.Name, moduleProxy.Name())
}

func TestModuleProxy_ContentPath(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := fake.NewMockModuleService(controller)

	response := component.ContentResponse{}
	service.EXPECT().
		Content(gomock.Any(), "/path").
		Return(response, nil)

	metadata := &plugin.Metadata{
		Name: "Test Plugin",
	}

	moduleProxy, err := plugin.NewModuleProxy("plugin-name", metadata, service)
	require.NoError(t, err)

	ctx := context.Background()
	got, err := moduleProxy.Content(ctx, "/path", module.ContentOptions{})
	require.NoError(t, err)

	assert.Equal(t, response, got)
}

func TestModuleProxy_Navigation(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := fake.NewMockModuleService(controller)

	nav := navigation.Navigation{}
	service.EXPECT().
		Navigation(gomock.Any()).
		Return(nav, nil)

	metadata := &plugin.Metadata{
		Name: "Test Plugin",
	}

	moduleProxy, err := plugin.NewModuleProxy("plugin-name", metadata, service)
	require.NoError(t, err)

	ctx := context.Background()
	got, err := moduleProxy.Navigation(ctx, "", "")
	require.NoError(t, err)

	expected := []navigation.Navigation{
		nav,
	}

	assert.Equal(t, expected, got)
}
