package kubefun

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/kubenext/kubefun/pkg/store"
	"github.com/kubenext/kubefun/pkg/store/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"testing"

	actionFake "github.com/kubenext/kubefun/pkg/action/fake"
)

func TestServiceConfigurationEditor(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := testutil.CreateService("service")
	service.Namespace = "default"

	objectStore := fake.NewMockStore(controller)
	alerter := actionFake.NewMockAlerter(controller)

	key, err := store.KeyFromObject(service)
	require.NoError(t, err)

	updatedService := service.DeepCopy()
	updatedService.Spec.Selector = map[string]string{
		"foo": "bar",
	}

	objectStore.EXPECT().
		Update(gomock.Any(), key, gomock.Any()).
		DoAndReturn(func(ctx context.Context, key store.Key, fn func(object *unstructured.Unstructured) error) error {
			return nil
		})

	alerter.EXPECT().
		SendAlert(gomock.Any()).
		DoAndReturn(func(alert action.Alert) {
			assert.Equal(t, action.AlertTypeInfo, alert.Type)
			assert.Equal(t, `Updated Service "service"`, alert.Message)
			assert.NotNil(t, alert.Expiration)
		})

	configurationEditor := NewServiceConfigurationEditor(objectStore)
	assert.Equal(t, "overview/serviceEditor", configurationEditor.ActionName())

	ctx := context.Background()

	payload := action.Payload{
		"apiVersion": "v1",
		"kind":       "Service",
		"namespace":  "default",
		"name":       "service",
		"selectors": []interface{}{
			"foo:bar",
		},
	}

	require.NoError(t, configurationEditor.Handle(ctx, alerter, payload))
}
