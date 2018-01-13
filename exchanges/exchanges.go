package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
)

// All Exchanges must follow this interface
type Exchange interface {
	GetBitcoinMarkets() ([]*Market, error)
	GetCandles(tradingPair string, interval string) ([]*charts.Candle, error)
	GetLatestCandle(tradingPair string, chartInterval string) (*charts.Candle, error)
	ExecuteLimitBuy(tradingPair string, price string, quantity string) (string, error)
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
		db.ThirtyMinuteInterval: "thirtyMin",
		db.OneHourInterval:      "hour",
		db.OneDayInterval:       "day",
	},
}

var ValidExchanges = map[string]bool{
	"bittrex": true,
}
