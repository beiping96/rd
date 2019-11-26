// Code generated by MockGen. DO NOT EDIT.
// Source: rd.go

// Package mock is a generated GoMock package.
package mock

import (
	rd "github.com/beiping96/rd"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRD is a mock of RD interface
type MockRD struct {
	ctrl     *gomock.Controller
	recorder *MockRDMockRecorder
}

// MockRDMockRecorder is the mock recorder for MockRD
type MockRDMockRecorder struct {
	mock *MockRD
}

// NewMockRD creates a new mock instance
func NewMockRD(ctrl *gomock.Controller) *MockRD {
	mock := &MockRD{ctrl: ctrl}
	mock.recorder = &MockRDMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRD) EXPECT() *MockRDMockRecorder {
	return m.recorder
}

// ErrNil mocks base method
func (m *MockRD) ErrNil() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ErrNil")
	ret0, _ := ret[0].(error)
	return ret0
}

// ErrNil indicates an expected call of ErrNil
func (mr *MockRDMockRecorder) ErrNil() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrNil", reflect.TypeOf((*MockRD)(nil).ErrNil))
}

// Do mocks base method
func (m *MockRD) Do(cmd string, args ...interface{}) (interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Do", varargs...)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do
func (mr *MockRDMockRecorder) Do(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockRD)(nil).Do), varargs...)
}

// DoBool mocks base method
func (m *MockRD) DoBool(cmd string, args ...interface{}) (bool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoBool", varargs...)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoBool indicates an expected call of DoBool
func (mr *MockRDMockRecorder) DoBool(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoBool", reflect.TypeOf((*MockRD)(nil).DoBool), varargs...)
}

// DoBytes mocks base method
func (m *MockRD) DoBytes(cmd string, args ...interface{}) ([]byte, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoBytes", varargs...)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoBytes indicates an expected call of DoBytes
func (mr *MockRDMockRecorder) DoBytes(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoBytes", reflect.TypeOf((*MockRD)(nil).DoBytes), varargs...)
}

// DoFloat64 mocks base method
func (m *MockRD) DoFloat64(cmd string, args ...interface{}) (float64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoFloat64", varargs...)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoFloat64 indicates an expected call of DoFloat64
func (mr *MockRDMockRecorder) DoFloat64(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoFloat64", reflect.TypeOf((*MockRD)(nil).DoFloat64), varargs...)
}

// DoInt mocks base method
func (m *MockRD) DoInt(cmd string, args ...interface{}) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoInt", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoInt indicates an expected call of DoInt
func (mr *MockRDMockRecorder) DoInt(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInt", reflect.TypeOf((*MockRD)(nil).DoInt), varargs...)
}

// DoInt64 mocks base method
func (m *MockRD) DoInt64(cmd string, args ...interface{}) (int64, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoInt64", varargs...)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoInt64 indicates an expected call of DoInt64
func (mr *MockRDMockRecorder) DoInt64(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInt64", reflect.TypeOf((*MockRD)(nil).DoInt64), varargs...)
}

// DoInts mocks base method
func (m *MockRD) DoInts(cmd string, args ...interface{}) ([]int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoInts", varargs...)
	ret0, _ := ret[0].([]int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoInts indicates an expected call of DoInts
func (mr *MockRDMockRecorder) DoInts(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoInts", reflect.TypeOf((*MockRD)(nil).DoInts), varargs...)
}

// DoString mocks base method
func (m *MockRD) DoString(cmd string, args ...interface{}) (string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoString", varargs...)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoString indicates an expected call of DoString
func (mr *MockRDMockRecorder) DoString(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoString", reflect.TypeOf((*MockRD)(nil).DoString), varargs...)
}

// DoStringMap mocks base method
func (m *MockRD) DoStringMap(cmd string, args ...interface{}) (map[string]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoStringMap", varargs...)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoStringMap indicates an expected call of DoStringMap
func (mr *MockRDMockRecorder) DoStringMap(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoStringMap", reflect.TypeOf((*MockRD)(nil).DoStringMap), varargs...)
}

// DoStrings mocks base method
func (m *MockRD) DoStrings(cmd string, args ...interface{}) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DoStrings", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DoStrings indicates an expected call of DoStrings
func (mr *MockRDMockRecorder) DoStrings(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DoStrings", reflect.TypeOf((*MockRD)(nil).DoStrings), varargs...)
}

// Values mocks base method
func (m *MockRD) Values(cmd string, args ...interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{cmd}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Values", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Values indicates an expected call of Values
func (mr *MockRDMockRecorder) Values(cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Values", reflect.TypeOf((*MockRD)(nil).Values), varargs...)
}

// Scan mocks base method
func (m *MockRD) Scan(src []interface{}, dst ...interface{}) ([]interface{}, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{src}
	for _, a := range dst {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Scan", varargs...)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan
func (mr *MockRDMockRecorder) Scan(src interface{}, dst ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{src}, dst...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRD)(nil).Scan), varargs...)
}

// Close mocks base method
func (m *MockRD) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockRDMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRD)(nil).Close))
}

// construct mocks base method
func (m *MockRD) construct(arg0 rd.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "construct", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// construct indicates an expected call of construct
func (mr *MockRDMockRecorder) construct(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "construct", reflect.TypeOf((*MockRD)(nil).construct), arg0)
}
