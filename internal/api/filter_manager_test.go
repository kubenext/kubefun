package api_test

import (
	"github.com/golang/mock/gomock"
	"github.com/kubenext/kubefun/internal/api"
	"github.com/kubenext/kubefun/internal/kubefun"
	kubefunFake "github.com/kubenext/kubefun/internal/kubefun/fake"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/labels"
	"reflect"
	"sort"
	"testing"
)

func TestFilterManager_Handlers(t *testing.T) {
	manager := api.NewFilterManager()

	handlers := manager.Handlers()
	var got []string
	for _, h := range handlers {
		got = append(got, h.RequestType)
	}
	sort.Strings(got)

	expected := []string{
		api.RequestClearFilters,
		api.RequestAddFilter,
		api.RequestRemoveFilter,
	}
	sort.Strings(expected)

	assert.Equal(t, expected, got)
}

func TestFilterManager_AddFilter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().AddFilter(kubefun.Filter{Key: "foo", Value: "bar"})
	state.EXPECT().SendAlert(gomock.Any())

	manager := api.NewFilterManager()

	payload := action.Payload{
		"filter": map[string]interface{}{
			"key":   "foo",
			"value": "bar",
		},
	}

	require.NoError(t, manager.AddFilter(state, payload))
}

func TestFilterManager_ClearFilters(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().SetFilters([]kubefun.Filter{})
	state.EXPECT().SendAlert(gomock.Any())

	manager := api.NewFilterManager()

	payload := action.Payload{}
	require.NoError(t, manager.ClearFilters(state, payload))
}

func TestFilterManager_RemoveFilter(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	state := kubefunFake.NewMockState(controller)
	state.EXPECT().RemoveFilter(kubefun.Filter{Key: "foo", Value: "bar"})
	state.EXPECT().SendAlert(gomock.Any())

	manager := api.NewFilterManager()

	payload := action.Payload{
		"filter": map[string]interface{}{
			"key":   "foo",
			"value": "bar",
		},
	}

	require.NoError(t, manager.RemoveFilter(state, payload))
}

func TestFilterFromPayload(t *testing.T) {
	type args struct {
		in action.Payload
	}
	tests := []struct {
		name    string
		args    args
		want    kubefun.Filter
		isFound bool
	}{
		{
			name: "in general",
			args: args{
				in: action.Payload{
					"filter": map[string]interface{}{
						"key": "foo", "value": "bar",
					},
				},
			},
			want:    kubefun.Filter{Key: "foo", Value: "bar"},
			isFound: true,
		},
		{
			name: "missing filter block",
			args: args{
				in: action.Payload{},
			},
			want:    kubefun.Filter{},
			isFound: false,
		},
		{
			name: "missing value",
			args: args{
				in: action.Payload{
					"filter": map[string]interface{}{
						"key": "foo",
					},
				},
			},
			want:    kubefun.Filter{},
			isFound: false,
		},
		{
			name: "missing key",
			args: args{
				in: action.Payload{
					"filter": map[string]interface{}{
						"value": "bar",
					},
				},
			},
			want:    kubefun.Filter{},
			isFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, isFound := api.FilterFromPayload(tt.args.in)
			require.Equal(t, tt.isFound, isFound)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestFiltersFromQueryParams(t *testing.T) {
	type args struct {
		in interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []kubefun.Filter
		wantErr bool
	}{
		{
			name: "single filter",
			args: args{in: "foo:bar"},
			want: []kubefun.Filter{{
				Key:   "foo",
				Value: "bar",
			}},
			wantErr: false,
		},
		{
			name: "multiple filters",
			args: args{in: []interface{}{"foo:bar", "baz:qux"}},
			want: []kubefun.Filter{
				{Key: "foo", Value: "bar"},
				{Key: "baz", Value: "qux"},
			},
			wantErr: false,
		},
		{
			name:    "unknown input",
			args:    args{in: 1},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.FiltersFromQueryParams(tt.args.in)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestParseFilterQueryParam(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    kubefun.Filter
		wantErr bool
	}{
		{
			name: "valid",
			args: args{"foo:bar"},
			want: kubefun.Filter{
				Key:   "foo",
				Value: "bar",
			},
			wantErr: false,
		},
		{
			name:    "invalid",
			args:    args{"foobar"},
			want:    kubefun.Filter{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.ParseFilterQueryParam(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFilterQueryParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFilterQueryParam() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFiltersToLabelSet(t *testing.T) {
	type args struct {
		filters []kubefun.Filter
	}
	tests := []struct {
		name string
		args args
		want *labels.Set
	}{
		{
			name: "in general",
			args: args{
				filters: []kubefun.Filter{
					{Key: "foo", Value: "bar"},
					{Key: "baz", Value: "qux"},
				},
			},
			want: &labels.Set{
				"foo": "bar",
				"baz": "qux",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.FiltersToLabelSet(tt.args.filters); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FiltersToLabelSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
