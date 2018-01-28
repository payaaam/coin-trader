package orders

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

var BaseCurrency = "BTC"
var MarketCurrency = "LTC"
var MarketKey = "btc-ltc"
var someError = errors.New("some error")
var limit = "0.05"
var quantity = "10"
var orderID = "some-order-id"
var ctx = context.Background()
var marketID = 1

func TestBuySuccess(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("BTC", "2.0")

	openOrderMatcher := getTestOpenOrder(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Save(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitBuy(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Nil(t, err, "should not return notEnoughFundsError")

	limitDecimal := utils.StringToDecimal(limit)
	quantityDecimal := utils.StringToDecimal(quantity)
	expectedBalance := utils.StringToDecimal("2").Sub(limitDecimal.Mul(quantityDecimal))
	actualBalance := manager.GetBalances()["btc"].Available
	assert.Equal(t, expectedBalance, actualBalance, "should deduct purchase price from balance")
}

func TestBuyExecuteError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("BTC", "2.0")

	openOrderMatcher := getTestOpenOrder(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return("", someError)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitBuy(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestBuyOrderStoreError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("BTC", "2.0")

	openOrderMatcher := getTestOpenOrder(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Save(ctx, orderModelMatcher).Return(someError)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitBuy(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestBuyMarketStoreError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("BTC", "2.0")

	openOrderMatcher := getTestOpenOrder(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(nil, someError)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitBuy(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestBuyInsufficentFunds(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("BTC", "0.0")

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitBuy(context.Background(), &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, ErrNotEnoughFunds, err, "should return notEnoughFundsError")
}

// SELL

func TestSellSuccess(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("LTC", "2.0")

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Save(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Nil(t, err, "should not return notEnoughFundsError")

	limitDecimal := utils.StringToDecimal(limit)
	quantityDecimal := utils.StringToDecimal(quantity)
	expectedBalance := utils.StringToDecimal("2").Sub(limitDecimal.Mul(quantityDecimal))
	actualBalance := manager.GetBalances()["ltc"].Available
	assert.Equal(t, expectedBalance, actualBalance, "should deduct purchase price from balance")
}

func TestSellExecuteError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("LTC", "2.0")

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return("", someError)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Save(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestSellOrderStoreError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("LTC", "2.0")

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Save(ctx, orderModelMatcher).Return(someError)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestSellMarketStoreError(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("LTC", "2.0")

	openOrderMatcher := getTestOpenOrder(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(nil, someError)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")
}

func TestSellInsufficentFunds(t *testing.T) {
	exchange, orderStore, marketStore, orderMonitor := newMockDependencies(t)
	manager := NewManager(orderMonitor, exchange, orderStore, marketStore)

	balances := getTestBalances("LTC", "0.0")

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(context.Background(), &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, ErrNotEnoughFunds, err, "should return notEnoughFundsError")
}

func getTestMarket() *models.Market {
	return &models.Market{
		ID: marketID,
	}
}

func getTestOrderModel(orderType string) *db.OrderModelMatcher {
	return &db.OrderModelMatcher{
		Type:            orderType,
		MarketID:        marketID,
		ExchangeOrderID: orderID,
		Limit:           limit,
		Quantity:        quantity,
		Status:          db.OpenOrderStatus,
	}
}

func getTestOpenOrder(orderType string) *OpenOrderMatcher {
	return &OpenOrderMatcher{
		Type:           orderType,
		BaseCurrency:   BaseCurrency,
		MarketCurrency: MarketCurrency,
		MarketKey:      MarketKey,
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
		Status:         OpenOrderStatus,
	}
}

func newMockDependencies(t *testing.T) (*mocks.MockExchange, *mocks.MockOrderStoreInterface, *mocks.MockMarketStoreInterface, *MockOrderMonitor) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockExchange := mocks.NewMockExchange(mockCtrl)
	mockOrderStore := mocks.NewMockOrderStoreInterface(mockCtrl)
	mockMarketStore := mocks.NewMockMarketStoreInterface(mockCtrl)
	mockOrderMonitor := NewMockOrderMonitor(mockCtrl)

	return mockExchange, mockOrderStore, mockMarketStore, mockOrderMonitor
}

func getTestBalances(currency string, balance string) []*exchanges.Balance {
	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: currency,
		Total:        utils.StringToDecimal(balance),
		Available:    utils.StringToDecimal(balance),
	})

	return balances
}
