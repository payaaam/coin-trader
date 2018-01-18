package strategies

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/utils"
	"github.com/stretchr/testify/assert"
	//"github.com/shopspring/decimal"
	"testing"
)

var timestampOne = int64(1)
var timestampTwo = int64(2)

// Buy
func TestShouldBuyStrongBullishTkCross(t *testing.T) {
	strat := NewCloudStrategy()

	chartToTest := getBullishTestChart()

	chartToTest.Cloud[timestampTwo] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("8"),
		SenkouB: utils.StringToDecimal("7"),
		Color:   "green",
	}

	assert.True(t, strat.ShouldBuy(chartToTest), "should buy with this chart")
}

func TestShouldBuyNormalBullishTkCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getBullishTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.True(t, strat.ShouldBuy(chartToTest), "should buy with this chart")
}

func TestShouldBuyWeakBullishTkCross(t *testing.T) {
	strat := NewCloudStrategy()

	chartToTest := getBullishTestChart()

	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("5"),
		SenkouB: utils.StringToDecimal("4"),
		Color:   "green",
	}

	assert.True(t, strat.ShouldBuy(chartToTest), "should buy with this chart")
}

func TestShouldNotBuyOnBearishCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getBearishTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.False(t, strat.ShouldBuy(chartToTest), "should NOT buy with this chart")
}

func TestShouldNotBuyOnNoCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getNoCrossTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.False(t, strat.ShouldBuy(chartToTest), "should NOT buy with this chart")
}

// Sell
func TestShouldSellOnBearishTkCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getBearishTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.True(t, strat.ShouldSell(chartToTest), "should sell with this chart")
}

func TestNotShouldSellOnNoTkCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getNoCrossTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.False(t, strat.ShouldSell(chartToTest), "should sell with this chart")
}

func TestNotShouldSellOnBullishTkCross(t *testing.T) {
	strat := NewCloudStrategy()
	chartToTest := getBullishTestChart()
	chartToTest.Cloud[2] = &charts.CloudPoint{
		SenkouA: utils.StringToDecimal("7"),
		SenkouB: utils.StringToDecimal("5"),
		Color:   "green",
	}

	assert.False(t, strat.ShouldSell(chartToTest), "should sell with this chart")
}

func getBullishTestChart() *charts.CloudChart {
	var candles []*charts.Candle
	// Second to last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampOne,
		High:      utils.StringToDecimal("5"),
		Low:       utils.StringToDecimal("5"),
		Open:      utils.StringToDecimal("5"),
		Close:     utils.StringToDecimal("5"),
		Kijun:     utils.StringToDecimal("1"),
		Tenkan:    utils.StringToDecimal("0"),
	})
	// Last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampTwo,
		High:      utils.StringToDecimal("6"),
		Low:       utils.StringToDecimal("6"),
		Open:      utils.StringToDecimal("6"),
		Close:     utils.StringToDecimal("6"),
		Kijun:     utils.StringToDecimal("0"),
		Tenkan:    utils.StringToDecimal("1"),
	})

	chartToTest := &charts.CloudChart{
		Test:  true,
		Cloud: make(map[int64]*charts.CloudPoint),
	}
	chartToTest.SetCandles(candles)

	return chartToTest
}

func getBearishTestChart() *charts.CloudChart {
	var candles []*charts.Candle
	// Second to last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampOne,
		High:      utils.StringToDecimal("5"),
		Low:       utils.StringToDecimal("5"),
		Open:      utils.StringToDecimal("5"),
		Close:     utils.StringToDecimal("5"),
		Kijun:     utils.StringToDecimal("0"),
		Tenkan:    utils.StringToDecimal("1"),
	})
	// Last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampTwo,
		High:      utils.StringToDecimal("6"),
		Low:       utils.StringToDecimal("6"),
		Open:      utils.StringToDecimal("6"),
		Close:     utils.StringToDecimal("6"),
		Kijun:     utils.StringToDecimal("1"),
		Tenkan:    utils.StringToDecimal("0"),
	})

	chartToTest := &charts.CloudChart{
		Test:  true,
		Cloud: make(map[int64]*charts.CloudPoint),
	}
	chartToTest.SetCandles(candles)

	return chartToTest
}

func getNoCrossTestChart() *charts.CloudChart {
	var candles []*charts.Candle
	// Second to last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampOne,
		High:      utils.StringToDecimal("5"),
		Low:       utils.StringToDecimal("5"),
		Open:      utils.StringToDecimal("5"),
		Close:     utils.StringToDecimal("5"),
		Kijun:     utils.StringToDecimal("1"),
		Tenkan:    utils.StringToDecimal("0"),
	})
	// Last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampTwo,
		High:      utils.StringToDecimal("6"),
		Low:       utils.StringToDecimal("6"),
		Open:      utils.StringToDecimal("6"),
		Close:     utils.StringToDecimal("6"),
		Kijun:     utils.StringToDecimal("1"),
		Tenkan:    utils.StringToDecimal("0"),
	})

	chartToTest := &charts.CloudChart{
		Test:  true,
		Cloud: make(map[int64]*charts.CloudPoint),
	}
	chartToTest.SetCandles(candles)

	return chartToTest
}
