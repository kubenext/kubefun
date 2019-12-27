package applications

import (
	"context"
	"github.com/golang/mock/gomock"
	configFake "github.com/kubenext/kubefun/internal/config/fake"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/store/fake"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_summarizer(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	ctx := context.Background()

	objectStore := fake.NewMockStore(controller)

	key := store.Key{
		Namespace:  "default",
		APIVersion: "v1",
		Kind:       "Pod",
	}
	podList := testutil.ToUnstructuredList(t, testutil.CreatePod("pod", withPodLabels(map[string]string{
		appLabelName:     "name",
		appLabelInstance: "instance",
		appLabelVersion:  "version",
	})))
	objectStore.EXPECT().
		List(gomock.Any(), key).
		Return(podList, true, nil)

	dashConfig := configFake.NewMockDash(controller)
	dashConfig.EXPECT().ObjectStore().Return(objectStore)

	s := summarizer{}
	actual, err := s.Summarize(ctx, "default", dashConfig)
	require.NoError(t, err)

	expected := component.NewTableWithRows("Applications", "applications", applicationListColumns, []component.TableRow{
		{
			"Name":     component.NewLink("", "name", "/applications/namespace/default/name/instance/version"),
			"Instance": component.NewText("instance"),
			"Version":  component.NewText("version"),
			"State":    component.NewText("state"),
			"Pods":     component.NewText("1"),
		},
	})

	component.AssertEqual(t, expected, actual)
}
