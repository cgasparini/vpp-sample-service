// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/ligato/vpp-agent/clientv1/defaultplugins/data_resync_api.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	. "github.com/ligato/vpp-agent/clientv1/defaultplugins"
	acl "github.com/ligato/vpp-agent/plugins/defaultplugins/aclplugin/model/acl"
	interfaces "github.com/ligato/vpp-agent/plugins/defaultplugins/ifplugin/model/interfaces"
	l2 "github.com/ligato/vpp-agent/plugins/defaultplugins/l2plugin/model/l2"
	l3 "github.com/ligato/vpp-agent/plugins/defaultplugins/l3plugin/model/l3"
	reflect "reflect"
)

// MockDataResyncDSL is a mock of DataResyncDSL interface
type MockDataResyncDSL struct {
	ctrl     *gomock.Controller
	recorder *MockDataResyncDSLMockRecorder
}

// MockDataResyncDSLMockRecorder is the mock recorder for MockDataResyncDSL
type MockDataResyncDSLMockRecorder struct {
	mock *MockDataResyncDSL
}

// NewMockDataResyncDSL creates a new mock instance
func NewMockDataResyncDSL(ctrl *gomock.Controller) *MockDataResyncDSL {
	mock := &MockDataResyncDSL{ctrl: ctrl}
	mock.recorder = &MockDataResyncDSLMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDataResyncDSL) EXPECT() *MockDataResyncDSLMockRecorder {
	return m.recorder
}

// Interface mocks base method
func (m *MockDataResyncDSL) Interface(intf *interfaces.Interfaces_Interface) DataResyncDSL {
	ret := m.ctrl.Call(m, "Interface", intf)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// Interface indicates an expected call of Interface
func (mr *MockDataResyncDSLMockRecorder) Interface(intf interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Interface", reflect.TypeOf((*MockDataResyncDSL)(nil).Interface), intf)
}

// BD mocks base method
func (m *MockDataResyncDSL) BD(bd *l2.BridgeDomains_BridgeDomain) DataResyncDSL {
	ret := m.ctrl.Call(m, "BD", bd)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// BD indicates an expected call of BD
func (mr *MockDataResyncDSLMockRecorder) BD(bd interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BD", reflect.TypeOf((*MockDataResyncDSL)(nil).BD), bd)
}

// BDFIB mocks base method
func (m *MockDataResyncDSL) BDFIB(fib *l2.FibTableEntries_FibTableEntry) DataResyncDSL {
	ret := m.ctrl.Call(m, "BDFIB", fib)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// BDFIB indicates an expected call of BDFIB
func (mr *MockDataResyncDSLMockRecorder) BDFIB(fib interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BDFIB", reflect.TypeOf((*MockDataResyncDSL)(nil).BDFIB), fib)
}

// XConnect mocks base method
func (m *MockDataResyncDSL) XConnect(xcon *l2.XConnectPairs_XConnectPair) DataResyncDSL {
	ret := m.ctrl.Call(m, "XConnect", xcon)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// XConnect indicates an expected call of XConnect
func (mr *MockDataResyncDSLMockRecorder) XConnect(xcon interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XConnect", reflect.TypeOf((*MockDataResyncDSL)(nil).XConnect), xcon)
}

// StaticRoute mocks base method
func (m *MockDataResyncDSL) StaticRoute(staticRoute *l3.StaticRoutes_Route) DataResyncDSL {
	ret := m.ctrl.Call(m, "StaticRoute", staticRoute)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// StaticRoute indicates an expected call of StaticRoute
func (mr *MockDataResyncDSLMockRecorder) StaticRoute(staticRoute interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StaticRoute", reflect.TypeOf((*MockDataResyncDSL)(nil).StaticRoute), staticRoute)
}

// ACL mocks base method
func (m *MockDataResyncDSL) ACL(acl *acl.AccessLists_Acl) DataResyncDSL {
	ret := m.ctrl.Call(m, "ACL", acl)
	ret0, _ := ret[0].(DataResyncDSL)
	return ret0
}

// ACL indicates an expected call of ACL
func (mr *MockDataResyncDSLMockRecorder) ACL(acl interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ACL", reflect.TypeOf((*MockDataResyncDSL)(nil).ACL), acl)
}

// Send mocks base method
func (m *MockDataResyncDSL) Send() Reply {
	ret := m.ctrl.Call(m, "Send")
	ret0, _ := ret[0].(Reply)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockDataResyncDSLMockRecorder) Send() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockDataResyncDSL)(nil).Send))
}