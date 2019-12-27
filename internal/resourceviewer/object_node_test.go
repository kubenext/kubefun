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

	linkFake "github.com/kubenext/kubefun/internal/link/fake"
	pluginFake "github.com/kubenext/kubefun/pkg/plugin/fake"
)

func Test_objectNode(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	deployment := testutil.ToUnstructured(t, testutil.CreateDeployment("deployment"))
	deploymentLink := component.NewLink("", deployment.GetName(), "/deployment")

	l := linkFake.NewMockInterface(controller)
	l.EXPECT().
		ForObjectWithQuery(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(deploymentLink, nil)

	pluginPrinter := pluginFake.NewMockManagerInterface(controller)
	objectStatus := fake.NewMockObjectStatus(controller)
	objectStatus.EXPECT().
		Status(gomock.Any(), gomock.Any()).
		Return(&objectstatus.ObjectStatus{}, nil)

	on := objectNode{
		link:          l,
		pluginPrinter: pluginPrinter,
		objectStatus:  objectStatus,
	}

	ctx := context.Background()

	got, err := on.Create(ctx, deployment)
	require.NoError(t, err)

	expected := &component.Node{
		Name:       deployment.GetName(),
		APIVersion: deployment.GetAPIVersion(),
		Kind:       deployment.GetKind(),
		Status:     component.NodeStatusOK,
		Path:       deploymentLink,
	}

	testutil.AssertJSONEqual(t, expected, got)
}
