package yamlviewer

import (
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"testing"
)

func Test_ToComponent(t *testing.T) {
	object := &corev1.Pod{}

	got, err := ToComponent(object)
	require.NoError(t, err)

	data := "---\nmetadata:\n  creationTimestamp: null\nspec:\n  containers: null\nstatus: {}\n"
	expected := component.NewYAML(component.TitleFromString("YAML"), data)

	assert.Equal(t, expected, got)
}
