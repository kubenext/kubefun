// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubenext/kubefun/internal/api (interfaces: KubefunClient)

// Package fake is a generated GoMock package.
package fake

import (
	gomock "github.com/golang/mock/gomock"
	kubefun "github.com/kubenext/kubefun/internal/kubefun"
	reflect "reflect"
)

// MockKubefunClient is a mock of KubefunClient interface
type MockKubefunClient struct {
	ctrl     *gomock.Controller
	recorder *MockKubefunClientMockRecorder
}

// MockKubefunClientMockRecorder is the mock recorder for MockKubefunClient
type MockKubefunClientMockRecorder struct {
	mock *MockKubefunClient
}

// NewMockKubefunClient creates a new mock instance
func NewMockKubefunClient(ctrl *gomock.Controller) *MockKubefunClient {
	mock := &MockKubefunClient{ctrl: ctrl}
	mock.recorder = &MockKubefunClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKubefunClient) EXPECT() *MockKubefunClientMockRecorder {
	return m.recorder
}

// ID mocks base method
func (m *MockKubefunClient) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockKubefunClientMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockKubefunClient)(nil).ID))
}

// Send mocks base method
func (m *MockKubefunClient) Send(arg0 kubefun.Event) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send
func (mr *MockKubefunClientMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockKubefunClient)(nil).Send), arg0)
}