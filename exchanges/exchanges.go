package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
)

type Exchange interface {
	NewClient() *Exchange
	GetBitcoinMarkets() []*Market
	GetCandles(tradingPair string, interval string) (*charts.CloudChart, error)
}

type Market struct {
	ExchangeName       string
	MarketKey          string
	BaseCurrency       string
	BaseCurrencyName   string
	MarketCurrency     string
	MarketCurrencyName string
}

var Intervals = map[string]map[string]string{
	"bittrex": map[string]string{
		"thirtyMin": db.ThirtyMinuteInterval,
		"hour":      db.OneDayInterval,
		"day":       db.OneDayInterval,
	},
}
