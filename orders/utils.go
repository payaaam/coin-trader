package orders

import (
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
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
	return &models.Order{
		Type:            BuyOrder,
		Limit:           order.Limit.String(),
		Quantity:        order.Quantity.String(),
		ExchangeOrderID: order.ID,
		Status:          db.OpenOrderStatus,
		OpenTime:        time.Now().Unix(),
	}
}
