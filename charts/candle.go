package charts

import (
	"fmt"
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

func (c *Candle) Print() {
	fmt.Println("---- Candle ----")
	fmt.Printf("TimeStamp: %v\n", c.TimeStamp)
	fmt.Printf("Open: %v\n", c.Open)
	fmt.Printf("Close: %v\n", c.Close)
	fmt.Printf("High: %v\n", c.High)
	fmt.Printf("Low: %v\n", c.Low)
	fmt.Printf("Day: %v\n", c.Day)
	fmt.Printf("Tenkan: %v\n", c.Tenkan)
	fmt.Printf("Kijun: %v\n", c.Kijun)
}
