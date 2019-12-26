package objectstatus

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"testing"

	storefake "github.com/kubenext/kubefun/pkg/store/fake"
)

func Test_pod(t *testing.T) {
	cases := []struct {
		name     string
		init     func(*testing.T) runtime.Object
		expected ObjectStatus
		isErr    bool
	}{
		{
			name: "in general",
			init: func(t *testing.T) runtime.Object {
				objectFile := "pod_ok.yaml"
				return testutil.LoadObjectFromFile(t, objectFile)
			},
			expected: ObjectStatus{
				nodeStatus: component.NodeStatusOK,
				Details: []component.Component{
					component.NewText(""),
				},
			},
		},
		{
			name: "pod is in unknown state",
			init: func(t *testing.T) runtime.Object {
				objectFile := "pod_unknown.yaml"
				return testutil.LoadObjectFromFile(t, objectFile)
			},
			expected: ObjectStatus{
				nodeStatus: component.NodeStatusError,
				Details: []component.Component{
					component.NewText(""),
				},
			},
		},
		{
			name: "pod is pending",
			init: func(t *testing.T) runtime.Object {
				objectFile := "pod_pending.yaml"
				return testutil.LoadObjectFromFile(t, objectFile)
			},
			expected: ObjectStatus{
				nodeStatus: component.NodeStatusWarning,
				Details: []component.Component{
					component.NewText(""),
				},
			},
		},
		{
			name: "object is nil",
			init: func(t *testing.T) runtime.Object {
				return nil
			},
			isErr: true,
		},
		{
			name: "object is not a pod",
			init: func(t *testing.T) runtime.Object {
				return &unstructured.Unstructured{}
			},
			isErr: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			controller := gomock.NewController(t)
			defer controller.Finish()

			o := storefake.NewMockStore(controller)

			object := tc.init(t)

			ctx := context.Background()
			status, err := pod(ctx, object, o)
			if tc.isErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tc.expected, status)
		})
	}
}
