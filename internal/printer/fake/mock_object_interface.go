// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubenext/kubefun/internal/printer (interfaces: ObjectInterface)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	action "github.com/kubenext/kubefun/pkg/action"
	component "github.com/kubenext/kubefun/pkg/view/component"
	reflect "reflect"
)

// MockObjectInterface is a mock of ObjectInterface interface
type MockObjectInterface struct {
	ctrl     *gomock.Controller
	recorder *MockObjectInterfaceMockRecorder
}

// MockObjectInterfaceMockRecorder is the mock recorder for MockObjectInterface
type MockObjectInterfaceMockRecorder struct {
	mock *MockObjectInterface
}

// NewMockObjectInterface creates a new mock instance
func NewMockObjectInterface(ctrl *gomock.Controller) *MockObjectInterface {
	mock := &MockObjectInterface{ctrl: ctrl}
	mock.recorder = &MockObjectInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockObjectInterface) EXPECT() *MockObjectInterfaceMockRecorder {
	return m.recorder
}

// AddButton mocks base method
func (m *MockObjectInterface) AddButton(arg0 string, arg1 action.Payload, arg2 ...component.ButtonOption) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddButton", varargs...)
}

// AddButton indicates an expected call of AddButton
func (mr *MockObjectInterfaceMockRecorder) AddButton(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddButton", reflect.TypeOf((*MockObjectInterface)(nil).AddButton), varargs...)
}
