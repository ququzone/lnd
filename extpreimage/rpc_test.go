// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/lightningnetwork/lnd/extpreimage (interfaces: ExternalPreimageServiceClient,ExternalPreimageService_GetPreimageClient)

// Package extpreimage_test is a generated GoMock package.
package extpreimage_test

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	extpreimage "github.com/lightningnetwork/lnd/extpreimage"
	grpc "google.golang.org/grpc"
	metadata "google.golang.org/grpc/metadata"
	reflect "reflect"
)

// MockExternalPreimageServiceClient is a mock of ExternalPreimageServiceClient interface
type MockExternalPreimageServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockExternalPreimageServiceClientMockRecorder
}

// MockExternalPreimageServiceClientMockRecorder is the mock recorder for MockExternalPreimageServiceClient
type MockExternalPreimageServiceClientMockRecorder struct {
	mock *MockExternalPreimageServiceClient
}

// NewMockExternalPreimageServiceClient creates a new mock instance
func NewMockExternalPreimageServiceClient(ctrl *gomock.Controller) *MockExternalPreimageServiceClient {
	mock := &MockExternalPreimageServiceClient{ctrl: ctrl}
	mock.recorder = &MockExternalPreimageServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalPreimageServiceClient) EXPECT() *MockExternalPreimageServiceClientMockRecorder {
	return m.recorder
}

// GetPreimage mocks base method
func (m *MockExternalPreimageServiceClient) GetPreimage(arg0 context.Context, arg1 *extpreimage.GetPreimageRequest, arg2 ...grpc.CallOption) (extpreimage.ExternalPreimageService_GetPreimageClient, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPreimage", varargs...)
	ret0, _ := ret[0].(extpreimage.ExternalPreimageService_GetPreimageClient)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreimage indicates an expected call of GetPreimage
func (mr *MockExternalPreimageServiceClientMockRecorder) GetPreimage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreimage", reflect.TypeOf((*MockExternalPreimageServiceClient)(nil).GetPreimage), varargs...)
}

// MockExternalPreimageService_GetPreimageClient is a mock of ExternalPreimageService_GetPreimageClient interface
type MockExternalPreimageService_GetPreimageClient struct {
	ctrl     *gomock.Controller
	recorder *MockExternalPreimageService_GetPreimageClientMockRecorder
}

// MockExternalPreimageService_GetPreimageClientMockRecorder is the mock recorder for MockExternalPreimageService_GetPreimageClient
type MockExternalPreimageService_GetPreimageClientMockRecorder struct {
	mock *MockExternalPreimageService_GetPreimageClient
}

// NewMockExternalPreimageService_GetPreimageClient creates a new mock instance
func NewMockExternalPreimageService_GetPreimageClient(ctrl *gomock.Controller) *MockExternalPreimageService_GetPreimageClient {
	mock := &MockExternalPreimageService_GetPreimageClient{ctrl: ctrl}
	mock.recorder = &MockExternalPreimageService_GetPreimageClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExternalPreimageService_GetPreimageClient) EXPECT() *MockExternalPreimageService_GetPreimageClientMockRecorder {
	return m.recorder
}

// CloseSend mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) CloseSend() error {
	ret := m.ctrl.Call(m, "CloseSend")
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseSend indicates an expected call of CloseSend
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) CloseSend() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseSend", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).CloseSend))
}

// Context mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) Context() context.Context {
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) Context() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).Context))
}

// Header mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) Header() (metadata.MD, error) {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(metadata.MD)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Header indicates an expected call of Header
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).Header))
}

// Recv mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) Recv() (*extpreimage.GetPreimageResponse, error) {
	ret := m.ctrl.Call(m, "Recv")
	ret0, _ := ret[0].(*extpreimage.GetPreimageResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Recv indicates an expected call of Recv
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) Recv() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Recv", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).Recv))
}

// RecvMsg mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) RecvMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "RecvMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecvMsg indicates an expected call of RecvMsg
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) RecvMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecvMsg", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).RecvMsg), arg0)
}

// SendMsg mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) SendMsg(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "SendMsg", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMsg indicates an expected call of SendMsg
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) SendMsg(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMsg", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).SendMsg), arg0)
}

// Trailer mocks base method
func (m *MockExternalPreimageService_GetPreimageClient) Trailer() metadata.MD {
	ret := m.ctrl.Call(m, "Trailer")
	ret0, _ := ret[0].(metadata.MD)
	return ret0
}

// Trailer indicates an expected call of Trailer
func (mr *MockExternalPreimageService_GetPreimageClientMockRecorder) Trailer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Trailer", reflect.TypeOf((*MockExternalPreimageService_GetPreimageClient)(nil).Trailer))
}