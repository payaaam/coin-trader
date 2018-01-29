package orders

import (
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/utils"
	"testing"
	"time"
)

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

func getTestOpenOrderMatcher(orderType string) *OpenOrderMatcher {
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

func getTestOpenOrder(orderType string) *OpenOrder {
	return &OpenOrder{
		OpenTimestamp:  time.Now().Unix(),
		Type:           orderType,
		Status:         OpenOrderStatus,
		BaseCurrency:   BaseCurrency,
		MarketCurrency: MarketCurrency,
		MarketKey:      MarketKey,
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
	}
}
