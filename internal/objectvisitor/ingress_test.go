package objectvisitor_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/objectvisitor"
	"github.com/kubenext/kubefun/internal/objectvisitor/fake"
	"github.com/kubenext/kubefun/internal/testutil"
	"testing"

	queryerFake "github.com/kubenext/kubefun/internal/queryer/fake"
	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func TestIngress_Visit(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	object := testutil.CreateIngress("ingress")
	u := testutil.ToUnstructured(t, object)

	q := queryerFake.NewMockQueryer(controller)
	service := testutil.CreateService("service")
	q.EXPECT().
		ServicesForIngress(gomock.Any(), object).
		Return(testutil.ToUnstructuredList(t, service), nil)

	handler := fake.NewMockObjectHandler(controller)
	handler.EXPECT().
		AddEdge(gomock.Any(), u, testutil.ToUnstructured(t, service)).
		Return(nil)

	var visited []unstructured.Unstructured
	visitor := fake.NewMockVisitor(controller)
	visitor.EXPECT().
		Visit(gomock.Any(), gomock.Any(), handler, true).
		DoAndReturn(func(ctx context.Context, object *unstructured.Unstructured, handler objectvisitor.ObjectHandler, _ bool) error {
			visited = append(visited, *object)
			return nil
		})

	ingress := objectvisitor.NewIngress(q)

	ctx := context.Background()
	err := ingress.Visit(ctx, u, handler, visitor, true)

	sortObjectsByName(t, visited)

	expected := testutil.ToUnstructuredList(t, service)
	assert.Equal(t, expected.Items, visited)
	assert.NoError(t, err)
}

func TestIngress_Visit_invalid_service_name(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	object := testutil.CreateIngress("ingress")
	u := testutil.ToUnstructured(t, object)

	q := queryerFake.NewMockQueryer(controller)
	q.EXPECT().
		ServicesForIngress(gomock.Any(), object).
		Return(testutil.ToUnstructuredList(t), nil)

	handler := fake.NewMockObjectHandler(controller)

	visitor := fake.NewMockVisitor(controller)

	ingress := objectvisitor.NewIngress(q)

	ctx := context.Background()
	err := ingress.Visit(ctx, u, handler, visitor, true)

	assert.NoError(t, err)

}
