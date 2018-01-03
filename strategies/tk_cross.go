package strategies

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

func FindTKCrosses(chart *charts.CloudChart) {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1
	for day, candle := range candles {
		if day+1 == chartLength {
			break
		}

		nextCandle := candles[day+1]
		currentDiff := candle.Tenkan.Sub(candle.Kijun)
		nextDiff := nextCandle.Tenkan.Sub(nextCandle.Kijun)

		if currentDiff.Sign() != nextDiff.Sign() {
			log.Infof("Timestamp: %v", candle.TimeStamp)
			previousCandle := candles[day-1]
			tSlope := candle.Tenkan.Sub(previousCandle.Tenkan)
			kSlope := candle.Kijun.Sub(previousCandle.Kijun)
			log.Infof("TenkanSlope: %v", tSlope)
			log.Infof("KijunSlope: %v", kSlope)
		}
	}
}

func FindNextTKCross(chart *charts.CloudChart) {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1
	lastCandle := candles[chartLength]
	secondToLast := candles[chartLength-1]

	numOfPeriods := findIntersection(lastCandle, secondToLast)
	log.Infof("Cross in: %v", numOfPeriods)
}

func findIntersection(currentCandle *charts.Candle, previousCandle *charts.Candle) decimal.Decimal {
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

	return num.Div(denom)
}
