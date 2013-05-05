// Automatically generated by MockGen. DO NOT EDIT!
// Source: spoj1.go

package mock_spoj

import (
	gomock "code.google.com/p/gomock/gomock"
	myio "dlluncor/myio"
)

// Mock of BitmapperI interface
type MockBitmapperI struct {
	ctrl     *gomock.Controller
	recorder *_MockBitmapperIRecorder
}

// Recorder for MockBitmapperI (not exported)
type _MockBitmapperIRecorder struct {
	mock *MockBitmapperI
}

func NewMockBitmapperI(ctrl *gomock.Controller) *MockBitmapperI {
	mock := &MockBitmapperI{ctrl: ctrl}
	mock.recorder = &_MockBitmapperIRecorder{mock}
	return mock
}

func (_m *MockBitmapperI) EXPECT() *_MockBitmapperIRecorder {
	return _m.recorder
}

func (_m *MockBitmapperI) ReadInput(in myio.Reader) {
	_m.ctrl.Call(_m, "ReadInput", in)
}

func (_mr *_MockBitmapperIRecorder) ReadInput(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ReadInput", arg0)
}

func (_m *MockBitmapperI) Solve() string {
	ret := _m.ctrl.Call(_m, "Solve")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockBitmapperIRecorder) Solve() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Solve")
}
