package orders

import (
	"github.com/payaaam/coin-trader/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var closedTimestamp = int64(1234567)
var tradePrice = "2.1"
var partiallyFilledQuantity = "8.5"

func TestExecuteBuy(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return(orderID, nil)

	openOrder := getTestOpenOrder(BuyOrder)

	receivedOrderID, err := monitor.Execute(openOrder)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, orderID, receivedOrderID, "should receive orderID from exchange")

	orders := monitor.GetOrders()
	order := orders[0]
	assert.Equal(t, order.ID, orderID)
	assert.Equal(t, order.Status, OpenOrderStatus)
	assert.True(t, order.Limit.Equal(utils.StringToDecimal(limit)))
	assert.True(t, order.Quantity.Equal(utils.StringToDecimal(quantity)))
	assert.Equal(t, 1, len(orders), "should have 1 order")
}

func TestExecuteBuyError(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return("", someError)

	openOrder := getTestOpenOrder(BuyOrder)

	_, err := monitor.Execute(openOrder)
	assert.Equal(t, someError, errors.Cause(err))
}

func TestExecuteSell(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitSell(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return(orderID, nil)

	openOrder := getTestOpenOrder(SellOrder)

	receivedOrderID, err := monitor.Execute(openOrder)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, orderID, receivedOrderID, "should receive orderID from exchange")

	orders := monitor.GetOrders()
	order := orders[0]
	assert.Equal(t, order.ID, orderID)
	assert.Equal(t, order.Status, OpenOrderStatus)
	assert.True(t, order.Limit.Equal(utils.StringToDecimal(limit)))
	assert.True(t, order.Quantity.Equal(utils.StringToDecimal(quantity)))
	assert.Equal(t, 1, len(orders), "should have 1 order")
}

func TestExecuteSellError(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitSell(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return("", someError)

	openOrder := getTestOpenOrder(SellOrder)

	_, err := monitor.Execute(openOrder)
	assert.Equal(t, someError, errors.Cause(err))
}

func TestOrderUpdateFilled(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return(orderID, nil)

	exchangeOrder := getTestFilledExchangeOrder()
	exchange.EXPECT().GetOrder(orderID).Return(exchangeOrder, nil)
	openOrder := getTestOpenOrder(BuyOrder)

	receivedOrderID, err := monitor.Execute(openOrder)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, orderID, receivedOrderID, "should receive orderID from exchange")

	go monitor.process()
	receivedOrder := <-orderUpdateChannel

	assert.Equal(t, FilledOrderStatus, receivedOrder.Status)
	assert.Equal(t, exchangeOrder.CloseTimestamp, receivedOrder.CloseTimestamp)
	assert.Equal(t, exchangeOrder.TradePrice, receivedOrder.TradePrice)
	assert.Equal(t, exchangeOrder.QuantityFilled, receivedOrder.QuantityFilled)
	assert.Equal(t, orderID, receivedOrder.ID)

	monitor.process()
	assert.Equal(t, 0, len(monitor.GetOrders()))
}

func TestOrderUpdatePartiallyFilled(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return(orderID, nil)

	exchangeOrder := getTestPartiallyFilledExchangeOrder()
	exchange.EXPECT().GetOrder(orderID).Return(exchangeOrder, nil)
	openOrder := getTestOpenOrder(BuyOrder)

	receivedOrderID, err := monitor.Execute(openOrder)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, orderID, receivedOrderID, "should receive orderID from exchange")

	go monitor.process()
	receivedOrder := <-orderUpdateChannel

	assert.Equal(t, FilledOrderStatus, receivedOrder.Status)
	assert.Equal(t, exchangeOrder.CloseTimestamp, receivedOrder.CloseTimestamp)
	assert.Equal(t, exchangeOrder.TradePrice, receivedOrder.TradePrice)
	assert.Equal(t, exchangeOrder.QuantityFilled, receivedOrder.QuantityFilled)
	assert.Equal(t, orderID, receivedOrder.ID)

	// Ensure item is removed from array
	monitor.process()
	assert.Equal(t, 0, len(monitor.GetOrders()))

}

func TestOrderTimeoutSuccess(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitBuy(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return(orderID, nil)

	exchangeOrder := getTestUnfilledExchangeOrder()
	exchangeOrderClosed := getTestUnfilledClosedExchangeOrder()
	exchange.EXPECT().GetOrder(orderID).Return(exchangeOrder, nil)
	exchange.EXPECT().CancelOrder(orderID).Return(nil)
	exchange.EXPECT().GetOrder(orderID).Return(exchangeOrderClosed, nil)

	openOrder := getTestOpenOrder(BuyOrder)
	openOrder.OpenTimestamp = openOrder.OpenTimestamp - OrderOpenTimeoutMS

	receivedOrderID, err := monitor.Execute(openOrder)
	assert.Nil(t, err, "should not error")
	assert.Equal(t, orderID, receivedOrderID, "should receive orderID from exchange")

	go monitor.process()
	receivedOrder := <-orderUpdateChannel

	assert.Equal(t, FilledOrderStatus, receivedOrder.Status)
	assert.Equal(t, closedTimestamp, receivedOrder.CloseTimestamp)

	// Ensure order is removed from array
	monitor.process()
	assert.Equal(t, 0, len(monitor.GetOrders()))
}
