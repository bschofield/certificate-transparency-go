// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/google/trillian (interfaces: TrillianLogClient)

package mockclient

import (
	gomock "github.com/golang/mock/gomock"
	trillian "github.com/google/trillian"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Mock of TrillianLogClient interface
type MockTrillianLogClient struct {
	ctrl     *gomock.Controller
	recorder *_MockTrillianLogClientRecorder
}

// Recorder for MockTrillianLogClient (not exported)
type _MockTrillianLogClientRecorder struct {
	mock *MockTrillianLogClient
}

func NewMockTrillianLogClient(ctrl *gomock.Controller) *MockTrillianLogClient {
	mock := &MockTrillianLogClient{ctrl: ctrl}
	mock.recorder = &_MockTrillianLogClientRecorder{mock}
	return mock
}

func (_m *MockTrillianLogClient) EXPECT() *_MockTrillianLogClientRecorder {
	return _m.recorder
}

func (_m *MockTrillianLogClient) GetConsistencyProof(_param0 context.Context, _param1 *trillian.GetConsistencyProofRequest, _param2 ...grpc.CallOption) (*trillian.GetConsistencyProofResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetConsistencyProof", _s...)
	ret0, _ := ret[0].(*trillian.GetConsistencyProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetConsistencyProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetConsistencyProof", _s...)
}

func (_m *MockTrillianLogClient) GetEntryAndProof(_param0 context.Context, _param1 *trillian.GetEntryAndProofRequest, _param2 ...grpc.CallOption) (*trillian.GetEntryAndProofResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetEntryAndProof", _s...)
	ret0, _ := ret[0].(*trillian.GetEntryAndProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetEntryAndProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEntryAndProof", _s...)
}

func (_m *MockTrillianLogClient) GetInclusionProof(_param0 context.Context, _param1 *trillian.GetInclusionProofRequest, _param2 ...grpc.CallOption) (*trillian.GetInclusionProofResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetInclusionProof", _s...)
	ret0, _ := ret[0].(*trillian.GetInclusionProofResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetInclusionProof(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInclusionProof", _s...)
}

func (_m *MockTrillianLogClient) GetInclusionProofByHash(_param0 context.Context, _param1 *trillian.GetInclusionProofByHashRequest, _param2 ...grpc.CallOption) (*trillian.GetInclusionProofByHashResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetInclusionProofByHash", _s...)
	ret0, _ := ret[0].(*trillian.GetInclusionProofByHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetInclusionProofByHash(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetInclusionProofByHash", _s...)
}

func (_m *MockTrillianLogClient) GetLatestSignedLogRoot(_param0 context.Context, _param1 *trillian.GetLatestSignedLogRootRequest, _param2 ...grpc.CallOption) (*trillian.GetLatestSignedLogRootResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetLatestSignedLogRoot", _s...)
	ret0, _ := ret[0].(*trillian.GetLatestSignedLogRootResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetLatestSignedLogRoot(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLatestSignedLogRoot", _s...)
}

func (_m *MockTrillianLogClient) GetLeavesByHash(_param0 context.Context, _param1 *trillian.GetLeavesByHashRequest, _param2 ...grpc.CallOption) (*trillian.GetLeavesByHashResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetLeavesByHash", _s...)
	ret0, _ := ret[0].(*trillian.GetLeavesByHashResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetLeavesByHash(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLeavesByHash", _s...)
}

func (_m *MockTrillianLogClient) GetLeavesByIndex(_param0 context.Context, _param1 *trillian.GetLeavesByIndexRequest, _param2 ...grpc.CallOption) (*trillian.GetLeavesByIndexResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetLeavesByIndex", _s...)
	ret0, _ := ret[0].(*trillian.GetLeavesByIndexResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetLeavesByIndex(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetLeavesByIndex", _s...)
}

func (_m *MockTrillianLogClient) GetSequencedLeafCount(_param0 context.Context, _param1 *trillian.GetSequencedLeafCountRequest, _param2 ...grpc.CallOption) (*trillian.GetSequencedLeafCountResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "GetSequencedLeafCount", _s...)
	ret0, _ := ret[0].(*trillian.GetSequencedLeafCountResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) GetSequencedLeafCount(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSequencedLeafCount", _s...)
}

func (_m *MockTrillianLogClient) QueueLeaf(_param0 context.Context, _param1 *trillian.QueueLeafRequest, _param2 ...grpc.CallOption) (*trillian.QueueLeafResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "QueueLeaf", _s...)
	ret0, _ := ret[0].(*trillian.QueueLeafResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) QueueLeaf(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueueLeaf", _s...)
}

func (_m *MockTrillianLogClient) QueueLeaves(_param0 context.Context, _param1 *trillian.QueueLeavesRequest, _param2 ...grpc.CallOption) (*trillian.QueueLeavesResponse, error) {
	_s := []interface{}{_param0, _param1}
	for _, _x := range _param2 {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "QueueLeaves", _s...)
	ret0, _ := ret[0].(*trillian.QueueLeavesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTrillianLogClientRecorder) QueueLeaves(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0, arg1}, arg2...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueueLeaves", _s...)
}