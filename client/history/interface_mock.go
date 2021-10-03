// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go

// Package history is a generated GoMock package.
package history

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	yarpc "go.uber.org/yarpc"

	types "github.com/uber/cadence/common/types"
)

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// CloseShard mocks base method
func (m *MockClient) CloseShard(arg0 context.Context, arg1 *types.CloseShardRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CloseShard", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseShard indicates an expected call of CloseShard
func (mr *MockClientMockRecorder) CloseShard(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseShard", reflect.TypeOf((*MockClient)(nil).CloseShard), varargs...)
}

// DescribeHistoryHost mocks base method
func (m *MockClient) DescribeHistoryHost(arg0 context.Context, arg1 *types.DescribeHistoryHostRequest, arg2 ...yarpc.CallOption) (*types.DescribeHistoryHostResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeHistoryHost", varargs...)
	ret0, _ := ret[0].(*types.DescribeHistoryHostResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeHistoryHost indicates an expected call of DescribeHistoryHost
func (mr *MockClientMockRecorder) DescribeHistoryHost(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeHistoryHost", reflect.TypeOf((*MockClient)(nil).DescribeHistoryHost), varargs...)
}

// DescribeMutableState mocks base method
func (m *MockClient) DescribeMutableState(arg0 context.Context, arg1 *types.DescribeMutableStateRequest, arg2 ...yarpc.CallOption) (*types.DescribeMutableStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeMutableState", varargs...)
	ret0, _ := ret[0].(*types.DescribeMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeMutableState indicates an expected call of DescribeMutableState
func (mr *MockClientMockRecorder) DescribeMutableState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeMutableState", reflect.TypeOf((*MockClient)(nil).DescribeMutableState), varargs...)
}

// DescribeQueue mocks base method
func (m *MockClient) DescribeQueue(arg0 context.Context, arg1 *types.DescribeQueueRequest, arg2 ...yarpc.CallOption) (*types.DescribeQueueResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeQueue", varargs...)
	ret0, _ := ret[0].(*types.DescribeQueueResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeQueue indicates an expected call of DescribeQueue
func (mr *MockClientMockRecorder) DescribeQueue(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeQueue", reflect.TypeOf((*MockClient)(nil).DescribeQueue), varargs...)
}

// DescribeWorkflowExecution mocks base method
func (m *MockClient) DescribeWorkflowExecution(arg0 context.Context, arg1 *types.HistoryDescribeWorkflowExecutionRequest, arg2 ...yarpc.CallOption) (*types.DescribeWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DescribeWorkflowExecution", varargs...)
	ret0, _ := ret[0].(*types.DescribeWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeWorkflowExecution indicates an expected call of DescribeWorkflowExecution
func (mr *MockClientMockRecorder) DescribeWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeWorkflowExecution", reflect.TypeOf((*MockClient)(nil).DescribeWorkflowExecution), varargs...)
}

// GetDLQReplicationMessages mocks base method
func (m *MockClient) GetDLQReplicationMessages(arg0 context.Context, arg1 *types.GetDLQReplicationMessagesRequest, arg2 ...yarpc.CallOption) (*types.GetDLQReplicationMessagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDLQReplicationMessages", varargs...)
	ret0, _ := ret[0].(*types.GetDLQReplicationMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDLQReplicationMessages indicates an expected call of GetDLQReplicationMessages
func (mr *MockClientMockRecorder) GetDLQReplicationMessages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDLQReplicationMessages", reflect.TypeOf((*MockClient)(nil).GetDLQReplicationMessages), varargs...)
}

// GetMutableState mocks base method
func (m *MockClient) GetMutableState(arg0 context.Context, arg1 *types.GetMutableStateRequest, arg2 ...yarpc.CallOption) (*types.GetMutableStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetMutableState", varargs...)
	ret0, _ := ret[0].(*types.GetMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMutableState indicates an expected call of GetMutableState
func (mr *MockClientMockRecorder) GetMutableState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMutableState", reflect.TypeOf((*MockClient)(nil).GetMutableState), varargs...)
}

// GetReplicationMessages mocks base method
func (m *MockClient) GetReplicationMessages(arg0 context.Context, arg1 *types.GetReplicationMessagesRequest, arg2 ...yarpc.CallOption) (*types.GetReplicationMessagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetReplicationMessages", varargs...)
	ret0, _ := ret[0].(*types.GetReplicationMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetReplicationMessages indicates an expected call of GetReplicationMessages
func (mr *MockClientMockRecorder) GetReplicationMessages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicationMessages", reflect.TypeOf((*MockClient)(nil).GetReplicationMessages), varargs...)
}

// MergeDLQMessages mocks base method
func (m *MockClient) MergeDLQMessages(arg0 context.Context, arg1 *types.MergeDLQMessagesRequest, arg2 ...yarpc.CallOption) (*types.MergeDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "MergeDLQMessages", varargs...)
	ret0, _ := ret[0].(*types.MergeDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MergeDLQMessages indicates an expected call of MergeDLQMessages
func (mr *MockClientMockRecorder) MergeDLQMessages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MergeDLQMessages", reflect.TypeOf((*MockClient)(nil).MergeDLQMessages), varargs...)
}

// NotifyFailoverMarkers mocks base method
func (m *MockClient) NotifyFailoverMarkers(arg0 context.Context, arg1 *types.NotifyFailoverMarkersRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NotifyFailoverMarkers", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// NotifyFailoverMarkers indicates an expected call of NotifyFailoverMarkers
func (mr *MockClientMockRecorder) NotifyFailoverMarkers(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyFailoverMarkers", reflect.TypeOf((*MockClient)(nil).NotifyFailoverMarkers), varargs...)
}

// PollMutableState mocks base method
func (m *MockClient) PollMutableState(arg0 context.Context, arg1 *types.PollMutableStateRequest, arg2 ...yarpc.CallOption) (*types.PollMutableStateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PollMutableState", varargs...)
	ret0, _ := ret[0].(*types.PollMutableStateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PollMutableState indicates an expected call of PollMutableState
func (mr *MockClientMockRecorder) PollMutableState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PollMutableState", reflect.TypeOf((*MockClient)(nil).PollMutableState), varargs...)
}

// PurgeDLQMessages mocks base method
func (m *MockClient) PurgeDLQMessages(arg0 context.Context, arg1 *types.PurgeDLQMessagesRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PurgeDLQMessages", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PurgeDLQMessages indicates an expected call of PurgeDLQMessages
func (mr *MockClientMockRecorder) PurgeDLQMessages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PurgeDLQMessages", reflect.TypeOf((*MockClient)(nil).PurgeDLQMessages), varargs...)
}

// QueryWorkflow mocks base method
func (m *MockClient) QueryWorkflow(arg0 context.Context, arg1 *types.HistoryQueryWorkflowRequest, arg2 ...yarpc.CallOption) (*types.HistoryQueryWorkflowResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "QueryWorkflow", varargs...)
	ret0, _ := ret[0].(*types.HistoryQueryWorkflowResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryWorkflow indicates an expected call of QueryWorkflow
func (mr *MockClientMockRecorder) QueryWorkflow(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryWorkflow", reflect.TypeOf((*MockClient)(nil).QueryWorkflow), varargs...)
}

// ReadDLQMessages mocks base method
func (m *MockClient) ReadDLQMessages(arg0 context.Context, arg1 *types.ReadDLQMessagesRequest, arg2 ...yarpc.CallOption) (*types.ReadDLQMessagesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReadDLQMessages", varargs...)
	ret0, _ := ret[0].(*types.ReadDLQMessagesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDLQMessages indicates an expected call of ReadDLQMessages
func (mr *MockClientMockRecorder) ReadDLQMessages(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDLQMessages", reflect.TypeOf((*MockClient)(nil).ReadDLQMessages), varargs...)
}

// ReapplyEvents mocks base method
func (m *MockClient) ReapplyEvents(arg0 context.Context, arg1 *types.HistoryReapplyEventsRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReapplyEvents", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReapplyEvents indicates an expected call of ReapplyEvents
func (mr *MockClientMockRecorder) ReapplyEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReapplyEvents", reflect.TypeOf((*MockClient)(nil).ReapplyEvents), varargs...)
}

// RecordActivityTaskHeartbeat mocks base method
func (m *MockClient) RecordActivityTaskHeartbeat(arg0 context.Context, arg1 *types.HistoryRecordActivityTaskHeartbeatRequest, arg2 ...yarpc.CallOption) (*types.RecordActivityTaskHeartbeatResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecordActivityTaskHeartbeat", varargs...)
	ret0, _ := ret[0].(*types.RecordActivityTaskHeartbeatResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskHeartbeat indicates an expected call of RecordActivityTaskHeartbeat
func (mr *MockClientMockRecorder) RecordActivityTaskHeartbeat(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskHeartbeat", reflect.TypeOf((*MockClient)(nil).RecordActivityTaskHeartbeat), varargs...)
}

// RecordActivityTaskStarted mocks base method
func (m *MockClient) RecordActivityTaskStarted(arg0 context.Context, arg1 *types.RecordActivityTaskStartedRequest, arg2 ...yarpc.CallOption) (*types.RecordActivityTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecordActivityTaskStarted", varargs...)
	ret0, _ := ret[0].(*types.RecordActivityTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordActivityTaskStarted indicates an expected call of RecordActivityTaskStarted
func (mr *MockClientMockRecorder) RecordActivityTaskStarted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordActivityTaskStarted", reflect.TypeOf((*MockClient)(nil).RecordActivityTaskStarted), varargs...)
}

// RecordChildExecutionCompleted mocks base method
func (m *MockClient) RecordChildExecutionCompleted(arg0 context.Context, arg1 *types.RecordChildExecutionCompletedRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecordChildExecutionCompleted", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordChildExecutionCompleted indicates an expected call of RecordChildExecutionCompleted
func (mr *MockClientMockRecorder) RecordChildExecutionCompleted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordChildExecutionCompleted", reflect.TypeOf((*MockClient)(nil).RecordChildExecutionCompleted), varargs...)
}

// RecordDecisionTaskStarted mocks base method
func (m *MockClient) RecordDecisionTaskStarted(arg0 context.Context, arg1 *types.RecordDecisionTaskStartedRequest, arg2 ...yarpc.CallOption) (*types.RecordDecisionTaskStartedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RecordDecisionTaskStarted", varargs...)
	ret0, _ := ret[0].(*types.RecordDecisionTaskStartedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RecordDecisionTaskStarted indicates an expected call of RecordDecisionTaskStarted
func (mr *MockClientMockRecorder) RecordDecisionTaskStarted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordDecisionTaskStarted", reflect.TypeOf((*MockClient)(nil).RecordDecisionTaskStarted), varargs...)
}

// RefreshWorkflowTasks mocks base method
func (m *MockClient) RefreshWorkflowTasks(arg0 context.Context, arg1 *types.HistoryRefreshWorkflowTasksRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RefreshWorkflowTasks", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RefreshWorkflowTasks indicates an expected call of RefreshWorkflowTasks
func (mr *MockClientMockRecorder) RefreshWorkflowTasks(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshWorkflowTasks", reflect.TypeOf((*MockClient)(nil).RefreshWorkflowTasks), varargs...)
}

// RemoveSignalMutableState mocks base method
func (m *MockClient) RemoveSignalMutableState(arg0 context.Context, arg1 *types.RemoveSignalMutableStateRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSignalMutableState", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveSignalMutableState indicates an expected call of RemoveSignalMutableState
func (mr *MockClientMockRecorder) RemoveSignalMutableState(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSignalMutableState", reflect.TypeOf((*MockClient)(nil).RemoveSignalMutableState), varargs...)
}

// RemoveTask mocks base method
func (m *MockClient) RemoveTask(arg0 context.Context, arg1 *types.RemoveTaskRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTask indicates an expected call of RemoveTask
func (mr *MockClientMockRecorder) RemoveTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockClient)(nil).RemoveTask), varargs...)
}

// ReplicateEventsV2 mocks base method
func (m *MockClient) ReplicateEventsV2(arg0 context.Context, arg1 *types.ReplicateEventsV2Request, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ReplicateEventsV2", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReplicateEventsV2 indicates an expected call of ReplicateEventsV2
func (mr *MockClientMockRecorder) ReplicateEventsV2(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicateEventsV2", reflect.TypeOf((*MockClient)(nil).ReplicateEventsV2), varargs...)
}

// RequestCancelWorkflowExecution mocks base method
func (m *MockClient) RequestCancelWorkflowExecution(arg0 context.Context, arg1 *types.HistoryRequestCancelWorkflowExecutionRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RequestCancelWorkflowExecution", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestCancelWorkflowExecution indicates an expected call of RequestCancelWorkflowExecution
func (mr *MockClientMockRecorder) RequestCancelWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestCancelWorkflowExecution", reflect.TypeOf((*MockClient)(nil).RequestCancelWorkflowExecution), varargs...)
}

// ResetQueue mocks base method
func (m *MockClient) ResetQueue(arg0 context.Context, arg1 *types.ResetQueueRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetQueue", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ResetQueue indicates an expected call of ResetQueue
func (mr *MockClientMockRecorder) ResetQueue(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetQueue", reflect.TypeOf((*MockClient)(nil).ResetQueue), varargs...)
}

// ResetStickyTaskList mocks base method
func (m *MockClient) ResetStickyTaskList(arg0 context.Context, arg1 *types.HistoryResetStickyTaskListRequest, arg2 ...yarpc.CallOption) (*types.HistoryResetStickyTaskListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetStickyTaskList", varargs...)
	ret0, _ := ret[0].(*types.HistoryResetStickyTaskListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetStickyTaskList indicates an expected call of ResetStickyTaskList
func (mr *MockClientMockRecorder) ResetStickyTaskList(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetStickyTaskList", reflect.TypeOf((*MockClient)(nil).ResetStickyTaskList), varargs...)
}

// ResetWorkflowExecution mocks base method
func (m *MockClient) ResetWorkflowExecution(arg0 context.Context, arg1 *types.HistoryResetWorkflowExecutionRequest, arg2 ...yarpc.CallOption) (*types.ResetWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ResetWorkflowExecution", varargs...)
	ret0, _ := ret[0].(*types.ResetWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResetWorkflowExecution indicates an expected call of ResetWorkflowExecution
func (mr *MockClientMockRecorder) ResetWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResetWorkflowExecution", reflect.TypeOf((*MockClient)(nil).ResetWorkflowExecution), varargs...)
}

// RespondActivityTaskCanceled mocks base method
func (m *MockClient) RespondActivityTaskCanceled(arg0 context.Context, arg1 *types.HistoryRespondActivityTaskCanceledRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondActivityTaskCanceled", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCanceled indicates an expected call of RespondActivityTaskCanceled
func (mr *MockClientMockRecorder) RespondActivityTaskCanceled(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCanceled", reflect.TypeOf((*MockClient)(nil).RespondActivityTaskCanceled), varargs...)
}

// RespondActivityTaskCompleted mocks base method
func (m *MockClient) RespondActivityTaskCompleted(arg0 context.Context, arg1 *types.HistoryRespondActivityTaskCompletedRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondActivityTaskCompleted", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskCompleted indicates an expected call of RespondActivityTaskCompleted
func (mr *MockClientMockRecorder) RespondActivityTaskCompleted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskCompleted", reflect.TypeOf((*MockClient)(nil).RespondActivityTaskCompleted), varargs...)
}

// RespondActivityTaskFailed mocks base method
func (m *MockClient) RespondActivityTaskFailed(arg0 context.Context, arg1 *types.HistoryRespondActivityTaskFailedRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondActivityTaskFailed", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondActivityTaskFailed indicates an expected call of RespondActivityTaskFailed
func (mr *MockClientMockRecorder) RespondActivityTaskFailed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondActivityTaskFailed", reflect.TypeOf((*MockClient)(nil).RespondActivityTaskFailed), varargs...)
}

// RespondDecisionTaskCompleted mocks base method
func (m *MockClient) RespondDecisionTaskCompleted(arg0 context.Context, arg1 *types.HistoryRespondDecisionTaskCompletedRequest, arg2 ...yarpc.CallOption) (*types.HistoryRespondDecisionTaskCompletedResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondDecisionTaskCompleted", varargs...)
	ret0, _ := ret[0].(*types.HistoryRespondDecisionTaskCompletedResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RespondDecisionTaskCompleted indicates an expected call of RespondDecisionTaskCompleted
func (mr *MockClientMockRecorder) RespondDecisionTaskCompleted(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskCompleted", reflect.TypeOf((*MockClient)(nil).RespondDecisionTaskCompleted), varargs...)
}

// RespondDecisionTaskFailed mocks base method
func (m *MockClient) RespondDecisionTaskFailed(arg0 context.Context, arg1 *types.HistoryRespondDecisionTaskFailedRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RespondDecisionTaskFailed", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// RespondDecisionTaskFailed indicates an expected call of RespondDecisionTaskFailed
func (mr *MockClientMockRecorder) RespondDecisionTaskFailed(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RespondDecisionTaskFailed", reflect.TypeOf((*MockClient)(nil).RespondDecisionTaskFailed), varargs...)
}

// ScheduleDecisionTask mocks base method
func (m *MockClient) ScheduleDecisionTask(arg0 context.Context, arg1 *types.ScheduleDecisionTaskRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ScheduleDecisionTask", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ScheduleDecisionTask indicates an expected call of ScheduleDecisionTask
func (mr *MockClientMockRecorder) ScheduleDecisionTask(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleDecisionTask", reflect.TypeOf((*MockClient)(nil).ScheduleDecisionTask), varargs...)
}

// SignalWithStartWorkflowExecution mocks base method
func (m *MockClient) SignalWithStartWorkflowExecution(arg0 context.Context, arg1 *types.HistorySignalWithStartWorkflowExecutionRequest, arg2 ...yarpc.CallOption) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignalWithStartWorkflowExecution", varargs...)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignalWithStartWorkflowExecution indicates an expected call of SignalWithStartWorkflowExecution
func (mr *MockClientMockRecorder) SignalWithStartWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWithStartWorkflowExecution", reflect.TypeOf((*MockClient)(nil).SignalWithStartWorkflowExecution), varargs...)
}

// SignalWorkflowExecution mocks base method
func (m *MockClient) SignalWorkflowExecution(arg0 context.Context, arg1 *types.HistorySignalWorkflowExecutionRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SignalWorkflowExecution", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignalWorkflowExecution indicates an expected call of SignalWorkflowExecution
func (mr *MockClientMockRecorder) SignalWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignalWorkflowExecution", reflect.TypeOf((*MockClient)(nil).SignalWorkflowExecution), varargs...)
}

// StartWorkflowExecution mocks base method
func (m *MockClient) StartWorkflowExecution(arg0 context.Context, arg1 *types.HistoryStartWorkflowExecutionRequest, arg2 ...yarpc.CallOption) (*types.StartWorkflowExecutionResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "StartWorkflowExecution", varargs...)
	ret0, _ := ret[0].(*types.StartWorkflowExecutionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StartWorkflowExecution indicates an expected call of StartWorkflowExecution
func (mr *MockClientMockRecorder) StartWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartWorkflowExecution", reflect.TypeOf((*MockClient)(nil).StartWorkflowExecution), varargs...)
}

// SyncActivity mocks base method
func (m *MockClient) SyncActivity(arg0 context.Context, arg1 *types.SyncActivityRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncActivity", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncActivity indicates an expected call of SyncActivity
func (mr *MockClientMockRecorder) SyncActivity(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncActivity", reflect.TypeOf((*MockClient)(nil).SyncActivity), varargs...)
}

// SyncShardStatus mocks base method
func (m *MockClient) SyncShardStatus(arg0 context.Context, arg1 *types.SyncShardStatusRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SyncShardStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncShardStatus indicates an expected call of SyncShardStatus
func (mr *MockClientMockRecorder) SyncShardStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SyncShardStatus", reflect.TypeOf((*MockClient)(nil).SyncShardStatus), varargs...)
}

// TerminateWorkflowExecution mocks base method
func (m *MockClient) TerminateWorkflowExecution(arg0 context.Context, arg1 *types.HistoryTerminateWorkflowExecutionRequest, arg2 ...yarpc.CallOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TerminateWorkflowExecution", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TerminateWorkflowExecution indicates an expected call of TerminateWorkflowExecution
func (mr *MockClientMockRecorder) TerminateWorkflowExecution(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TerminateWorkflowExecution", reflect.TypeOf((*MockClient)(nil).TerminateWorkflowExecution), varargs...)
}
