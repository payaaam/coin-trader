package charts

import (
	"github.com/shopspring/decimal"
	//log "github.com/sirupsen/logrus"
	"time"
)

func FindLastTKCross(chart *CloudChart) time.Time {
	candles := chart.GetCandles()
	chartLength := len(candles) - 1
	var lastTime time.Time
	for day, candle := range candles {
		if day+1 == chartLength {
			break
		}

		nextCandle := candles[day+1]
		currentDiff := candle.Tenkan.Sub(candle.Kijun)
		nextDiff := nextCandle.Tenkan.Sub(nextCandle.Kijun)

		if currentDiff.Sign() != nextDiff.Sign() {
			lastTime = candle.TimeStamp
		}
	}
	return lastTime
}

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

	return num.Div(denom)
}
