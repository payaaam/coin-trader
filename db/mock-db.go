// Code generated by MockGen. DO NOT EDIT.
// Source: ./db/db.go

// Package db is a generated GoMock package.
package db

import (
	gomock "github.com/golang/mock/gomock"
	charts "github.com/payaaam/coin-trader/charts"
	models "github.com/payaaam/coin-trader/db/models"
	context "golang.org/x/net/context"
	reflect "reflect"
)

// MockOrderStoreInterface is a mock of OrderStoreInterface interface
type MockOrderStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockOrderStoreInterfaceMockRecorder
}

// MockOrderStoreInterfaceMockRecorder is the mock recorder for MockOrderStoreInterface
type MockOrderStoreInterfaceMockRecorder struct {
	mock *MockOrderStoreInterface
}

// NewMockOrderStoreInterface creates a new mock instance
func NewMockOrderStoreInterface(ctrl *gomock.Controller) *MockOrderStoreInterface {
	mock := &MockOrderStoreInterface{ctrl: ctrl}
	mock.recorder = &MockOrderStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderStoreInterface) EXPECT() *MockOrderStoreInterfaceMockRecorder {
	return m.recorder
}

// Save mocks base method
func (m *MockOrderStoreInterface) Save(ctx context.Context, order *models.Order) error {
	ret := m.ctrl.Call(m, "Save", ctx, order)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockOrderStoreInterfaceMockRecorder) Save(ctx, order interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOrderStoreInterface)(nil).Save), ctx, order)
}

// MockChartStoreInterface is a mock of ChartStoreInterface interface
type MockChartStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockChartStoreInterfaceMockRecorder
}

// MockChartStoreInterfaceMockRecorder is the mock recorder for MockChartStoreInterface
type MockChartStoreInterfaceMockRecorder struct {
	mock *MockChartStoreInterface
}

// NewMockChartStoreInterface creates a new mock instance
func NewMockChartStoreInterface(ctrl *gomock.Controller) *MockChartStoreInterface {
	mock := &MockChartStoreInterface{ctrl: ctrl}
	mock.recorder = &MockChartStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockChartStoreInterface) EXPECT() *MockChartStoreInterfaceMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockChartStoreInterface) Upsert(ctx context.Context, chart *models.Chart) error {
	ret := m.ctrl.Call(m, "Upsert", ctx, chart)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockChartStoreInterfaceMockRecorder) Upsert(ctx, chart interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockChartStoreInterface)(nil).Upsert), ctx, chart)
}

// Save mocks base method
func (m *MockChartStoreInterface) Save(ctx context.Context, chart *models.Chart) error {
	ret := m.ctrl.Call(m, "Save", ctx, chart)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockChartStoreInterfaceMockRecorder) Save(ctx, chart interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockChartStoreInterface)(nil).Save), ctx, chart)
}

// MockMarketStoreInterface is a mock of MarketStoreInterface interface
type MockMarketStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockMarketStoreInterfaceMockRecorder
}

// MockMarketStoreInterfaceMockRecorder is the mock recorder for MockMarketStoreInterface
type MockMarketStoreInterfaceMockRecorder struct {
	mock *MockMarketStoreInterface
}

// NewMockMarketStoreInterface creates a new mock instance
func NewMockMarketStoreInterface(ctrl *gomock.Controller) *MockMarketStoreInterface {
	mock := &MockMarketStoreInterface{ctrl: ctrl}
	mock.recorder = &MockMarketStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketStoreInterface) EXPECT() *MockMarketStoreInterfaceMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockMarketStoreInterface) Upsert(ctx context.Context, market *models.Market) error {
	ret := m.ctrl.Call(m, "Upsert", ctx, market)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockMarketStoreInterfaceMockRecorder) Upsert(ctx, market interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockMarketStoreInterface)(nil).Upsert), ctx, market)
}

// Save mocks base method
func (m *MockMarketStoreInterface) Save(ctx context.Context, market *models.Market) error {
	ret := m.ctrl.Call(m, "Save", ctx, market)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockMarketStoreInterfaceMockRecorder) Save(ctx, market interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMarketStoreInterface)(nil).Save), ctx, market)
}

// GetMarkets mocks base method
func (m *MockMarketStoreInterface) GetMarkets(ctx context.Context, exchange string) ([]*models.Market, error) {
	ret := m.ctrl.Call(m, "GetMarkets", ctx, exchange)
	ret0, _ := ret[0].([]*models.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarkets indicates an expected call of GetMarkets
func (mr *MockMarketStoreInterfaceMockRecorder) GetMarkets(ctx, exchange interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarkets", reflect.TypeOf((*MockMarketStoreInterface)(nil).GetMarkets), ctx, exchange)
}

// GetMarket mocks base method
func (m *MockMarketStoreInterface) GetMarket(ctx context.Context, exchangeName, marketKey string) (*models.Market, error) {
	ret := m.ctrl.Call(m, "GetMarket", ctx, exchangeName, marketKey)
	ret0, _ := ret[0].(*models.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarket indicates an expected call of GetMarket
func (mr *MockMarketStoreInterfaceMockRecorder) GetMarket(ctx, exchangeName, marketKey interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarket", reflect.TypeOf((*MockMarketStoreInterface)(nil).GetMarket), ctx, exchangeName, marketKey)
}

// MockTickStoreInterface is a mock of TickStoreInterface interface
type MockTickStoreInterface struct {
	ctrl     *gomock.Controller
	recorder *MockTickStoreInterfaceMockRecorder
}

// MockTickStoreInterfaceMockRecorder is the mock recorder for MockTickStoreInterface
type MockTickStoreInterfaceMockRecorder struct {
	mock *MockTickStoreInterface
}

// NewMockTickStoreInterface creates a new mock instance
func NewMockTickStoreInterface(ctrl *gomock.Controller) *MockTickStoreInterface {
	mock := &MockTickStoreInterface{ctrl: ctrl}
	mock.recorder = &MockTickStoreInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTickStoreInterface) EXPECT() *MockTickStoreInterfaceMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockTickStoreInterface) Upsert(ctx context.Context, chartID int, candle *charts.Candle) error {
	ret := m.ctrl.Call(m, "Upsert", ctx, chartID, candle)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockTickStoreInterfaceMockRecorder) Upsert(ctx, chartID, candle interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTickStoreInterface)(nil).Upsert), ctx, chartID, candle)
}

// Save mocks base method
func (m *MockTickStoreInterface) Save(ctx context.Context, tick *models.Tick) error {
	ret := m.ctrl.Call(m, "Save", ctx, tick)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockTickStoreInterfaceMockRecorder) Save(ctx, tick interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTickStoreInterface)(nil).Save), ctx, tick)
}

// GetAllChartCandles mocks base method
func (m *MockTickStoreInterface) GetAllChartCandles(ctx context.Context, marketKey, exchange, interval string) ([]*charts.Candle, error) {
	ret := m.ctrl.Call(m, "GetAllChartCandles", ctx, marketKey, exchange, interval)
	ret0, _ := ret[0].([]*charts.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllChartCandles indicates an expected call of GetAllChartCandles
func (mr *MockTickStoreInterfaceMockRecorder) GetAllChartCandles(ctx, marketKey, exchange, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllChartCandles", reflect.TypeOf((*MockTickStoreInterface)(nil).GetAllChartCandles), ctx, marketKey, exchange, interval)
}

// GetChartCandles mocks base method
func (m *MockTickStoreInterface) GetChartCandles(ctx context.Context, marketKey, exchange, interval string) ([]*charts.Candle, error) {
	ret := m.ctrl.Call(m, "GetChartCandles", ctx, marketKey, exchange, interval)
	ret0, _ := ret[0].([]*charts.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChartCandles indicates an expected call of GetChartCandles
func (mr *MockTickStoreInterfaceMockRecorder) GetChartCandles(ctx, marketKey, exchange, interval interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChartCandles", reflect.TypeOf((*MockTickStoreInterface)(nil).GetChartCandles), ctx, marketKey, exchange, interval)
}

// GetLatestChartCandle mocks base method
func (m *MockTickStoreInterface) GetLatestChartCandle(ctx context.Context, chartID int) (*charts.Candle, error) {
	ret := m.ctrl.Call(m, "GetLatestChartCandle", ctx, chartID)
	ret0, _ := ret[0].(*charts.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestChartCandle indicates an expected call of GetLatestChartCandle
func (mr *MockTickStoreInterfaceMockRecorder) GetLatestChartCandle(ctx, chartID interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestChartCandle", reflect.TypeOf((*MockTickStoreInterface)(nil).GetLatestChartCandle), ctx, chartID)
}

// GetCandlesFromRange mocks base method
func (m *MockTickStoreInterface) GetCandlesFromRange(ctx context.Context, chartID int, start, end int64) ([]*charts.Candle, error) {
	ret := m.ctrl.Call(m, "GetCandlesFromRange", ctx, chartID, start, end)
	ret0, _ := ret[0].([]*charts.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCandlesFromRange indicates an expected call of GetCandlesFromRange
func (mr *MockTickStoreInterfaceMockRecorder) GetCandlesFromRange(ctx, chartID, start, end interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCandlesFromRange", reflect.TypeOf((*MockTickStoreInterface)(nil).GetCandlesFromRange), ctx, chartID, start, end)
}
