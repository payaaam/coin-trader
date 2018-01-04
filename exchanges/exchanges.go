package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
)

type Exchange interface {
	NewClient() *Exchange
	GetBitcoinMarkets() []*Market
	GetCandles(tradingPair string, interval string) (*charts.CloudChart, error)
}

type Market struct {
	TradingPair    string
	BaseCurrency   string
	MarketCurrency string
}
