package mocks

import (
	"github.com/payaaam/coin-trader/db/models"
)

type OrderModelMatcher struct {
	MarketID        int
	ExchangeOrderID string
	Limit           string
	Quantity        string
	Status          string
}

func (o *OrderModelMatcher) Matches(order interface{}) bool {
	toTest := order.(*models.Order)

	if toTest.ExchangeOrderID != o.ExchangeOrderID {
		return false
	}

	if toTest.MarketID != o.MarketID {
		return false
	}

	if toTest.Limit != o.Limit {
		return false
	}

	if toTest.Quantity != o.Quantity {
		return false
	}

	if toTest.Status != o.Status {
		return false
	}

	return true
}

func (o *OrderModelMatcher) String() string {
	return "does not match the order model"
}
