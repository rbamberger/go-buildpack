// Automatically generated by MockGen. DO NOT EDIT!
// Source: ../vendor/github.com/cloudfoundry/libbuildpack/command_runner.go

package finalize_test

import (
	gomock "github.com/golang/mock/gomock"
	io "io"
)

// Mock of CommandRunner interface
type MockCommandRunner struct {
	ctrl     *gomock.Controller
	recorder *_MockCommandRunnerRecorder
}

// Recorder for MockCommandRunner (not exported)
type _MockCommandRunnerRecorder struct {
	mock *MockCommandRunner
}

func NewMockCommandRunner(ctrl *gomock.Controller) *MockCommandRunner {
	mock := &MockCommandRunner{ctrl: ctrl}
	mock.recorder = &_MockCommandRunnerRecorder{mock}
	return mock
}

func (_m *MockCommandRunner) EXPECT() *_MockCommandRunnerRecorder {
	return _m.recorder
}

func (_m *MockCommandRunner) SetOutput(_param0 io.Writer) {
	_m.ctrl.Call(_m, "SetOutput", _param0)
}

func (_mr *_MockCommandRunnerRecorder) SetOutput(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetOutput", arg0)
}

func (_m *MockCommandRunner) SetStdout(_param0 io.Writer) {
	_m.ctrl.Call(_m, "SetStdout", _param0)
}

func (_mr *_MockCommandRunnerRecorder) SetStdout(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetStdout", arg0)
}

func (_m *MockCommandRunner) SetStderr(_param0 io.Writer) {
	_m.ctrl.Call(_m, "SetStderr", _param0)
}

func (_mr *_MockCommandRunnerRecorder) SetStderr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetStderr", arg0)
}

func (_m *MockCommandRunner) SetDir(_param0 string) {
	_m.ctrl.Call(_m, "SetDir", _param0)
}

func (_mr *_MockCommandRunnerRecorder) SetDir(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetDir", arg0)
}

func (_m *MockCommandRunner) Reset() {
	_m.ctrl.Call(_m, "Reset")
}

func (_mr *_MockCommandRunnerRecorder) Reset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Reset")
}

func (_m *MockCommandRunner) ResetOutput() {
	_m.ctrl.Call(_m, "ResetOutput")
}

func (_mr *_MockCommandRunnerRecorder) ResetOutput() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ResetOutput")
}

func (_m *MockCommandRunner) Run(program string, args ...string) error {
	_s := []interface{}{program}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "Run", _s...)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockCommandRunnerRecorder) Run(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", _s...)
}

func (_m *MockCommandRunner) CaptureOutput(program string, args ...string) (string, error) {
	_s := []interface{}{program}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CaptureOutput", _s...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommandRunnerRecorder) CaptureOutput(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CaptureOutput", _s...)
}

func (_m *MockCommandRunner) CaptureStdout(program string, args ...string) (string, error) {
	_s := []interface{}{program}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CaptureStdout", _s...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommandRunnerRecorder) CaptureStdout(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CaptureStdout", _s...)
}

func (_m *MockCommandRunner) CaptureStderr(program string, args ...string) (string, error) {
	_s := []interface{}{program}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	ret := _m.ctrl.Call(_m, "CaptureStderr", _s...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockCommandRunnerRecorder) CaptureStderr(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CaptureStderr", _s...)
}
