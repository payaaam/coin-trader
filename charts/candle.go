package charts

import (
	"github.com/shopspring/decimal"
)

type Candle struct {
	TimeStamp int64
	Day       int
	Open      decimal.Decimal
	Close     decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Volume    decimal.Decimal
	Kijun     decimal.Decimal
	Tenkan    decimal.Decimal
}
