package orders

import (
	"github.com/shopspring/decimal"
	//log "github.com/sirupsen/logrus"
)

type OpenOrderMatcher struct {
	Type           string
	MarketKey      string
	BaseCurrency   string
	MarketCurrency string
	Quantity       decimal.Decimal
	Limit          decimal.Decimal
	Status         string
}

func (o *OpenOrderMatcher) Matches(order interface{}) bool {
	toTest := order.(*OpenOrder)

	if toTest.Type != o.Type {
		return false
	}

	if toTest.MarketKey != o.MarketKey {
		return false
	}

	if toTest.BaseCurrency != o.BaseCurrency {
		return false
	}

	if toTest.MarketCurrency != o.MarketCurrency {
		return false
	}

	if toTest.Limit.Equal(o.Limit) == false {
		return false
	}

	if toTest.Quantity.Equals(o.Quantity) == false {
		return false
	}

	if toTest.Status != o.Status {
		return false
	}

	if toTest.OpenTimestamp == 0 {
		return false
	}

	return true
}

func (o *OpenOrderMatcher) String() string {
	return "does not match the open order"
}

type LimitOrderMatcher struct {
	Limit          decimal.Decimal
	Quantity       decimal.Decimal
	MarketCurrency string
	BaseCurrency   string
}

func (o *LimitOrderMatcher) Matches(order interface{}) bool {
	toTest := order.(*LimitOrder)

	if toTest.BaseCurrency != o.BaseCurrency {
		return false
	}

	if toTest.MarketCurrency != o.MarketCurrency {
		return false
	}

	if toTest.Limit.Equal(o.Limit) == false {
		return false
	}

	if toTest.Quantity.Equals(o.Quantity) == false {
		return false
	}
	return true

}

func (o *LimitOrderMatcher) String() string {
	return "does not match the limit order"
}
