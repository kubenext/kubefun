package describer

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	configFake "github.com/kubenext/kubefun/internal/config/fake"
)

func TestSectionDescriber(t *testing.T) {
	namespace := "default"

	controller := gomock.NewController(t)
	defer controller.Finish()

	dashConfig := configFake.NewMockDash(controller)

	options := Options{
		Dash: dashConfig,
	}

	ctx := context.Background()

	tests := []struct {
		name     string
		d        *Section
		expected component.ContentResponse
	}{
		{
			name: "general",
			d: NewSection(
				"/section",
				"section",
				NewStubDescriber("/foo"),
			),
			expected: component.ContentResponse{
				Title: component.Title(component.NewText("section")),
				Components: []component.Component{
					component.NewList("section", nil),
				},
			},
		},
		{
			name: "empty",
			d: NewSection(
				"/section",
				"section",
				NewEmptyDescriber("/foo"),
				NewEmptyDescriber("/bar"),
			),
			expected: component.ContentResponse{
				Title: component.Title(component.NewText("section")),
				Components: []component.Component{
					component.NewList("section", nil),
				},
			},
		},
		{
			name: "empty component",
			d: NewSection(
				"/section",
				"section",
				NewStubDescriber("/foo", &emptyComponent{}),
			),
			expected: component.ContentResponse{
				Title: component.Title(component.NewText("section")),
				Components: []component.Component{
					component.NewList("section", nil),
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			got, err := tc.d.Describe(ctx, namespace, options)
			require.NoError(t, err)

			assert.Equal(t, tc.expected, got)
		})
	}

}
