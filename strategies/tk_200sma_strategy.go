package strategies

import (
	"github.com/payaaam/coin-trader/charts"
	//"github.com/shopspring/decimal"
)

type TKCross200SMA struct{}

func NewTKCross200SMAStrategy() *TKCross200SMA {
	return &TKCross200SMA{}
}

func (c *TKCross200SMA) ShouldBuy(chart *charts.CloudChart) bool {
	var lastCandle = chart.GetLastCandle()

	if chart.Test == false && chart.GetCandleCount() < 120 {
		return false
	}

	// Buy on bullish TK cross when price > 200SMA
	if lastCandle.Tenkan.GreaterThan(lastCandle.Kijun) && lastCandle.Close.GreaterThan(lastCandle.SMA200) {
		return true
	}

	return false

}

func (c *TKCross200SMA) ShouldSell(chart *charts.CloudChart) bool {
	var lastCandle = chart.GetLastCandle()

	if chart.Test == false && chart.GetCandleCount() < 120 {
		return false
	}

	// Sell on bearish TK cross when price > 200SMA
	if lastCandle.Kijun.GreaterThan(lastCandle.Tenkan) {
		return true
	}

	return false
}
