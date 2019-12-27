package api

import (
	"context"
	"fmt"
	"github.com/kubenext/kubefun/internal/kubefun"
	"github.com/kubenext/kubefun/pkg/action"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"strings"
)

const (
	RequestAddFilter    = "addFilter"
	RequestClearFilters = "clearFilters"
	RequestRemoveFilter = "removeFilter"
)

// FilterManager manages filters.
type FilterManager struct {
}

var _ StateManager = (*FilterManager)(nil)

// NewFilterManager creates an instance of FilterManager.
func NewFilterManager() *FilterManager {
	return &FilterManager{}
}

// Start starts the manager. Current is a no-op.
func (fm *FilterManager) Start(ctx context.Context, state kubefun.State, s KubefunClient) {
}

// Handlers returns a slice of handlers.
func (fm *FilterManager) Handlers() []kubefun.ClientRequestHandler {
	return []kubefun.ClientRequestHandler{
		{
			RequestType: RequestAddFilter,
			Handler:     fm.AddFilter,
		},
		{
			RequestType: RequestClearFilters,
			Handler:     fm.ClearFilters,
		},
		{
			RequestType: RequestRemoveFilter,
			Handler:     fm.RemoveFilter,
		},
	}
}

// AddFilter adds a filter.
func (fm *FilterManager) AddFilter(state kubefun.State, payload action.Payload) error {
	if filter, ok := FilterFromPayload(payload); ok {
		state.AddFilter(filter)
		message := fmt.Sprintf("Added filter for label %s", filter.String())
		state.SendAlert(action.CreateAlert(action.AlertTypeInfo, message, action.DefaultAlertExpiration))
	}

	return nil
}

// ClearFilters clears all filters.
func (fm *FilterManager) ClearFilters(state kubefun.State, payload action.Payload) error {
	state.SetFilters([]kubefun.Filter{})
	message := "Cleared filters"
	state.SendAlert(action.CreateAlert(action.AlertTypeInfo, message, action.DefaultAlertExpiration))
	return nil
}

// RemoveFilters removes a filter.
func (fm *FilterManager) RemoveFilter(state kubefun.State, payload action.Payload) error {
	if filter, ok := FilterFromPayload(payload); ok {
		state.RemoveFilter(filter)
		message := fmt.Sprintf("Removed filter for label %s", filter.String())
		state.SendAlert(action.CreateAlert(action.AlertTypeInfo, message, action.DefaultAlertExpiration))
	}
	return nil
}

// FilterFromPayload creates a filter from a payload. Returns false
// if the payload is invalid.
func FilterFromPayload(in action.Payload) (kubefun.Filter, bool) {
	filters, found, err := unstructured.NestedMap(in, "filter")
	if err != nil || !found {
		return kubefun.Filter{}, false
	}

	key, found, err := unstructured.NestedString(filters, "key")
	if err != nil || !found {
		return kubefun.Filter{}, false
	}

	value, found, err := unstructured.NestedString(filters, "value")
	if err != nil || !found {
		return kubefun.Filter{}, false
	}

	return kubefun.Filter{
		Key:   key,
		Value: value,
	}, true
}

// FiltersFromQueryParams converts query params to filters. Can handle
// one or multiple query params.
func FiltersFromQueryParams(in interface{}) ([]kubefun.Filter, error) {
	var filters []kubefun.Filter

	switch t := in.(type) {
	case []interface{}:
		for i := range t {
			if raw, ok := t[i].(string); ok {
				filter, err := ParseFilterQueryParam(raw)
				if err != nil {
					return nil, err
				}
				filters = append(filters, filter)
			}
		}
	case string:
		filter, err := ParseFilterQueryParam(t)
		if err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	default:
		return nil, errors.Errorf("not sure what to do with filter of type %T\n", in)
	}

	return filters, nil
}

// ParseFilterQueryParam parsers a single filter from a query param in the format `key:value`.
func ParseFilterQueryParam(in string) (kubefun.Filter, error) {
	parts := strings.Split(in, ":")
	if len(parts) != 2 {
		return kubefun.Filter{}, errors.Errorf("invalid filter parameter %s", in)
	}

	return kubefun.Filter{
		Key:   parts[0],
		Value: parts[1],
	}, nil
}

// FiltersToLabelSet converts a slice of filters to a label set.
func FiltersToLabelSet(filters []kubefun.Filter) *labels.Set {
	set := labels.Set{}
	for i := range filters {
		set[filters[i].Key] = filters[i].Value
	}
	return &set

}
