package resourceviewer

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/objectstatus"
	"github.com/kubenext/kubefun/internal/resourceviewer/fake"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_podGroupNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	pod := testutil.ToUnstructured(t, testutil.CreatePod("pod"))
	objectStatus := fake.NewMockObjectStatus(controller)
	objectStatus.EXPECT().
		Status(gomock.Any(), pod).
		Return(&objectstatus.ObjectStatus{}, nil)

	pgn := podGroupNode{objectStatus: objectStatus}

	objects := testutil.ToUnstructuredList(t, pod)
	name := "foo pods"

	ctx := context.Background()

	got, err := pgn.Create(ctx, name, objects.Items)
	require.NoError(t, err)

	podStatus := component.NewPodStatus()
	podStatus.AddSummary(pod.GetName(), nil, component.NodeStatusOK)

	expected := &component.Node{
		Name:       "foo pods",
		APIVersion: "v1",
		Kind:       "Pod",
		Status:     component.NodeStatusOK,
		Details:    []component.Component{podStatus},
	}

	testutil.AssertJSONEqual(t, expected, got)
}
