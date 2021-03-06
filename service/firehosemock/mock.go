// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/firehose/firehoseiface (interfaces: FirehoseAPI)

// Package firehosemock is a generated GoMock package.
package firehosemock

import (
	context "context"
	request "github.com/aws/aws-sdk-go/aws/request"
	firehose "github.com/aws/aws-sdk-go/service/firehose"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockFirehoseAPI is a mock of FirehoseAPI interface
type MockFirehoseAPI struct {
	ctrl     *gomock.Controller
	recorder *MockFirehoseAPIMockRecorder
}

// MockFirehoseAPIMockRecorder is the mock recorder for MockFirehoseAPI
type MockFirehoseAPIMockRecorder struct {
	mock *MockFirehoseAPI
}

// NewMockFirehoseAPI creates a new mock instance
func NewMockFirehoseAPI(ctrl *gomock.Controller) *MockFirehoseAPI {
	mock := &MockFirehoseAPI{ctrl: ctrl}
	mock.recorder = &MockFirehoseAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFirehoseAPI) EXPECT() *MockFirehoseAPIMockRecorder {
	return m.recorder
}

// CreateDeliveryStream mocks base method
func (m *MockFirehoseAPI) CreateDeliveryStream(arg0 *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.CreateDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeliveryStream indicates an expected call of CreateDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) CreateDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).CreateDeliveryStream), arg0)
}

// CreateDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) CreateDeliveryStreamRequest(arg0 *firehose.CreateDeliveryStreamInput) (*request.Request, *firehose.CreateDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.CreateDeliveryStreamOutput)
	return ret0, ret1
}

// CreateDeliveryStreamRequest indicates an expected call of CreateDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) CreateDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).CreateDeliveryStreamRequest), arg0)
}

// CreateDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) CreateDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.CreateDeliveryStreamInput, arg2 ...request.Option) (*firehose.CreateDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.CreateDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDeliveryStreamWithContext indicates an expected call of CreateDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) CreateDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).CreateDeliveryStreamWithContext), varargs...)
}

// DeleteDeliveryStream mocks base method
func (m *MockFirehoseAPI) DeleteDeliveryStream(arg0 *firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.DeleteDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDeliveryStream indicates an expected call of DeleteDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) DeleteDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).DeleteDeliveryStream), arg0)
}

// DeleteDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) DeleteDeliveryStreamRequest(arg0 *firehose.DeleteDeliveryStreamInput) (*request.Request, *firehose.DeleteDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.DeleteDeliveryStreamOutput)
	return ret0, ret1
}

// DeleteDeliveryStreamRequest indicates an expected call of DeleteDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) DeleteDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).DeleteDeliveryStreamRequest), arg0)
}

// DeleteDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) DeleteDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.DeleteDeliveryStreamInput, arg2 ...request.Option) (*firehose.DeleteDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.DeleteDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDeliveryStreamWithContext indicates an expected call of DeleteDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) DeleteDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).DeleteDeliveryStreamWithContext), varargs...)
}

// DescribeDeliveryStream mocks base method
func (m *MockFirehoseAPI) DescribeDeliveryStream(arg0 *firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.DescribeDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDeliveryStream indicates an expected call of DescribeDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) DescribeDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).DescribeDeliveryStream), arg0)
}

// DescribeDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) DescribeDeliveryStreamRequest(arg0 *firehose.DescribeDeliveryStreamInput) (*request.Request, *firehose.DescribeDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.DescribeDeliveryStreamOutput)
	return ret0, ret1
}

// DescribeDeliveryStreamRequest indicates an expected call of DescribeDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) DescribeDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).DescribeDeliveryStreamRequest), arg0)
}

// DescribeDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) DescribeDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.DescribeDeliveryStreamInput, arg2 ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.DescribeDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDeliveryStreamWithContext indicates an expected call of DescribeDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) DescribeDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).DescribeDeliveryStreamWithContext), varargs...)
}

// ListDeliveryStreams mocks base method
func (m *MockFirehoseAPI) ListDeliveryStreams(arg0 *firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeliveryStreams", arg0)
	ret0, _ := ret[0].(*firehose.ListDeliveryStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeliveryStreams indicates an expected call of ListDeliveryStreams
func (mr *MockFirehoseAPIMockRecorder) ListDeliveryStreams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeliveryStreams", reflect.TypeOf((*MockFirehoseAPI)(nil).ListDeliveryStreams), arg0)
}

// ListDeliveryStreamsRequest mocks base method
func (m *MockFirehoseAPI) ListDeliveryStreamsRequest(arg0 *firehose.ListDeliveryStreamsInput) (*request.Request, *firehose.ListDeliveryStreamsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDeliveryStreamsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.ListDeliveryStreamsOutput)
	return ret0, ret1
}

// ListDeliveryStreamsRequest indicates an expected call of ListDeliveryStreamsRequest
func (mr *MockFirehoseAPIMockRecorder) ListDeliveryStreamsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeliveryStreamsRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).ListDeliveryStreamsRequest), arg0)
}

// ListDeliveryStreamsWithContext mocks base method
func (m *MockFirehoseAPI) ListDeliveryStreamsWithContext(arg0 context.Context, arg1 *firehose.ListDeliveryStreamsInput, arg2 ...request.Option) (*firehose.ListDeliveryStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListDeliveryStreamsWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.ListDeliveryStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDeliveryStreamsWithContext indicates an expected call of ListDeliveryStreamsWithContext
func (mr *MockFirehoseAPIMockRecorder) ListDeliveryStreamsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDeliveryStreamsWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).ListDeliveryStreamsWithContext), varargs...)
}

// ListTagsForDeliveryStream mocks base method
func (m *MockFirehoseAPI) ListTagsForDeliveryStream(arg0 *firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.ListTagsForDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForDeliveryStream indicates an expected call of ListTagsForDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) ListTagsForDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).ListTagsForDeliveryStream), arg0)
}

// ListTagsForDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) ListTagsForDeliveryStreamRequest(arg0 *firehose.ListTagsForDeliveryStreamInput) (*request.Request, *firehose.ListTagsForDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTagsForDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.ListTagsForDeliveryStreamOutput)
	return ret0, ret1
}

// ListTagsForDeliveryStreamRequest indicates an expected call of ListTagsForDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) ListTagsForDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).ListTagsForDeliveryStreamRequest), arg0)
}

// ListTagsForDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) ListTagsForDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.ListTagsForDeliveryStreamInput, arg2 ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTagsForDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.ListTagsForDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTagsForDeliveryStreamWithContext indicates an expected call of ListTagsForDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) ListTagsForDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTagsForDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).ListTagsForDeliveryStreamWithContext), varargs...)
}

// PutRecord mocks base method
func (m *MockFirehoseAPI) PutRecord(arg0 *firehose.PutRecordInput) (*firehose.PutRecordOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecord", arg0)
	ret0, _ := ret[0].(*firehose.PutRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecord indicates an expected call of PutRecord
func (mr *MockFirehoseAPIMockRecorder) PutRecord(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecord", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecord), arg0)
}

// PutRecordBatch mocks base method
func (m *MockFirehoseAPI) PutRecordBatch(arg0 *firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecordBatch", arg0)
	ret0, _ := ret[0].(*firehose.PutRecordBatchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecordBatch indicates an expected call of PutRecordBatch
func (mr *MockFirehoseAPIMockRecorder) PutRecordBatch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordBatch", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecordBatch), arg0)
}

// PutRecordBatchRequest mocks base method
func (m *MockFirehoseAPI) PutRecordBatchRequest(arg0 *firehose.PutRecordBatchInput) (*request.Request, *firehose.PutRecordBatchOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecordBatchRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.PutRecordBatchOutput)
	return ret0, ret1
}

// PutRecordBatchRequest indicates an expected call of PutRecordBatchRequest
func (mr *MockFirehoseAPIMockRecorder) PutRecordBatchRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordBatchRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecordBatchRequest), arg0)
}

// PutRecordBatchWithContext mocks base method
func (m *MockFirehoseAPI) PutRecordBatchWithContext(arg0 context.Context, arg1 *firehose.PutRecordBatchInput, arg2 ...request.Option) (*firehose.PutRecordBatchOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutRecordBatchWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.PutRecordBatchOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecordBatchWithContext indicates an expected call of PutRecordBatchWithContext
func (mr *MockFirehoseAPIMockRecorder) PutRecordBatchWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordBatchWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecordBatchWithContext), varargs...)
}

// PutRecordRequest mocks base method
func (m *MockFirehoseAPI) PutRecordRequest(arg0 *firehose.PutRecordInput) (*request.Request, *firehose.PutRecordOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutRecordRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.PutRecordOutput)
	return ret0, ret1
}

// PutRecordRequest indicates an expected call of PutRecordRequest
func (mr *MockFirehoseAPIMockRecorder) PutRecordRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecordRequest), arg0)
}

// PutRecordWithContext mocks base method
func (m *MockFirehoseAPI) PutRecordWithContext(arg0 context.Context, arg1 *firehose.PutRecordInput, arg2 ...request.Option) (*firehose.PutRecordOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutRecordWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.PutRecordOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutRecordWithContext indicates an expected call of PutRecordWithContext
func (mr *MockFirehoseAPIMockRecorder) PutRecordWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutRecordWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).PutRecordWithContext), varargs...)
}

// StartDeliveryStreamEncryption mocks base method
func (m *MockFirehoseAPI) StartDeliveryStreamEncryption(arg0 *firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDeliveryStreamEncryption", arg0)
	ret0, _ := ret[0].(*firehose.StartDeliveryStreamEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartDeliveryStreamEncryption indicates an expected call of StartDeliveryStreamEncryption
func (mr *MockFirehoseAPIMockRecorder) StartDeliveryStreamEncryption(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeliveryStreamEncryption", reflect.TypeOf((*MockFirehoseAPI)(nil).StartDeliveryStreamEncryption), arg0)
}

// StartDeliveryStreamEncryptionRequest mocks base method
func (m *MockFirehoseAPI) StartDeliveryStreamEncryptionRequest(arg0 *firehose.StartDeliveryStreamEncryptionInput) (*request.Request, *firehose.StartDeliveryStreamEncryptionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartDeliveryStreamEncryptionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.StartDeliveryStreamEncryptionOutput)
	return ret0, ret1
}

// StartDeliveryStreamEncryptionRequest indicates an expected call of StartDeliveryStreamEncryptionRequest
func (mr *MockFirehoseAPIMockRecorder) StartDeliveryStreamEncryptionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeliveryStreamEncryptionRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).StartDeliveryStreamEncryptionRequest), arg0)
}

// StartDeliveryStreamEncryptionWithContext mocks base method
func (m *MockFirehoseAPI) StartDeliveryStreamEncryptionWithContext(arg0 context.Context, arg1 *firehose.StartDeliveryStreamEncryptionInput, arg2 ...request.Option) (*firehose.StartDeliveryStreamEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartDeliveryStreamEncryptionWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.StartDeliveryStreamEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartDeliveryStreamEncryptionWithContext indicates an expected call of StartDeliveryStreamEncryptionWithContext
func (mr *MockFirehoseAPIMockRecorder) StartDeliveryStreamEncryptionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartDeliveryStreamEncryptionWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).StartDeliveryStreamEncryptionWithContext), varargs...)
}

// StopDeliveryStreamEncryption mocks base method
func (m *MockFirehoseAPI) StopDeliveryStreamEncryption(arg0 *firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopDeliveryStreamEncryption", arg0)
	ret0, _ := ret[0].(*firehose.StopDeliveryStreamEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopDeliveryStreamEncryption indicates an expected call of StopDeliveryStreamEncryption
func (mr *MockFirehoseAPIMockRecorder) StopDeliveryStreamEncryption(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDeliveryStreamEncryption", reflect.TypeOf((*MockFirehoseAPI)(nil).StopDeliveryStreamEncryption), arg0)
}

// StopDeliveryStreamEncryptionRequest mocks base method
func (m *MockFirehoseAPI) StopDeliveryStreamEncryptionRequest(arg0 *firehose.StopDeliveryStreamEncryptionInput) (*request.Request, *firehose.StopDeliveryStreamEncryptionOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StopDeliveryStreamEncryptionRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.StopDeliveryStreamEncryptionOutput)
	return ret0, ret1
}

// StopDeliveryStreamEncryptionRequest indicates an expected call of StopDeliveryStreamEncryptionRequest
func (mr *MockFirehoseAPIMockRecorder) StopDeliveryStreamEncryptionRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDeliveryStreamEncryptionRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).StopDeliveryStreamEncryptionRequest), arg0)
}

// StopDeliveryStreamEncryptionWithContext mocks base method
func (m *MockFirehoseAPI) StopDeliveryStreamEncryptionWithContext(arg0 context.Context, arg1 *firehose.StopDeliveryStreamEncryptionInput, arg2 ...request.Option) (*firehose.StopDeliveryStreamEncryptionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StopDeliveryStreamEncryptionWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.StopDeliveryStreamEncryptionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StopDeliveryStreamEncryptionWithContext indicates an expected call of StopDeliveryStreamEncryptionWithContext
func (mr *MockFirehoseAPIMockRecorder) StopDeliveryStreamEncryptionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopDeliveryStreamEncryptionWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).StopDeliveryStreamEncryptionWithContext), varargs...)
}

// TagDeliveryStream mocks base method
func (m *MockFirehoseAPI) TagDeliveryStream(arg0 *firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.TagDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagDeliveryStream indicates an expected call of TagDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) TagDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).TagDeliveryStream), arg0)
}

// TagDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) TagDeliveryStreamRequest(arg0 *firehose.TagDeliveryStreamInput) (*request.Request, *firehose.TagDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TagDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.TagDeliveryStreamOutput)
	return ret0, ret1
}

// TagDeliveryStreamRequest indicates an expected call of TagDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) TagDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).TagDeliveryStreamRequest), arg0)
}

// TagDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) TagDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.TagDeliveryStreamInput, arg2 ...request.Option) (*firehose.TagDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TagDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.TagDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TagDeliveryStreamWithContext indicates an expected call of TagDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) TagDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TagDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).TagDeliveryStreamWithContext), varargs...)
}

// UntagDeliveryStream mocks base method
func (m *MockFirehoseAPI) UntagDeliveryStream(arg0 *firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagDeliveryStream", arg0)
	ret0, _ := ret[0].(*firehose.UntagDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagDeliveryStream indicates an expected call of UntagDeliveryStream
func (mr *MockFirehoseAPIMockRecorder) UntagDeliveryStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagDeliveryStream", reflect.TypeOf((*MockFirehoseAPI)(nil).UntagDeliveryStream), arg0)
}

// UntagDeliveryStreamRequest mocks base method
func (m *MockFirehoseAPI) UntagDeliveryStreamRequest(arg0 *firehose.UntagDeliveryStreamInput) (*request.Request, *firehose.UntagDeliveryStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UntagDeliveryStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.UntagDeliveryStreamOutput)
	return ret0, ret1
}

// UntagDeliveryStreamRequest indicates an expected call of UntagDeliveryStreamRequest
func (mr *MockFirehoseAPIMockRecorder) UntagDeliveryStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagDeliveryStreamRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).UntagDeliveryStreamRequest), arg0)
}

// UntagDeliveryStreamWithContext mocks base method
func (m *MockFirehoseAPI) UntagDeliveryStreamWithContext(arg0 context.Context, arg1 *firehose.UntagDeliveryStreamInput, arg2 ...request.Option) (*firehose.UntagDeliveryStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UntagDeliveryStreamWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.UntagDeliveryStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UntagDeliveryStreamWithContext indicates an expected call of UntagDeliveryStreamWithContext
func (mr *MockFirehoseAPIMockRecorder) UntagDeliveryStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UntagDeliveryStreamWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).UntagDeliveryStreamWithContext), varargs...)
}

// UpdateDestination mocks base method
func (m *MockFirehoseAPI) UpdateDestination(arg0 *firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDestination", arg0)
	ret0, _ := ret[0].(*firehose.UpdateDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDestination indicates an expected call of UpdateDestination
func (mr *MockFirehoseAPIMockRecorder) UpdateDestination(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDestination", reflect.TypeOf((*MockFirehoseAPI)(nil).UpdateDestination), arg0)
}

// UpdateDestinationRequest mocks base method
func (m *MockFirehoseAPI) UpdateDestinationRequest(arg0 *firehose.UpdateDestinationInput) (*request.Request, *firehose.UpdateDestinationOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDestinationRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*firehose.UpdateDestinationOutput)
	return ret0, ret1
}

// UpdateDestinationRequest indicates an expected call of UpdateDestinationRequest
func (mr *MockFirehoseAPIMockRecorder) UpdateDestinationRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDestinationRequest", reflect.TypeOf((*MockFirehoseAPI)(nil).UpdateDestinationRequest), arg0)
}

// UpdateDestinationWithContext mocks base method
func (m *MockFirehoseAPI) UpdateDestinationWithContext(arg0 context.Context, arg1 *firehose.UpdateDestinationInput, arg2 ...request.Option) (*firehose.UpdateDestinationOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateDestinationWithContext", varargs...)
	ret0, _ := ret[0].(*firehose.UpdateDestinationOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateDestinationWithContext indicates an expected call of UpdateDestinationWithContext
func (mr *MockFirehoseAPIMockRecorder) UpdateDestinationWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDestinationWithContext", reflect.TypeOf((*MockFirehoseAPI)(nil).UpdateDestinationWithContext), varargs...)
}
