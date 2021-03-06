// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubenext/kubefun/internal/objectstore (interfaces: InformerFactory)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers"
	reflect "reflect"
)

// MockInformerFactory is a mock of InformerFactory interface
type MockInformerFactory struct {
	ctrl     *gomock.Controller
	recorder *MockInformerFactoryMockRecorder
}

// MockInformerFactoryMockRecorder is the mock recorder for MockInformerFactory
type MockInformerFactoryMockRecorder struct {
	mock *MockInformerFactory
}

// NewMockInformerFactory creates a new mock instance
func NewMockInformerFactory(ctrl *gomock.Controller) *MockInformerFactory {
	mock := &MockInformerFactory{ctrl: ctrl}
	mock.recorder = &MockInformerFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInformerFactory) EXPECT() *MockInformerFactoryMockRecorder {
	return m.recorder
}

// Delete mocks base method
func (m *MockInformerFactory) Delete(arg0 schema.GroupVersionResource) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete
func (mr *MockInformerFactoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockInformerFactory)(nil).Delete), arg0)
}

// ForResource mocks base method
func (m *MockInformerFactory) ForResource(arg0 schema.GroupVersionResource) informers.GenericInformer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForResource", arg0)
	ret0, _ := ret[0].(informers.GenericInformer)
	return ret0
}

// ForResource indicates an expected call of ForResource
func (mr *MockInformerFactoryMockRecorder) ForResource(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForResource", reflect.TypeOf((*MockInformerFactory)(nil).ForResource), arg0)
}

// WaitForCacheSync mocks base method
func (m *MockInformerFactory) WaitForCacheSync(arg0 <-chan struct{}) map[schema.GroupVersionResource]bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitForCacheSync", arg0)
	ret0, _ := ret[0].(map[schema.GroupVersionResource]bool)
	return ret0
}

// WaitForCacheSync indicates an expected call of WaitForCacheSync
func (mr *MockInformerFactoryMockRecorder) WaitForCacheSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitForCacheSync", reflect.TypeOf((*MockInformerFactory)(nil).WaitForCacheSync), arg0)
}
