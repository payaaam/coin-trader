package strategies

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	//"github.com/shopspring/decimal"
	"errors"
)

const BullishCross = 1
const BearishCross = -1
const InCloud = 0
const AboveCloud = 0
const BelowCloud = 0
const NoCrossNumber = 1000000

type CloudStrategy struct{}

func NewCloudStrategy() *CloudStrategy {
	return &CloudStrategy{}
}

func (c *CloudStrategy) ShouldBuy(chart *charts.CloudChart) bool {
	if chart.GetCandleCount() < 120 {
		return false
	}
	periodsSinceLastCross := findLastTKCross(chart, BullishCross)
	//log.Infof("Buy SINCE LAST: %v", periodsSinceLastCross)
	if periodsSinceLastCross == NoCrossNumber {
		return false
	}

	if periodsSinceLastCross == 0 {

		hasStrongTenkan := hasStrongTenkanSlope(chart)
		//log.Infof("STRONG TENKAN: %v", hasStrongTenkan)

		cloudPosition, err := getPriceCloudPosition(chart)
		if err != nil {
			log.Error(err)
			return false
		}
		//log.Infof("Cloud Position: %v", cloudPosition)

		// If String Bullish TK Cross just occured
		// -- TK Cross + Price Above cloud
		// -- -- YES
		if cloudPosition == AboveCloud {
			return true
		}

		// If Normal Bullish TK Cross just occured
		// -- TK Cross + Price in Cloud
		// -- -- YES
		if cloudPosition == InCloud {
			return true
		}

		// If Weak Bullish TK Cross Just Occured
		// -- TK Cross + Below Cloud
		// -- Check slope of TK Cross
		// -- If Slope is aggressive
		// -- -- YES
		// -- If flat slope
		// -- -- NO
		if cloudPosition == BelowCloud {

			if hasStrongTenkan == true {
				return true
			}
		}

	}
	return false

}

func (c *CloudStrategy) ShouldSell(chart *charts.CloudChart) bool {
	if chart.GetCandleCount() < 120 {
		return false
	}
	periodsSinceLastCross := findLastTKCross(chart, BearishCross)
	if periodsSinceLastCross == NoCrossNumber {
		return false
	}

	// If Bearish TK Cross just occured
	// -- TK Cross + Price Above / In / Below cloud
	// -- -- YES
	if periodsSinceLastCross == 0 {
		return true
	}

	return false
}

// crossType == -1 - Bearish Cross
// crossType == 1 - Bullish Cross
func findLastTKCross(chart *charts.CloudChart, crossType int) int {
	candles := chart.GetCandles()

	chartLength := len(candles) - 1
	var periodsSinceLastCross = 0
	var hasCrossed = false

	for day := chartLength; day >= 0; day-- {
		if day-1 < 0 {
			break
		}

		candle := candles[day]
		previousCandle := candles[day-1]

		currentDiff := candle.Tenkan.Sub(candle.Kijun)
		previousDiff := previousCandle.Tenkan.Sub(previousCandle.Kijun)
		if currentDiff.Sign() != previousDiff.Sign() && currentDiff.Sign() == crossType {
			hasCrossed = true
			break
		}
		periodsSinceLastCross = periodsSinceLastCross + 1
	}
	if hasCrossed == true {
		return periodsSinceLastCross
	}

	return NoCrossNumber
}

// Determines if the latest close price is above the cloud
func getPriceCloudPosition(chart *charts.CloudChart) (int, error) {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1

	if chartLength < 3 {
		return 0, errors.New("not enough candles")
	}

	lastCandle := candles[chartLength]
	candleCloud, err := chart.GetCloud(lastCandle.Day)
	if err != nil {
		return 0, err
	}

	// Above the Cloud
	if lastCandle.Close.GreaterThan(candleCloud.SenkouB) && lastCandle.Close.GreaterThan(candleCloud.SenkouA) {
		return AboveCloud, nil
	}

	// In a Bullish Cloud
	if candleCloud.SenkouA.GreaterThan(lastCandle.Close) && lastCandle.Close.GreaterThan(candleCloud.SenkouB) {
		return InCloud, nil
	}

	// In a Bearish Cloud
	if candleCloud.SenkouB.GreaterThan(lastCandle.Close) && lastCandle.Close.GreaterThan(candleCloud.SenkouA) {
		return InCloud, nil
	}

	// Below the Cloud
	if lastCandle.Close.LessThan(candleCloud.SenkouB) && lastCandle.Close.LessThan(candleCloud.SenkouA) {
		return BelowCloud, nil
	}

	// This should never happen
	return 0, errors.New("No conditions met")
}

// Returns a decimal value representing the slope of the Tenkan
func getSlopeOfTenkan(chart *charts.CloudChart) (decimal.Decimal, error) {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1

	if chartLength > 2 {
		return utils.StringToDecimal("0"), errors.New("not enough candles")
	}

	currentCandle := candles[chartLength]
	previousCandle := candles[chartLength-1]
	return currentCandle.Tenkan.Sub(previousCandle.Tenkan), nil
}

// Measures the slope of the Tenkan line. Returns favorable if greater than 0
func hasStrongTenkanSlope(chart *charts.CloudChart) bool {
	slopeOfTenkan, err := getSlopeOfTenkan(chart)
	if err != nil {
		return false
	}

	if slopeOfTenkan.GreaterThan(utils.ZeroDecimal()) {
		return true
	}

	return false

	// Slope Calculation
	// -- Calculate next cross,
	// -- -- if its -10 < x Low Slope
	// -- -- if its -2 < x < 0 High Slope?
}

/*

func FindNextTKCross(chart *CloudChart) decimal.Decimal {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1
	lastCandle := candles[chartLength]
	secondToLast := candles[chartLength-1]

	return findIntersection(lastCandle, secondToLast)
}

func findIntersection(currentCandle *Candle, previousCandle *Candle) decimal.Decimal {
	zero, _ := decimal.NewFromString("0")
	tenkanSlope := currentCandle.Tenkan.Sub(previousCandle.Tenkan)
	kijunSlope := currentCandle.Kijun.Sub(previousCandle.Kijun)

	if tenkanSlope.Equals(zero) && kijunSlope.Equals(zero) {
		return zero
	}

	// (TenkanSlope * X) + prev.Tenkan = (KijunSlope * X) + prev.Kijun
	// ((TenkanSlope - KijunSlope) * X) = (prev.Kijun - prev.Tenkan)
	// X = (prev.Kijun - prev.Tenkan) / (TenkanSlope - KijunSlope)
	num := previousCandle.Kijun.Sub(previousCandle.Tenkan)
	denom := tenkanSlope.Sub(kijunSlope)

	if denom.Equals(zero) {
		return zero
	}

	return num.Div(denom)
}

func getCloudColor(chart *charts.CloudChart) string, bool {
	candle := chart.GetLastCandle()
	senkouA := candle.senkouA
	senkouB := candle.senkouB

	if senkouB.Equals(utils.ZeroDecimal()) {
		return "", false
	}

	if senkouA.GreaterThan(senkouB) {
		return "green"
	}

	if senkouB.GreaterThan(senkouA) {
		return "red"
	}

	if senkouA.Equals(senkouB) {
		return "NONE"
	}

	return "N/A"
}

*/
