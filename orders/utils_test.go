package orders

import (
	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/utils"
	//log "github.com/sirupsen/logrus"
	"gopkg.in/volatiletech/null.v6"
	"testing"
	"time"
)

func getTestMarket() *models.Market {
	return &models.Market{
		ID: MarketID,
	}
}

func getTestOpenOrderModel(orderType string) *db.OrderModelMatcher {

	o := &db.OrderModelMatcher{
		Type:            orderType,
		MarketID:        MarketID,
		ExchangeOrderID: orderID,
		Limit:           limit,
		Quantity:        quantity,
		Status:          OpenOrderStatus,
	}

	if orderType == SellOrder {
		o.Quantity = quantitySell
	}
	return o
}

func getTestClosedOrderModel(orderType string, tradePrice string, quantityFilled string) *db.OrderModelMatcher {
	return &db.OrderModelMatcher{
		Type:            orderType,
		MarketID:        MarketID,
		ExchangeOrderID: orderID,
		Limit:           limit,
		Quantity:        quantity,
		QuantityFilled:  null.StringFrom(quantityFilled),
		TradePrice:      null.StringFrom(tradePrice),
		Status:          FilledOrderStatus,
	}
}

func getTestOrderModel(orderType string, orderStatus string) *db.OrderModelMatcher {
	return &db.OrderModelMatcher{
		Type:            orderType,
		MarketID:        MarketID,
		ExchangeOrderID: orderID,
		Limit:           limit,
		Quantity:        quantity,
		QuantityFilled:  null.StringFrom(quantity),
		TradePrice:      null.StringFrom(limit),
		Status:          orderStatus,
	}
}

func getTestOpenOrderMatcher(orderType string) *OpenOrderMatcher {
	o := &OpenOrderMatcher{
		Type:           orderType,
		BaseCurrency:   BaseCurrency,
		MarketCurrency: MarketCurrency,
		MarketKey:      MarketKey,
		Limit:          utils.StringToDecimal(limit),
		Quantity:       utils.StringToDecimal(quantity),
		Status:         OpenOrderStatus,
	}

	if orderType == SellOrder {
		o.Quantity = utils.StringToDecimal(quantitySell)
	}

	return o
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

func getTestFilledExchangeOrder() *exchanges.Order {
	return &exchanges.Order{
		CloseTimestamp: closedTimestamp,
		TradePrice:     utils.StringToDecimal(tradePrice),
		QuantityFilled: utils.StringToDecimal(quantity),
	}
}

func getTestPartiallyFilledExchangeOrder() *exchanges.Order {
	return &exchanges.Order{
		CloseTimestamp: closedTimestamp,
		TradePrice:     utils.StringToDecimal(tradePrice),
		QuantityFilled: utils.StringToDecimal(partiallyFilledQuantity),
	}
}

func getTestUnfilledExchangeOrder() *exchanges.Order {
	return &exchanges.Order{}
}

func getTestUnfilledClosedExchangeOrder() *exchanges.Order {
	return &exchanges.Order{
		CloseTimestamp: closedTimestamp,
		TradePrice:     utils.StringToDecimal(tradePrice),
		QuantityFilled: utils.ZeroDecimal(),
	}
}
