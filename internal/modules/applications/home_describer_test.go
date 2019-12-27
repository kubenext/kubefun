package applications_test

import (
	"context"
	"github.com/golang/mock/gomock"
	configFake "github.com/kubenext/kubefun/internal/config/fake"
	"github.com/kubenext/kubefun/internal/describer"
	"github.com/kubenext/kubefun/internal/modules/applications"
	"github.com/kubenext/kubefun/internal/modules/applications/fake"
	"github.com/kubenext/kubefun/pkg/view/component"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_homeDescriber_Describe(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	table := component.NewTable("table", "table", component.NewTableCols("col"))

	s := fake.NewMockSummarizer(controller)
	s.EXPECT().
		Summarize(gomock.Any(), "default", gomock.Any()).
		Return(table, nil)

	dashConfig := configFake.NewMockDash(controller)

	d := applications.NewHomeDescriber(applications.WithHomeDescriberSummarizer(s))

	ctx := context.Background()
	options := describer.Options{
		Dash: dashConfig,
	}
	actual, err := d.Describe(ctx, "default", options)
	require.NoError(t, err)

	expected := component.ContentResponse{
		Title:      component.TitleFromString("Applications"),
		Components: []component.Component{table},
		IconName:   "",
		IconSource: "",
	}
	require.Equal(t, expected, actual)
}
