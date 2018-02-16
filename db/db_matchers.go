package db

import (
	"github.com/payaaam/coin-trader/db/models"
	log "github.com/sirupsen/logrus"
	"gopkg.in/volatiletech/null.v6"
)

type OrderModelMatcher struct {
	MarketID        int
	ExchangeOrderID string
	Limit           string
	Quantity        string
	Status          string
	Type            string
	QuantityFilled  null.String
	TradePrice      null.String
}

func (o *OrderModelMatcher) Matches(order interface{}) bool {
	toTest := order.(*models.Order)

	if toTest.ExchangeOrderID != o.ExchangeOrderID {
		log.Info("MATCH: OrderID")
		return false
	}

	if toTest.MarketID != o.MarketID {
		log.Info("MATCH: MarketID")
		return false
	}

	if toTest.Limit != o.Limit {
		log.Info("MATCH: LIMIT")
		return false
	}

	if toTest.Quantity != o.Quantity {
		log.Info("MATCH: Quantity")
		return false
	}

	if toTest.QuantityFilled != o.QuantityFilled {
		log.Info(toTest.QuantityFilled)
		log.Info(o.QuantityFilled)
		log.Info("MATCH: QuantityFilled")
		return false
	}

	if toTest.Status != o.Status {
		log.Info("MATCH: Status")
		return false
	}

	if toTest.Type != o.Type {
		log.Info("MATCH: Type")
		return false
	}

	return true
}

func (o *OrderModelMatcher) String() string {
	return "does not match the order model"
}
