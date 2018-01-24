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

func NewMockDependencies(t *testing.T) (*mocks.MockExchange, *mocks.MockOrderStoreInterface, *mocks.MockMarketStoreInterface) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockExchange := mocks.NewMockExchange(mockCtrl)
	mockOrderStore := mocks.NewMockOrderStoreInterface(mockCtrl)
	mockMarketStore := mocks.NewMockMarketStoreInterface(mockCtrl)

	return mockExchange, mockOrderStore, mockMarketStore
}

func TestExecuteLimitBuySuccess(t *testing.T) {
	exchange, orderStore, marketStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore, marketStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("2.0"),
		Available:    utils.StringToDecimal("2.0"),
	})

	var limit = utils.StringToDecimal("0.05")
	var quantity = utils.StringToDecimal("10")
	var orderID = "some-order-id"
	var ctx = context.Background()
	marketModel := &models.Market{
		ID: 1,
	}

	orderModelMatcher := &mocks.OrderModelMatcher{
		MarketID:        marketModel.ID,
		ExchangeOrderID: orderID,
		Limit:           limit.String(),
		Quantity:        quantity.String(),
		Status:          db.OpenOrderStatus,
	}
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, limit, quantity).Return(orderID, nil)
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

	assert.Equal(t, 1, len(manager.GetOpenOrders()), "should have 1 open order")
	assert.Equal(t, orderID, manager.GetOpenOrders()[0].ID, "should have correct order id")

	expectedBalance := utils.StringToDecimal("2").Sub(limit.Mul(quantity))
	actualBalance := manager.GetBalances()["btc"].Available
	assert.Equal(t, expectedBalance, actualBalance, "should deduct purchase price from balance")
}

func TestExecuteLimitBuyExchangeError(t *testing.T) {
	exchange, orderStore, marketStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore, marketStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("2.0"),
		Available:    utils.StringToDecimal("2.0"),
	})

	var someError = errors.New("some error")
	var limit = utils.StringToDecimal("0.05")
	var quantity = utils.StringToDecimal("10")
	var ctx = context.Background()

	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, limit, quantity).Return("", someError)

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

func TestExecuteLimitBuyOrderStoreError(t *testing.T) {
	exchange, orderStore, marketStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore, marketStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("2.0"),
		Available:    utils.StringToDecimal("2.0"),
	})

	var someError = errors.New("some error")
	var limit = utils.StringToDecimal("0.05")
	var quantity = utils.StringToDecimal("10")
	var orderID = "some-order-id"
	var ctx = context.Background()
	marketModel := &models.Market{
		ID: 1,
	}

	orderModelMatcher := &mocks.OrderModelMatcher{
		MarketID:        marketModel.ID,
		ExchangeOrderID: orderID,
		Limit:           limit.String(),
		Quantity:        quantity.String(),
		Status:          db.OpenOrderStatus,
	}
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, limit, quantity).Return(orderID, nil)
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

func TestExecuteLimitBuyMarketStoreError(t *testing.T) {
	exchange, orderStore, marketStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore, marketStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("2.0"),
		Available:    utils.StringToDecimal("2.0"),
	})

	var someError = errors.New("some error")
	var limit = utils.StringToDecimal("0.05")
	var quantity = utils.StringToDecimal("10")
	var orderID = "some-order-id"
	var ctx = context.Background()
	exchange.EXPECT().GetBalances().Return(balances, nil)
	exchange.EXPECT().GetMarketKey(BaseCurrency, MarketCurrency).Return(MarketKey)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, limit, quantity).Return(orderID, nil)
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
func TestExecuteLimitBuyInsufficentFunds(t *testing.T) {
	exchange, orderStore, marketStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore, marketStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("0.0"),
		Available:    utils.StringToDecimal("0.0"),
	})

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
