package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/shopspring/decimal"
)

// All Exchanges must follow this interface
type Exchange interface {
	GetMarketKey(base string, market string) string
	GetBitcoinMarkets() ([]*Market, error)
	GetCandles(tradingPair string, interval string) ([]*charts.Candle, error)
	GetLatestCandle(tradingPair string, chartInterval string) (*charts.Candle, error)
	ExecuteLimitBuy(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error)
	ExecuteLimitSell(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error)
	GetBalances() ([]*Balance, error)
	GetOrder(orderID string) (*Order, error)
	CancelOrder(orderID string) error
}

type Market struct {
	ExchangeName       string
	MarketKey          string
	BaseCurrency       string
	BaseCurrencyName   string
	MarketCurrency     string
	MarketCurrencyName string
}

type Balance struct {
	BaseCurrency string
	Total        decimal.Decimal
	Available    decimal.Decimal
}

type Order struct {
	Type           string
	MarketKey      string
	OpenTimestamp  int64
	CloseTimestamp int64
	Quantity       decimal.Decimal
	Limit          decimal.Decimal
	TradePrice     decimal.Decimal
	QuantityFilled decimal.Decimal
	Status         string
	ID             string
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
