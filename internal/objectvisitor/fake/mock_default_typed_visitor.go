// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubenext/kubefun/internal/objectvisitor (interfaces: DefaultTypedVisitor)

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	objectvisitor "github.com/kubenext/kubefun/internal/objectvisitor"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	reflect "reflect"
)

// MockDefaultTypedVisitor is a mock of DefaultTypedVisitor interface
type MockDefaultTypedVisitor struct {
	ctrl     *gomock.Controller
	recorder *MockDefaultTypedVisitorMockRecorder
}

// MockDefaultTypedVisitorMockRecorder is the mock recorder for MockDefaultTypedVisitor
type MockDefaultTypedVisitorMockRecorder struct {
	mock *MockDefaultTypedVisitor
}

// NewMockDefaultTypedVisitor creates a new mock instance
func NewMockDefaultTypedVisitor(ctrl *gomock.Controller) *MockDefaultTypedVisitor {
	mock := &MockDefaultTypedVisitor{ctrl: ctrl}
	mock.recorder = &MockDefaultTypedVisitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDefaultTypedVisitor) EXPECT() *MockDefaultTypedVisitorMockRecorder {
	return m.recorder
}

// Visit mocks base method
func (m *MockDefaultTypedVisitor) Visit(arg0 context.Context, arg1 *unstructured.Unstructured, arg2 objectvisitor.ObjectHandler, arg3 objectvisitor.Visitor, arg4 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Visit", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Visit indicates an expected call of Visit
func (mr *MockDefaultTypedVisitorMockRecorder) Visit(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Visit", reflect.TypeOf((*MockDefaultTypedVisitor)(nil).Visit), arg0, arg1, arg2, arg3, arg4)
}
