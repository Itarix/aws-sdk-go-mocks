// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/kinesisvideomedia/kinesisvideomediaiface (interfaces: KinesisVideoMediaAPI)

// Package kinesisvideomediamock is a generated GoMock package.
package kinesisvideomediamock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	kinesisvideomedia "github.com/aws/aws-sdk-go/service/kinesisvideomedia"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockKinesisVideoMediaAPI is a mock of KinesisVideoMediaAPI interface
type MockKinesisVideoMediaAPI struct {
	ctrl     *gomock.Controller
	recorder *MockKinesisVideoMediaAPIMockRecorder
}

// MockKinesisVideoMediaAPIMockRecorder is the mock recorder for MockKinesisVideoMediaAPI
type MockKinesisVideoMediaAPIMockRecorder struct {
	mock *MockKinesisVideoMediaAPI
}

// NewMockKinesisVideoMediaAPI creates a new mock instance
func NewMockKinesisVideoMediaAPI(ctrl *gomock.Controller) *MockKinesisVideoMediaAPI {
	mock := &MockKinesisVideoMediaAPI{ctrl: ctrl}
	mock.recorder = &MockKinesisVideoMediaAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockKinesisVideoMediaAPI) EXPECT() *MockKinesisVideoMediaAPIMockRecorder {
	return m.recorder
}

// GetMedia mocks base method
func (m *MockKinesisVideoMediaAPI) GetMedia(arg0 *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMedia", arg0)
	ret0, _ := ret[0].(*kinesisvideomedia.GetMediaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMedia indicates an expected call of GetMedia
func (mr *MockKinesisVideoMediaAPIMockRecorder) GetMedia(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMedia", reflect.TypeOf((*MockKinesisVideoMediaAPI)(nil).GetMedia), arg0)
}

// GetMediaRequest mocks base method
func (m *MockKinesisVideoMediaAPI) GetMediaRequest(arg0 *kinesisvideomedia.GetMediaInput) (*request.Request, *kinesisvideomedia.GetMediaOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMediaRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*kinesisvideomedia.GetMediaOutput)
	return ret0, ret1
}

// GetMediaRequest indicates an expected call of GetMediaRequest
func (mr *MockKinesisVideoMediaAPIMockRecorder) GetMediaRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaRequest", reflect.TypeOf((*MockKinesisVideoMediaAPI)(nil).GetMediaRequest), arg0)
}

// GetMediaWithContext mocks base method
func (m *MockKinesisVideoMediaAPI) GetMediaWithContext(arg0 context.Context, arg1 *kinesisvideomedia.GetMediaInput, arg2 ...request.Option) (*kinesisvideomedia.GetMediaOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMediaWithContext", varargs...)
	ret0, _ := ret[0].(*kinesisvideomedia.GetMediaOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMediaWithContext indicates an expected call of GetMediaWithContext
func (mr *MockKinesisVideoMediaAPIMockRecorder) GetMediaWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMediaWithContext", reflect.TypeOf((*MockKinesisVideoMediaAPI)(nil).GetMediaWithContext), varargs...)
}
