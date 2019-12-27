package objectvisitor_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/gvk"
	"github.com/kubenext/kubefun/internal/objectvisitor"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/stretchr/testify/require"
	"testing"

	configFake "github.com/kubenext/kubefun/internal/config/fake"
	ovFake "github.com/kubenext/kubefun/internal/objectvisitor/fake"
	queryerFake "github.com/kubenext/kubefun/internal/queryer/fake"
)

func TestDefaultVisitor_Visit_use_typed_visitor(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	dashConfig := configFake.NewMockDash(controller)

	pod := testutil.CreatePod("pod")
	unstructuredPod := testutil.ToUnstructured(t, pod)

	q := queryerFake.NewMockQueryer(controller)

	handler := ovFake.NewMockObjectHandler(controller)

	defaultHandler := ovFake.NewMockDefaultTypedVisitor(controller)
	defaultHandler.EXPECT().
		Visit(gomock.Any(), unstructuredPod, handler, gomock.Any(), true).Return(nil)

	tv := ovFake.NewMockTypedVisitor(controller)
	tv.EXPECT().Supports().Return(gvk.Pod).AnyTimes()
	tv.EXPECT().
		Visit(gomock.Any(), unstructuredPod, handler, gomock.Any(), true)
	tvList := []objectvisitor.TypedVisitor{tv}

	dv, err := objectvisitor.NewDefaultVisitor(dashConfig, q,
		objectvisitor.SetDefaultHandler(defaultHandler),
		objectvisitor.SetTypedVisitors(tvList))
	require.NoError(t, err)

	ctx := context.Background()
	err = dv.Visit(ctx, testutil.ToUnstructured(t, pod), handler, true)
	require.NoError(t, err)
}
