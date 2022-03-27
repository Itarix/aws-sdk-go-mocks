// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/aws-sdk-go/service/dynamodbstreams/dynamodbstreamsiface (interfaces: DynamoDBStreamsAPI)

// Package dynamodbstreamsmock is a generated GoMock package.
package dynamodbstreamsmock

import (
	aws "github.com/aws/aws-sdk-go/aws"
	request "github.com/aws/aws-sdk-go/aws/request"
	dynamodbstreams "github.com/aws/aws-sdk-go/service/dynamodbstreams"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDynamoDBStreamsAPI is a mock of DynamoDBStreamsAPI interface
type MockDynamoDBStreamsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockDynamoDBStreamsAPIMockRecorder
}

// MockDynamoDBStreamsAPIMockRecorder is the mock recorder for MockDynamoDBStreamsAPI
type MockDynamoDBStreamsAPIMockRecorder struct {
	mock *MockDynamoDBStreamsAPI
}

// NewMockDynamoDBStreamsAPI creates a new mock instance
func NewMockDynamoDBStreamsAPI(ctrl *gomock.Controller) *MockDynamoDBStreamsAPI {
	mock := &MockDynamoDBStreamsAPI{ctrl: ctrl}
	mock.recorder = &MockDynamoDBStreamsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamoDBStreamsAPI) EXPECT() *MockDynamoDBStreamsAPIMockRecorder {
	return m.recorder
}

// DescribeStream mocks base method
func (m *MockDynamoDBStreamsAPI) DescribeStream(arg0 *dynamodbstreams.DescribeStreamInput) (*dynamodbstreams.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStream", arg0)
	ret0, _ := ret[0].(*dynamodbstreams.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStream indicates an expected call of DescribeStream
func (mr *MockDynamoDBStreamsAPIMockRecorder) DescribeStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStream", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).DescribeStream), arg0)
}

// DescribeStreamRequest mocks base method
func (m *MockDynamoDBStreamsAPI) DescribeStreamRequest(arg0 *dynamodbstreams.DescribeStreamInput) (*request.Request, *dynamodbstreams.DescribeStreamOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeStreamRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodbstreams.DescribeStreamOutput)
	return ret0, ret1
}

// DescribeStreamRequest indicates an expected call of DescribeStreamRequest
func (mr *MockDynamoDBStreamsAPIMockRecorder) DescribeStreamRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStreamRequest", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).DescribeStreamRequest), arg0)
}

// DescribeStreamWithContext mocks base method
func (m *MockDynamoDBStreamsAPI) DescribeStreamWithContext(arg0 aws.Context, arg1 *dynamodbstreams.DescribeStreamInput, arg2 ...request.Option) (*dynamodbstreams.DescribeStreamOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeStreamWithContext", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.DescribeStreamOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeStreamWithContext indicates an expected call of DescribeStreamWithContext
func (mr *MockDynamoDBStreamsAPIMockRecorder) DescribeStreamWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeStreamWithContext", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).DescribeStreamWithContext), varargs...)
}

// GetRecords mocks base method
func (m *MockDynamoDBStreamsAPI) GetRecords(arg0 *dynamodbstreams.GetRecordsInput) (*dynamodbstreams.GetRecordsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecords", arg0)
	ret0, _ := ret[0].(*dynamodbstreams.GetRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecords indicates an expected call of GetRecords
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetRecords(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecords", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetRecords), arg0)
}

// GetRecordsRequest mocks base method
func (m *MockDynamoDBStreamsAPI) GetRecordsRequest(arg0 *dynamodbstreams.GetRecordsInput) (*request.Request, *dynamodbstreams.GetRecordsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodbstreams.GetRecordsOutput)
	return ret0, ret1
}

// GetRecordsRequest indicates an expected call of GetRecordsRequest
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetRecordsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsRequest", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetRecordsRequest), arg0)
}

// GetRecordsWithContext mocks base method
func (m *MockDynamoDBStreamsAPI) GetRecordsWithContext(arg0 aws.Context, arg1 *dynamodbstreams.GetRecordsInput, arg2 ...request.Option) (*dynamodbstreams.GetRecordsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRecordsWithContext", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.GetRecordsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordsWithContext indicates an expected call of GetRecordsWithContext
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetRecordsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordsWithContext", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetRecordsWithContext), varargs...)
}

// GetShardIterator mocks base method
func (m *MockDynamoDBStreamsAPI) GetShardIterator(arg0 *dynamodbstreams.GetShardIteratorInput) (*dynamodbstreams.GetShardIteratorOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardIterator", arg0)
	ret0, _ := ret[0].(*dynamodbstreams.GetShardIteratorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardIterator indicates an expected call of GetShardIterator
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetShardIterator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardIterator", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetShardIterator), arg0)
}

// GetShardIteratorRequest mocks base method
func (m *MockDynamoDBStreamsAPI) GetShardIteratorRequest(arg0 *dynamodbstreams.GetShardIteratorInput) (*request.Request, *dynamodbstreams.GetShardIteratorOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardIteratorRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodbstreams.GetShardIteratorOutput)
	return ret0, ret1
}

// GetShardIteratorRequest indicates an expected call of GetShardIteratorRequest
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetShardIteratorRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardIteratorRequest", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetShardIteratorRequest), arg0)
}

// GetShardIteratorWithContext mocks base method
func (m *MockDynamoDBStreamsAPI) GetShardIteratorWithContext(arg0 aws.Context, arg1 *dynamodbstreams.GetShardIteratorInput, arg2 ...request.Option) (*dynamodbstreams.GetShardIteratorOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetShardIteratorWithContext", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.GetShardIteratorOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetShardIteratorWithContext indicates an expected call of GetShardIteratorWithContext
func (mr *MockDynamoDBStreamsAPIMockRecorder) GetShardIteratorWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardIteratorWithContext", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).GetShardIteratorWithContext), varargs...)
}

// ListStreams mocks base method
func (m *MockDynamoDBStreamsAPI) ListStreams(arg0 *dynamodbstreams.ListStreamsInput) (*dynamodbstreams.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStreams", arg0)
	ret0, _ := ret[0].(*dynamodbstreams.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStreams indicates an expected call of ListStreams
func (mr *MockDynamoDBStreamsAPIMockRecorder) ListStreams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreams", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).ListStreams), arg0)
}

// ListStreamsRequest mocks base method
func (m *MockDynamoDBStreamsAPI) ListStreamsRequest(arg0 *dynamodbstreams.ListStreamsInput) (*request.Request, *dynamodbstreams.ListStreamsOutput) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStreamsRequest", arg0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*dynamodbstreams.ListStreamsOutput)
	return ret0, ret1
}

// ListStreamsRequest indicates an expected call of ListStreamsRequest
func (mr *MockDynamoDBStreamsAPIMockRecorder) ListStreamsRequest(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreamsRequest", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).ListStreamsRequest), arg0)
}

// ListStreamsWithContext mocks base method
func (m *MockDynamoDBStreamsAPI) ListStreamsWithContext(arg0 aws.Context, arg1 *dynamodbstreams.ListStreamsInput, arg2 ...request.Option) (*dynamodbstreams.ListStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListStreamsWithContext", varargs...)
	ret0, _ := ret[0].(*dynamodbstreams.ListStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStreamsWithContext indicates an expected call of ListStreamsWithContext
func (mr *MockDynamoDBStreamsAPIMockRecorder) ListStreamsWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStreamsWithContext", reflect.TypeOf((*MockDynamoDBStreamsAPI)(nil).ListStreamsWithContext), varargs...)
}
