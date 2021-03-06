// Code generated by MockGen. DO NOT EDIT.
// Source: ./orders/orders.go

// Package orders is a generated GoMock package.
package orders

import (
	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockOrderManager is a mock of OrderManager interface
type MockOrderManager struct {
	ctrl     *gomock.Controller
	recorder *MockOrderManagerMockRecorder
}

// MockOrderManagerMockRecorder is the mock recorder for MockOrderManager
type MockOrderManagerMockRecorder struct {
	mock *MockOrderManager
}

// NewMockOrderManager creates a new mock instance
func NewMockOrderManager(ctrl *gomock.Controller) *MockOrderManager {
	mock := &MockOrderManager{ctrl: ctrl}
	mock.recorder = &MockOrderManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderManager) EXPECT() *MockOrderManagerMockRecorder {
	return m.recorder
}

// Setup mocks base method
func (m *MockOrderManager) Setup() error {
	ret := m.ctrl.Call(m, "Setup")
	ret0, _ := ret[0].(error)
	return ret0
}

// Setup indicates an expected call of Setup
func (mr *MockOrderManagerMockRecorder) Setup() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Setup", reflect.TypeOf((*MockOrderManager)(nil).Setup))
}

// SetupSimulation mocks base method
func (m *MockOrderManager) SetupSimulation(arg0 map[string]*Balance) error {
	ret := m.ctrl.Call(m, "SetupSimulation", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetupSimulation indicates an expected call of SetupSimulation
func (mr *MockOrderManagerMockRecorder) SetupSimulation(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetupSimulation", reflect.TypeOf((*MockOrderManager)(nil).SetupSimulation), arg0)
}

// GetBalance mocks base method
func (m *MockOrderManager) GetBalance(marketKey string) *Balance {
	ret := m.ctrl.Call(m, "GetBalance", marketKey)
	ret0, _ := ret[0].(*Balance)
	return ret0
}

// GetBalance indicates an expected call of GetBalance
func (mr *MockOrderManagerMockRecorder) GetBalance(marketKey interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockOrderManager)(nil).GetBalance), marketKey)
}

// GetBalances mocks base method
func (m *MockOrderManager) GetBalances() map[string]*Balance {
	ret := m.ctrl.Call(m, "GetBalances")
	ret0, _ := ret[0].(map[string]*Balance)
	return ret0
}

// GetBalances indicates an expected call of GetBalances
func (mr *MockOrderManagerMockRecorder) GetBalances() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalances", reflect.TypeOf((*MockOrderManager)(nil).GetBalances))
}

// ExecuteLimitSell mocks base method
func (m *MockOrderManager) ExecuteLimitSell(ctx context.Context, order *LimitOrder) error {
	ret := m.ctrl.Call(m, "ExecuteLimitSell", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteLimitSell indicates an expected call of ExecuteLimitSell
func (mr *MockOrderManagerMockRecorder) ExecuteLimitSell(ctx, order interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteLimitSell", reflect.TypeOf((*MockOrderManager)(nil).ExecuteLimitSell), ctx, order)
}

// ExecuteLimitBuy mocks base method
func (m *MockOrderManager) ExecuteLimitBuy(ctx context.Context, order *LimitOrder) error {
	ret := m.ctrl.Call(m, "ExecuteLimitBuy", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecuteLimitBuy indicates an expected call of ExecuteLimitBuy
func (mr *MockOrderManagerMockRecorder) ExecuteLimitBuy(ctx, order interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecuteLimitBuy", reflect.TypeOf((*MockOrderManager)(nil).ExecuteLimitBuy), ctx, order)
}

// MockOrderMonitor is a mock of OrderMonitor interface
type MockOrderMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockOrderMonitorMockRecorder
}

// MockOrderMonitorMockRecorder is the mock recorder for MockOrderMonitor
type MockOrderMonitorMockRecorder struct {
	mock *MockOrderMonitor
}

// NewMockOrderMonitor creates a new mock instance
func NewMockOrderMonitor(ctrl *gomock.Controller) *MockOrderMonitor {
	mock := &MockOrderMonitor{ctrl: ctrl}
	mock.recorder = &MockOrderMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderMonitor) EXPECT() *MockOrderMonitorMockRecorder {
	return m.recorder
}

// Start mocks base method
func (m *MockOrderMonitor) Start(arg0 chan *OpenOrder) {
	m.ctrl.Call(m, "Start", arg0)
}

// Start indicates an expected call of Start
func (mr *MockOrderMonitorMockRecorder) Start(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOrderMonitor)(nil).Start), arg0)
}

// process mocks base method
func (m *MockOrderMonitor) process() {
	m.ctrl.Call(m, "process")
}

// process indicates an expected call of process
func (mr *MockOrderMonitorMockRecorder) process() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "process", reflect.TypeOf((*MockOrderMonitor)(nil).process))
}

// GetOrders mocks base method
func (m *MockOrderMonitor) GetOrders() []*OpenOrder {
	ret := m.ctrl.Call(m, "GetOrders")
	ret0, _ := ret[0].([]*OpenOrder)
	return ret0
}

// GetOrders indicates an expected call of GetOrders
func (mr *MockOrderMonitorMockRecorder) GetOrders() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderMonitor)(nil).GetOrders))
}

// Execute mocks base method
func (m *MockOrderMonitor) Execute(order *OpenOrder) (string, error) {
	ret := m.ctrl.Call(m, "Execute", order)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockOrderMonitorMockRecorder) Execute(order interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockOrderMonitor)(nil).Execute), order)
}
