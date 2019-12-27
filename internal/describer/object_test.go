package describer

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/plugin"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"

	configFake "github.com/kubenext/kubefun/internal/config/fake"
	printerFake "github.com/kubenext/kubefun/internal/printer/fake"
	pluginFake "github.com/kubenext/kubefun/pkg/plugin/fake"
)

func TestObjectDescriber(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ctx := context.Background()
	thePath := "/"

	pod := testutil.CreatePod("pod")
	pod.CreationTimestamp = *testutil.CreateTimestamp()

	key, err := store.KeyFromObject(pod)
	require.NoError(t, err)

	dashConfig := configFake.NewMockDash(controller)
	moduleRegistrar := pluginFake.NewMockModuleRegistrar(controller)
	actionRegistrar := pluginFake.NewMockActionRegistrar(controller)

	pluginManager := plugin.NewManager(nil, moduleRegistrar, actionRegistrar)
	dashConfig.EXPECT().PluginManager().Return(pluginManager).AnyTimes()

	objectPrinter := printerFake.NewMockPrinter(controller)

	podSummary := component.NewText("summary")
	objectPrinter.EXPECT().Print(gomock.Any(), pod, pluginManager).Return(podSummary, nil)

	options := Options{
		Dash:    dashConfig,
		Printer: objectPrinter,
		LoadObject: func(ctx context.Context, namespace string, fields map[string]string, objectStoreKey store.Key) (*unstructured.Unstructured, error) {
			return testutil.ToUnstructured(t, pod), nil
		},
	}

	objectConfig := ObjectConfig{
		Path:                  thePath,
		BaseTitle:             "object",
		StoreKey:              key,
		ObjectType:            podObjectType,
		DisableResourceViewer: true,
		IconName:              "icon-name",
		IconSource:            "icon-source",
	}
	d := NewObject(objectConfig)

	d.tabFuncDescriptors = []tabFuncDescriptor{
		{name: "summary", tabFunc: d.addSummaryTab},
	}

	cResponse, err := d.Describe(ctx, pod.Namespace, options)
	require.NoError(t, err)

	summary := component.NewText("summary")
	summary.SetAccessor("summary")

	expected := component.ContentResponse{
		Title:      component.Title(component.NewText("object"), component.NewText("pod")),
		IconName:   "icon-name",
		IconSource: "icon-source",
		Components: []component.Component{
			summary,
		},
	}
	assert.Equal(t, expected, cResponse)

}
