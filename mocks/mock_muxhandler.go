// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/kubeedge/viaduct/pkg/mux (interfaces: Handler)
// souce tesr
// Package mocks is a generated GoMock package.
// testing viaduct bot
//test viaduct
//testingggggggggggggggggggggggggggggg
//testtttttttttttttttttttttt
// tttttttttttttttttttttttttttttttttttt
//testingggggggggggggggggggggggggggg
//testingggggg
//+!!!
//test +1
//Source to test
//one more test
//one of the test
//testing
//+test
//add test
//testing
//++++
// VALUES +1 +2
//++++++++++ Values
//+0
//+1
//+2
//+3
//+4
//+5
//+6
//+7
//+8
//+1
//test12
//test13
//+9
//+10
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	model "github.com/kubeedge/beehive/pkg/core/model"
	mux "github.com/kubeedge/viaduct/pkg/mux"
	reflect "reflect"
)

// MockHandler is a mock of Handler interface
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// ServeConn mocks base method
func (m *MockHandler) ServeConn(arg0 *model.Message, arg1 mux.ResponseWriter) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeConn", arg0, arg1)
}

// ServeConn indicates an expected call of ServeConn
func (mr *MockHandlerMockRecorder) ServeConn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeConn", reflect.TypeOf((*MockHandler)(nil).ServeConn), arg0, arg1)
}
