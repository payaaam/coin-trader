package db

import (
	"github.com/payaaam/coin-trader/db/models"
	log "github.com/sirupsen/logrus"
)

type OrderModelMatcher struct {
	MarketID        int
	ExchangeOrderID string
	Limit           string
	Quantity        string
	Status          string
	Type            string
}

func (o *OrderModelMatcher) Matches(order interface{}) bool {
	toTest := order.(*models.Order)

	if toTest.ExchangeOrderID != o.ExchangeOrderID {
		log.Info("Blimit")
		return false
	}

	if toTest.MarketID != o.MarketID {
		log.Info("Alimit")
		log.Info(toTest.MarketID)
		return false
	}

	if toTest.Limit != o.Limit {
		log.Info("limit")
		return false
	}

	if toTest.Quantity != o.Quantity {
		log.Info("Xlimit")
		return false
	}

	if toTest.Status != o.Status {
		log.Info("Ylimit")
		return false
	}

	if toTest.Type != o.Type {
		log.Info("Zlimit")
		return false
	}

	return true
}

func (o *OrderModelMatcher) String() string {
	return "does not match the order model"
}
