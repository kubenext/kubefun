package printer

import (
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/testutil"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/kubenext/kubefun/pkg/view/flexlayout"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Metadata(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	tpo := newTestPrinterOptions(controller)

	fl := flexlayout.New()

	deployment := testutil.CreateDeployment("deployment")
	metadata, err := NewMetadata(deployment, tpo.link)
	require.NoError(t, err)

	require.NoError(t, metadata.AddToFlexLayout(fl))

	got := fl.ToComponent("Summary")

	expected := component.NewFlexLayout("Summary")
	expected.AddSections([]component.FlexLayoutSection{
		{
			{
				Width: component.WidthFull,
				View: component.NewSummary("Metadata", component.SummarySections{
					{
						Header:  "Age",
						Content: component.NewTimestamp(deployment.CreationTimestamp.Time),
					},
				}...),
			},
		},
	}...)

	assert.Equal(t, expected, got)
}
