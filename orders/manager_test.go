package orders

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
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
var marketID = 1

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

	openOrderMatcher := getTestOpenOrder(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(BuyOrder)

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
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	balances := getTestBalances("2.0", "2.0", "2.0", "2.0")

	openOrderMatcher := getTestOpenOrder(BuyOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(BuyOrder)

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

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

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

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

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

	openOrderMatcher := getTestOpenOrder(SellOrder)
	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(SellOrder)

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

func TestOpenOrderUpdate(t *testing.T) {
	mockConfig := newMockDependencies(t)
	orderMonitor := mockConfig.OrderMonitor
	exchange := mockConfig.Exchange
	orderStore := mockConfig.OrderStore
	marketStore := mockConfig.MarketStore
	orderUpdateChannel := mockConfig.OrderUpdateChannel
	manager := NewManager(orderMonitor, orderUpdateChannel, exchange, orderStore, marketStore)

	marketModel := getTestMarket()
	orderModelMatcher := getTestOrderModel(BuyOrder)

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
	//assert.Equal(t, utils.StringToDecimal("1.5"), bMap["btc"].Available, "should not update available balance for open order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("2.0")), "should not update total balance for open order")
	//assert.Equal(t, utils.StringToDecimal("2.0"), bMap["btc"].Total, "should not update total balance for open order")

	orderUpdateChannel <- &OpenOrder{
		ID:             orderID,
		Type:           BuyOrder,
		Status:         FilledOrderStatus,
		MarketKey:      "btc-ltc",
		MarketCurrency: "LTC",
		BaseCurrency:   "BTC",
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
		TradePrice:     utils.StringToDecimal("0.045"),
	}

	time.Sleep(time.Millisecond * 10)
	bMap = manager.GetBalances()
	assert.True(t, bMap["btc"].Available.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")
	assert.True(t, bMap["btc"].Total.Equals(utils.StringToDecimal("1.55")), "should update available balance on filled order")

	assert.True(t, bMap["ltc"].Available.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
	assert.True(t, bMap["ltc"].Total.Equals(utils.StringToDecimal("12.0")), "should update available balance on filled order")
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

func newMockDependencies(t *testing.T) *ManagerTestConfig {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return &ManagerTestConfig{
		Exchange:           mocks.NewMockExchange(mockCtrl),
		OrderStore:         mocks.NewMockOrderStoreInterface(mockCtrl),
		MarketStore:        mocks.NewMockMarketStoreInterface(mockCtrl),
		OrderMonitor:       NewMockOrderMonitor(mockCtrl),
		OrderUpdateChannel: make(chan *OpenOrder),
	}
}

func getTestBalances(baseAvailable, baseTotal, marketAvailable, marketTotal string) []*exchanges.Balance {
	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: MarketCurrency,
		Total:        utils.StringToDecimal(marketTotal),
		Available:    utils.StringToDecimal(marketAvailable),
	})

	balances = append(balances, &exchanges.Balance{
		BaseCurrency: BaseCurrency,
		Total:        utils.StringToDecimal(baseTotal),
		Available:    utils.StringToDecimal(baseAvailable),
	})

	return balances
}
