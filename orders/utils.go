package orders

import (
	"github.com/payaaam/coin-trader/db/models"
	"github.com/shopspring/decimal"
)

func hasAvailableFunds(balance decimal.Decimal, order *LimitOrder) bool {
	totalPrice := order.Amount.Mul(order.Limit)
	newBalance := balance.Sub(totalPrice)

	if balance.Equals(utils.ZeroDecimal()) || newBalance.Sign() == -1 {
		return false
	}

	return true
}

func convertToOrderModel(order *OpenOrder) *models.Order {
	return &models.Order{
		Type:            BuyOrder,
		Limit:           order.Limit,
		Quantity:        order.Quantity,
		ExchangeOrderID: order.ID,
		Status:          db.OpenOrderStatus,
		OpenTime:        time.Now().Unix(),
	}
}
