package orders

import (
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	"testing"
)

func NewMockManager(t *testing.T) OrderManager {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockExchange := exchanges.NewMockExchange(mockCtrl)
	mockOrderStore := db.NewMockOrderStoreInterface(mockCtrl)

	return NewManager(mockExchange, mockOrderStore)
}

func TestExecuteLimitBuyInsufficentFunds(t *testing.T) {
	manager := NewMockManager(t)

	var balances []*exchanges.Balance
	balances = append(balances, &exchanges.Balance{
		BaseCurrency: "BTC",
		Total:        utils.StringToDecimal("0.0"),
		Available:    utils.StringToDecimal("0.0"),
	})

	manager.client.EXPECT().GetBalances().Return(balances, nil)
}
