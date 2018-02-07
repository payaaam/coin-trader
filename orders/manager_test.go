package orders

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/utils"
	//log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
	"time"
	//"time"
)

var BaseCurrency = "BTC"
var MarketCurrency = "LTC"
var MarketKey = "btc-ltc"
var someError = errors.New("some error")
var limit = "0.05"
var quantity = "10"
var orderID = "some-order-id"
var ctx = context.Background()
var MarketID = 1

type ManagerTestConfig struct {
	Exchange           *mocks.MockExchange
	OrderStore         *mocks.MockOrderStoreInterface
	MarketStore        *mocks.MockMarketStoreInterface
	OrderMonitor       *MockOrderMonitor
	OrderUpdateChannel chan *OpenOrder
}

func TestBuySuccess(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOpenOrderModel(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(nil)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(BuyOrder)

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

	bMap := manager.GetBalances()
	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")

	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
}

func TestBuyOrderStoreError(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOpenOrderModel(BuyOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(someError)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(BuyOrder)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("0.0", "0.0", "0.0", "0.0")

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOpenOrderModel(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(nil)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder, OpenOrderStatus)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return("", someError)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	err = manager.ExecuteLimitSell(ctx, &LimitOrder{
		BaseCurrency:   "BTC",
		MarketCurrency: "LTC",
		Limit:          utils.StringToDecimal("0.05"),
		Quantity:       utils.StringToDecimal("10"),
	})
	assert.Equal(t, someError, err, "should return errors from exchange")

	bMap := manager.GetBalances()
	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")

	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("2.0")), "should update available balance on filled order")
}

func TestSellOrderStoreError(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOpenOrderModel(SellOrder)

	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	orderMonitor.EXPECT().Execute(openOrderMatcher).Return(orderID, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(someError)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrderMatcher(SellOrder)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("0.0", "0.0", "0.0", "0.0")

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

func TestOpenOrderUpdateBuy(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	marketModel := getTestMarket()
	orderModelMatcher := getTestClosedOrderModel(BuyOrder, "0.045", quantity)

	balances := getTestBalances("1.5", "2.0", "2.0", "2.0")
	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	orderUpdateChannel <- &OpenOrder{
		Status: OpenOrderStatus,
	}

	time.Sleep(time.Millisecond * 10)
	bMap := manager.GetBalances()
	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("1.5")), "should not update available balance for open order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("2.0")), "should not update total balance for open order")

	orderUpdateChannel <- &OpenOrder{
		ID:             orderID,
		Type:           BuyOrder,
		Status:         FilledOrderStatus,
		MarketKey:      "btc-ltc",
		MarketCurrency: "LTC",
		BaseCurrency:   "BTC",
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
		QuantityFilled: utils.StringToDecimal(quantity),
		TradePrice:     utils.StringToDecimal("0.045"),
		CloseTimestamp: int64(100000000),
	}

	time.Sleep(time.Millisecond * 10)
	bMap = manager.GetBalances()
	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")

	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
}

func TestOpenOrderUpdateSell(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	marketModel := getTestMarket()
	orderModelMatcher := getTestClosedOrderModel(SellOrder, "0.045", quantity)

	balances := getTestBalances("2.0", "2.0", "1.5", "2.0")
	orderMonitor.EXPECT().Start(gomock.Any())
	exchange.EXPECT().GetBalances().Return(balances, nil)
	marketStore.EXPECT().GetMarket(ctx, "bittrex", MarketKey).Return(marketModel, nil)
	orderStore.EXPECT().Upsert(ctx, orderModelMatcher).Return(nil)

	err := manager.Setup()
	assert.Nil(t, err, "should not error")

	orderUpdateChannel <- &OpenOrder{
		Status: OpenOrderStatus,
	}

	time.Sleep(time.Millisecond * 10)
	bMap := manager.GetBalances()
	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("1.5")), "should not update available balance for open order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("2.0")), "should not update total balance for open order")

	orderUpdateChannel <- &OpenOrder{
		ID:             orderID,
		Type:           SellOrder,
		Status:         FilledOrderStatus,
		MarketKey:      "btc-ltc",
		MarketCurrency: "LTC",
		BaseCurrency:   "BTC",
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
		QuantityFilled: utils.StringToDecimal(quantity),
		TradePrice:     utils.StringToDecimal("0.045"),
		CloseTimestamp: int64(10000000),
	}

	time.Sleep(time.Millisecond * 10)
	bMap = manager.GetBalances()
	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")

	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
}

func TestGetBalanceNil(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := manager.GetBalance("xvc")
	assert.Equal(t, balances.Total, utils.ZeroDecimal())
	assert.Equal(t, balances.Available, utils.ZeroDecimal())
}
