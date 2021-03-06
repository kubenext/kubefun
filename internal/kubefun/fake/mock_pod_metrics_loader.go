// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubenext/kubefun/internal/kubefun (interfaces: PodMetricsLoader)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	reflect "reflect"
)

// MockPodMetricsLoader is a mock of PodMetricsLoader interface
type MockPodMetricsLoader struct {
	ctrl     *gomock.Controller
	recorder *MockPodMetricsLoaderMockRecorder
}

// MockPodMetricsLoaderMockRecorder is the mock recorder for MockPodMetricsLoader
type MockPodMetricsLoaderMockRecorder struct {
	mock *MockPodMetricsLoader
}

// NewMockPodMetricsLoader creates a new mock instance
func NewMockPodMetricsLoader(ctrl *gomock.Controller) *MockPodMetricsLoader {
	mock := &MockPodMetricsLoader{ctrl: ctrl}
	mock.recorder = &MockPodMetricsLoaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPodMetricsLoader) EXPECT() *MockPodMetricsLoaderMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockPodMetricsLoader) Load(arg0, arg1 string) (*unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockPodMetricsLoaderMockRecorder) Load(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockPodMetricsLoader)(nil).Load), arg0, arg1)
}

// SupportsMetrics mocks base method
func (m *MockPodMetricsLoader) SupportsMetrics() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SupportsMetrics")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SupportsMetrics indicates an expected call of SupportsMetrics
func (mr *MockPodMetricsLoaderMockRecorder) SupportsMetrics() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportsMetrics", reflect.TypeOf((*MockPodMetricsLoader)(nil).SupportsMetrics))
}
