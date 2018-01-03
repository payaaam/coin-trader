package main

import (
	"github.com/payaaam/coin-trader/charts"
)

// Exchange under exchange folder. All exchange clients must abide by this interface
type Exchange interface {
	NewClient() *Exchange
	GetCandles(coin string) (*charts.CloudChart, error)
}
