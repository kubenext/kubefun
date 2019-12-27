// Code generated by MockGen. DO NOT EDIT.
// Source: dashboard/dashboard.pb.go

// Package fake is a generated GoMock package.
package fake

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	dashboard "github.com/kubenext/kubefun/pkg/plugin/dashboard"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockPluginClient is a mock of PluginClient interface
type MockPluginClient struct {
	ctrl     *gomock.Controller
	recorder *MockPluginClientMockRecorder
}

// MockPluginClientMockRecorder is the mock recorder for MockPluginClient
type MockPluginClientMockRecorder struct {
	mock *MockPluginClient
}

// NewMockPluginClient creates a new mock instance
func NewMockPluginClient(ctrl *gomock.Controller) *MockPluginClient {
	mock := &MockPluginClient{ctrl: ctrl}
	mock.recorder = &MockPluginClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPluginClient) EXPECT() *MockPluginClientMockRecorder {
	return m.recorder
}

// Content mocks base method
func (m *MockPluginClient) Content(ctx context.Context, in *dashboard.ContentRequest, opts ...grpc.CallOption) (*dashboard.ContentResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Content", varargs...)
	ret0, _ := ret[0].(*dashboard.ContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Content indicates an expected call of Content
func (mr *MockPluginClientMockRecorder) Content(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockPluginClient)(nil).Content), varargs...)
}

// HandleAction mocks base method
func (m *MockPluginClient) HandleAction(ctx context.Context, in *dashboard.HandleActionRequest, opts ...grpc.CallOption) (*dashboard.HandleActionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HandleAction", varargs...)
	ret0, _ := ret[0].(*dashboard.HandleActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleAction indicates an expected call of HandleAction
func (mr *MockPluginClientMockRecorder) HandleAction(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockPluginClient)(nil).HandleAction), varargs...)
}

// Navigation mocks base method
func (m *MockPluginClient) Navigation(ctx context.Context, in *dashboard.NavigationRequest, opts ...grpc.CallOption) (*dashboard.NavigationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Navigation", varargs...)
	ret0, _ := ret[0].(*dashboard.NavigationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Navigation indicates an expected call of Navigation
func (mr *MockPluginClientMockRecorder) Navigation(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Navigation", reflect.TypeOf((*MockPluginClient)(nil).Navigation), varargs...)
}

// Register mocks base method
func (m *MockPluginClient) Register(ctx context.Context, in *dashboard.RegisterRequest, opts ...grpc.CallOption) (*dashboard.RegisterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Register", varargs...)
	ret0, _ := ret[0].(*dashboard.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockPluginClientMockRecorder) Register(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPluginClient)(nil).Register), varargs...)
}

// Print mocks base method
func (m *MockPluginClient) Print(ctx context.Context, in *dashboard.ObjectRequest, opts ...grpc.CallOption) (*dashboard.PrintResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Print", varargs...)
	ret0, _ := ret[0].(*dashboard.PrintResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Print indicates an expected call of Print
func (mr *MockPluginClientMockRecorder) Print(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPluginClient)(nil).Print), varargs...)
}

// ObjectStatus mocks base method
func (m *MockPluginClient) ObjectStatus(ctx context.Context, in *dashboard.ObjectRequest, opts ...grpc.CallOption) (*dashboard.ObjectStatusResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ObjectStatus", varargs...)
	ret0, _ := ret[0].(*dashboard.ObjectStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStatus indicates an expected call of ObjectStatus
func (mr *MockPluginClientMockRecorder) ObjectStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStatus", reflect.TypeOf((*MockPluginClient)(nil).ObjectStatus), varargs...)
}

// PrintTab mocks base method
func (m *MockPluginClient) PrintTab(ctx context.Context, in *dashboard.ObjectRequest, opts ...grpc.CallOption) (*dashboard.PrintTabResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PrintTab", varargs...)
	ret0, _ := ret[0].(*dashboard.PrintTabResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrintTab indicates an expected call of PrintTab
func (mr *MockPluginClientMockRecorder) PrintTab(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTab", reflect.TypeOf((*MockPluginClient)(nil).PrintTab), varargs...)
}

// WatchAdd mocks base method
func (m *MockPluginClient) WatchAdd(ctx context.Context, in *dashboard.WatchRequest, opts ...grpc.CallOption) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchAdd", varargs...)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAdd indicates an expected call of WatchAdd
func (mr *MockPluginClientMockRecorder) WatchAdd(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAdd", reflect.TypeOf((*MockPluginClient)(nil).WatchAdd), varargs...)
}

// WatchUpdate mocks base method
func (m *MockPluginClient) WatchUpdate(ctx context.Context, in *dashboard.WatchRequest, opts ...grpc.CallOption) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchUpdate", varargs...)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpdate indicates an expected call of WatchUpdate
func (mr *MockPluginClientMockRecorder) WatchUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpdate", reflect.TypeOf((*MockPluginClient)(nil).WatchUpdate), varargs...)
}

// WatchDelete mocks base method
func (m *MockPluginClient) WatchDelete(ctx context.Context, in *dashboard.WatchRequest, opts ...grpc.CallOption) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WatchDelete", varargs...)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchDelete indicates an expected call of WatchDelete
func (mr *MockPluginClientMockRecorder) WatchDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchDelete", reflect.TypeOf((*MockPluginClient)(nil).WatchDelete), varargs...)
}

// MockPluginServer is a mock of PluginServer interface
type MockPluginServer struct {
	ctrl     *gomock.Controller
	recorder *MockPluginServerMockRecorder
}

// MockPluginServerMockRecorder is the mock recorder for MockPluginServer
type MockPluginServerMockRecorder struct {
	mock *MockPluginServer
}

// NewMockPluginServer creates a new mock instance
func NewMockPluginServer(ctrl *gomock.Controller) *MockPluginServer {
	mock := &MockPluginServer{ctrl: ctrl}
	mock.recorder = &MockPluginServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPluginServer) EXPECT() *MockPluginServerMockRecorder {
	return m.recorder
}

// Content mocks base method
func (m *MockPluginServer) Content(arg0 context.Context, arg1 *dashboard.ContentRequest) (*dashboard.ContentResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Content", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.ContentResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Content indicates an expected call of Content
func (mr *MockPluginServerMockRecorder) Content(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Content", reflect.TypeOf((*MockPluginServer)(nil).Content), arg0, arg1)
}

// HandleAction mocks base method
func (m *MockPluginServer) HandleAction(arg0 context.Context, arg1 *dashboard.HandleActionRequest) (*dashboard.HandleActionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleAction", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.HandleActionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HandleAction indicates an expected call of HandleAction
func (mr *MockPluginServerMockRecorder) HandleAction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleAction", reflect.TypeOf((*MockPluginServer)(nil).HandleAction), arg0, arg1)
}

// Navigation mocks base method
func (m *MockPluginServer) Navigation(arg0 context.Context, arg1 *dashboard.NavigationRequest) (*dashboard.NavigationResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Navigation", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.NavigationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Navigation indicates an expected call of Navigation
func (mr *MockPluginServerMockRecorder) Navigation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Navigation", reflect.TypeOf((*MockPluginServer)(nil).Navigation), arg0, arg1)
}

// Register mocks base method
func (m *MockPluginServer) Register(arg0 context.Context, arg1 *dashboard.RegisterRequest) (*dashboard.RegisterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.RegisterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register
func (mr *MockPluginServerMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockPluginServer)(nil).Register), arg0, arg1)
}

// Print mocks base method
func (m *MockPluginServer) Print(arg0 context.Context, arg1 *dashboard.ObjectRequest) (*dashboard.PrintResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.PrintResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Print indicates an expected call of Print
func (mr *MockPluginServerMockRecorder) Print(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockPluginServer)(nil).Print), arg0, arg1)
}

// ObjectStatus mocks base method
func (m *MockPluginServer) ObjectStatus(arg0 context.Context, arg1 *dashboard.ObjectRequest) (*dashboard.ObjectStatusResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObjectStatus", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.ObjectStatusResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ObjectStatus indicates an expected call of ObjectStatus
func (mr *MockPluginServerMockRecorder) ObjectStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObjectStatus", reflect.TypeOf((*MockPluginServer)(nil).ObjectStatus), arg0, arg1)
}

// PrintTab mocks base method
func (m *MockPluginServer) PrintTab(arg0 context.Context, arg1 *dashboard.ObjectRequest) (*dashboard.PrintTabResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PrintTab", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.PrintTabResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PrintTab indicates an expected call of PrintTab
func (mr *MockPluginServerMockRecorder) PrintTab(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PrintTab", reflect.TypeOf((*MockPluginServer)(nil).PrintTab), arg0, arg1)
}

// WatchAdd mocks base method
func (m *MockPluginServer) WatchAdd(arg0 context.Context, arg1 *dashboard.WatchRequest) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchAdd", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchAdd indicates an expected call of WatchAdd
func (mr *MockPluginServerMockRecorder) WatchAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchAdd", reflect.TypeOf((*MockPluginServer)(nil).WatchAdd), arg0, arg1)
}

// WatchUpdate mocks base method
func (m *MockPluginServer) WatchUpdate(arg0 context.Context, arg1 *dashboard.WatchRequest) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchUpdate", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchUpdate indicates an expected call of WatchUpdate
func (mr *MockPluginServerMockRecorder) WatchUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchUpdate", reflect.TypeOf((*MockPluginServer)(nil).WatchUpdate), arg0, arg1)
}

// WatchDelete mocks base method
func (m *MockPluginServer) WatchDelete(arg0 context.Context, arg1 *dashboard.WatchRequest) (*dashboard.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WatchDelete", arg0, arg1)
	ret0, _ := ret[0].(*dashboard.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchDelete indicates an expected call of WatchDelete
func (mr *MockPluginServerMockRecorder) WatchDelete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchDelete", reflect.TypeOf((*MockPluginServer)(nil).WatchDelete), arg0, arg1)
}
