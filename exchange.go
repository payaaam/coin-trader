package main

import (
	"github.com/payaaam/coin-trader/charts"
)

type Exchange interface {
	NewClient() *Exchange
	GetCandles(coin string) (*charts.CloudChart, error)
}
