package main

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/shopspring/decimal"
)

type OpenOrder struct {
	Type      string
	MarketKey string
	Quantity  decimal.Decimal
	Limit     decimal.Decimal
	UUID      string
}

type TraderState struct {
	Balances   map[string]decimal.Decimal
	OpenOrders []*OpenOrder
	Charts     map[string]*charts.CloudChart
}

func NewTraderState() *TraderState {
	return &TraderState{
		Balances:   make(map[string]decimal.Decimal),
		OpenOrders: []*OpenOrder{},
		Charts:     make(map[string]*charts.CloudChart),
	}
}

func (t *TraderState) SetBalance(marketKey string, amount decimal.Decimal) error {
	t.Balances[marketKey] = amount
	return nil
}

func (t *TraderState) GetBalance(marketKey string) decimal.Decimal {
	return t.Balances[marketKey]
}

func (t *TraderState) SetChart(marketKey string, chart *charts.CloudChart) error {
	t.Charts[marketKey] = chart
	return nil
}

func (t *TraderState) GetChart(marketKey string) *charts.CloudChart {
	return t.Charts[marketKey]
}
