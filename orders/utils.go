package orders

import (
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

func hasAvailableFunds(balance decimal.Decimal, order *LimitOrder) bool {
	totalPrice := order.Quantity.Mul(order.Limit)
	newBalance := balance.Sub(totalPrice)

	if balance.Equals(utils.ZeroDecimal()) || newBalance.Sign() == -1 {
		return false
	}

	return true
}

func convertToOrderModel(order *OpenOrder) *models.Order {
	orderModel := &models.Order{
		Type:            order.Type,
		Limit:           order.Limit.String(),
		Quantity:        order.Quantity.String(),
		ExchangeOrderID: order.ID,
		Status:          order.Status,
		OpenTime:        time.Now().Unix(),
	}

	if order.CloseTimestamp != 0 {
		orderModel.CloseTime = null.Int64From(order.CloseTimestamp)
		orderModel.SellPrice = null.StringFrom(order.TradePrice.String())
		orderModel.QuantityFilled = null.StringFrom(order.QuantityFilled.String())
	}

	return orderModel
}
