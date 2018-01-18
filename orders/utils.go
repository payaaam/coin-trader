package orders

import (
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
