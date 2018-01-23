package orders

import (
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"testing"
)

func NewMockDependencies(t *testing.T) (*mocks.MockExchange, *mocks.MockOrderStoreInterface) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockExchange := mocks.NewMockExchange(mockCtrl)
	mockOrderStore := mocks.NewMockOrderStoreInterface(mockCtrl)

	return mockExchange, mockOrderStore
}

func TestExecuteLimitBuySuccess(t *testing.T) {
	exchange, orderStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("2.0"),
		Available:    utils.StringToDecimal("2.0"),
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
	assert.Nil(t, err, "should not return notEnoughFundsError")
}

func TestExecuteLimitBuyInsufficentFunds(t *testing.T) {
	exchange, orderStore := NewMockDependencies(t)
	manager := NewManager(exchange, orderStore)

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
