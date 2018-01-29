package orders

import (
	"github.com/payaaam/coin-trader/utils"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	//"time"
)

func TestExecuteBuy(t *testing.T) {
	mockConfig := newMockDependencies(t)
	exchange := mockConfig.Exchange
	orderUpdateChannel := mockConfig.OrderUpdateChannel

	monitor := NewTestMonitor(mockConfig.Exchange)
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

	monitor := NewTestMonitor(mockConfig.Exchange)
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

	monitor := NewTestMonitor(mockConfig.Exchange)
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

	monitor := NewTestMonitor(mockConfig.Exchange)
	monitor.Start(orderUpdateChannel)
	exchange.EXPECT().ExecuteLimitSell(MarketKey, utils.StringToDecimal(limit), utils.StringToDecimal(quantity)).Return("", someError)

	openOrder := getTestOpenOrder(SellOrder)

	_, err := monitor.Execute(openOrder)
	assert.Equal(t, someError, errors.Cause(err))
}

func TestOrderUpdateFilled(t *testing.T) {

}

func TestOrderUpdatePartiallyFilled(t *testing.T) {

}

func TestOrderTimeout(t *testing.T) {

}
