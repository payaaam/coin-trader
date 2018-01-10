package strategies

import (
	"github.com/payaaam/coin-trader/charts"
	//"github.com/shopspring/decimal"
)

const BullishCross = 1
const BearishCross = -1

type CloudStrategy struct{}

func NewCloudStrategy() *CloudStrategy {
	return &CloudStrategy{}
}

func (c *CloudStrategy) ShouldBuy(chart *charts.CloudChart) bool {
	periodsSinceLastCross := findLastTKCross(chart, BullishCross)
	if periodsSinceLastCross == 0 || periodsSinceLastCross == 1 {
		return true
	}
	return false
	// If String Bullish TK Cross just occured
	// -- TK Cross + Price Above cloud
	// -- -- YES

	// If Normal Bullish TK Cross just occured
	// -- TK Cross + Price in Cloud
	// -- -- YES

	// If Weak Bullish TK Cross Just Occured
	// -- TK Cross + Below Cloud
	// -- Check slope of TK Cross
	// -- If Slope is aggressive
	// -- -- YES
	// -- If flat slope
	// -- -- NO

	// Slope Calculation
	// -- Calculate next cross,
	// -- -- if its -10 < x Low Slope
	// -- -- if its -2 < x < 0 High Slope?
	return false
}

func (c *CloudStrategy) ShouldSell(chart *charts.CloudChart) bool {
	periodsSinceLastCross := findLastTKCross(chart, BearishCross)
	if periodsSinceLastCross == 0 || periodsSinceLastCross == 1 {
		return true
	}
	return false
	// If Bearish TK Cross just occured
	// -- TK Cross + Price Above / In / Below cloud
	// -- -- YES

	return false
}

// crossType == -1 - Bearish Cross
// crossType == 1 - Bullish Cross
func findLastTKCross(chart *charts.CloudChart, crossType int) int {
	candles := chart.GetCandles()

	chartLength := len(candles) - 1
	var periodsSinceLastCross = 0

	for day := chartLength; day >= 0; day-- {
		if day-1 == 0 {
			break
		}

		candle := candles[day]
		previousCandle := candles[day-1]

		currentDiff := candle.Tenkan.Sub(candle.Kijun)
		previousDiff := previousCandle.Tenkan.Sub(previousCandle.Kijun)

		if currentDiff.Sign() != previousDiff.Sign() && currentDiff.Sign() == crossType {
			break
		}
		periodsSinceLastCross = periodsSinceLastCross + 1
	}

	return periodsSinceLastCross
}
