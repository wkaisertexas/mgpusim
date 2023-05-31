// Code generated by MockGen. DO NOT EDIT.
// Source: alg.go

// Package dispatching is a generated GoMock package.
package dispatching

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	kernels "github.com/sarchlab/mgpusim/v3/kernels"
	resource "github.com/sarchlab/mgpusim/v3/timing/cp/internal/resource"
)

// MockAlgorithm is a mock of algorithm interface.
type MockAlgorithm struct {
	ctrl     *gomock.Controller
	recorder *MockAlgorithmMockRecorder
}

// MockAlgorithmMockRecorder is the mock recorder for MockAlgorithm.
type MockAlgorithmMockRecorder struct {
	mock *MockAlgorithm
}

// NewMockAlgorithm creates a new mock instance.
func NewMockAlgorithm(ctrl *gomock.Controller) *MockAlgorithm {
	mock := &MockAlgorithm{ctrl: ctrl}
	mock.recorder = &MockAlgorithmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAlgorithm) EXPECT() *MockAlgorithmMockRecorder {
	return m.recorder
}

// FreeResources mocks base method.
func (m *MockAlgorithm) FreeResources(location dispatchLocation) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "FreeResources", location)
}

// FreeResources indicates an expected call of FreeResources.
func (mr *MockAlgorithmMockRecorder) FreeResources(location interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreeResources", reflect.TypeOf((*MockAlgorithm)(nil).FreeResources), location)
}

// HasNext mocks base method.
func (m *MockAlgorithm) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockAlgorithmMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockAlgorithm)(nil).HasNext))
}

// Next mocks base method.
func (m *MockAlgorithm) Next() dispatchLocation {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(dispatchLocation)
	return ret0
}

// Next indicates an expected call of Next.
func (mr *MockAlgorithmMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockAlgorithm)(nil).Next))
}

// NumWG mocks base method.
func (m *MockAlgorithm) NumWG() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumWG")
	ret0, _ := ret[0].(int)
	return ret0
}

// NumWG indicates an expected call of NumWG.
func (mr *MockAlgorithmMockRecorder) NumWG() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumWG", reflect.TypeOf((*MockAlgorithm)(nil).NumWG))
}

// RegisterCU mocks base method.
func (m *MockAlgorithm) RegisterCU(cu resource.DispatchableCU) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterCU", cu)
}

// RegisterCU indicates an expected call of RegisterCU.
func (mr *MockAlgorithmMockRecorder) RegisterCU(cu interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterCU", reflect.TypeOf((*MockAlgorithm)(nil).RegisterCU), cu)
}

// StartNewKernel mocks base method.
func (m *MockAlgorithm) StartNewKernel(info kernels.KernelLaunchInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartNewKernel", info)
}

// StartNewKernel indicates an expected call of StartNewKernel.
func (mr *MockAlgorithmMockRecorder) StartNewKernel(info interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartNewKernel", reflect.TypeOf((*MockAlgorithm)(nil).StartNewKernel), info)
}
