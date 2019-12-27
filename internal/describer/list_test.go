package describer

import (
	"context"
	"github.com/golang/mock/gomock"
	configFake "github.com/kubenext/kubefun/internal/config/fake"
	printerFake "github.com/kubenext/kubefun/internal/printer/fake"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/plugin"
	pluginFake "github.com/kubenext/kubefun/pkg/plugin/fake"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"
)

func TestListDescriber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	thePath := "/"

	pod := testutil.CreatePod("pod")
	pod.CreationTimestamp = *testutil.CreateTimestamp()

	key, err := store.KeyFromObject(pod)
	require.NoError(t, err)

	ctx := context.Background()
	namespace := "default"

	dashConfig := configFake.NewMockDash(controller)
	moduleRegistrar := pluginFake.NewMockModuleRegistrar(controller)
	actionRegistrar := pluginFake.NewMockActionRegistrar(controller)

	pluginManager := plugin.NewManager(nil, moduleRegistrar, actionRegistrar)
	dashConfig.EXPECT().PluginManager().Return(pluginManager)

	podListTable := createPodTable(*pod)

	objectPrinter := printerFake.NewMockPrinter(controller)
	podList := &corev1.PodList{Items: []corev1.Pod{*pod}}
	objectPrinter.EXPECT().Print(gomock.Any(), podList, pluginManager).Return(podListTable, nil)

	options := Options{
		Dash:    dashConfig,
		Printer: objectPrinter,
		LoadObjects: func(ctx context.Context, namespace string, fields map[string]string, objectStoreKeys []store.Key) (*unstructured.UnstructuredList, error) {
			return testutil.ToUnstructuredList(t, pod), nil
		},
	}

	listConfig := ListConfig{
		Path:          thePath,
		Title:         "list",
		StoreKey:      key,
		ListType:      podListType,
		ObjectType:    podObjectType,
		IsClusterWide: false,
		IconName:      "icon-name",
		IconSource:    "icon-source",
	}
	d := NewList(listConfig)
	cResponse, err := d.Describe(ctx, namespace, options)
	require.NoError(t, err)

	list := component.NewList("list", nil)
	list.Add(podListTable)
	list.SetIcon("icon-name", "icon-source")
	expected := component.ContentResponse{
		Components: []component.Component{list},
	}

	assert.Equal(t, expected, cResponse)
}
