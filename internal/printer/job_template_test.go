package printer

import (
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestJobTemplateHeader(t *testing.T) {
	labels := map[string]string{
		"app": "myapp",
	}

	jth := NewJobTemplateHeader(labels)
	got, err := jth.Create()

	require.NoError(t, err)

	assert.Len(t, got.Config.Labels, 1)

	expected := []component.TitleComponent{
		component.NewText("Job Template"),
	}

	assert.Equal(t, expected, got.Metadata.Title)
}
