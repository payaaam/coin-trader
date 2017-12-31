package main

import (
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"os"
)

const (
	TenkanPeriod       = 20
	KijunPeriod        = 60
	SenkouAPerioud     = 120
	CloudLeadingPeriod = 30
)

type Cloud struct {
	SenkouA decimal.Decimal
	SenkouB decimal.Decimal
}

func main() {
	API_KEY := os.Getenv("API_KEY")
	API_SECRET := os.Getenv("API_SECRET")

	// Bittrex client
	bittrex := bittrex.New(API_KEY, API_SECRET)

	getCandles(bittrex)
	//printMarketPairs(bittrex)

	// Get markets

	//log.Infof("%d days", len(candles))
}

func printMarketPairs(bittrex *bittrex.Bittrex) {
	markets, err := bittrex.GetMarkets()
	if err != nil {
		panic(err)
	}
	for _, market := range markets {
		if market.BaseCurrency == "BTC" {
			log.Infof("Market: %s", market.MarketCurrency)
		}

	}
}

func getCandles(bittrex *bittrex.Bittrex) {
	candles, err := bittrex.GetTicks("BTC-ETH", "day")
	if err != nil {
		panic(err)
	}

	cloudMap := make(map[int]*Cloud)

	for day, candle := range candles {

		log.Infof("Day: %d -----", day)
		log.Infof("Timestamp: %v", candle.TimeStamp)
		log.Infof("Open: %v", candle.Open)
		log.Infof("Close: %v", candle.Close)

		tenkan := getTenkan(candles, day)
		log.Infof("Tenkan: %v", tenkan)

		kijun := getKijun(candles, day)
		log.Infof("Kijun: %v", kijun)

		senkouPlotDay := day + CloudLeadingPeriod
		cloudMap[senkouPlotDay] = &Cloud{
			SenkouA: getSenkouA(tenkan, kijun),
			SenkouB: getSenkouB(candles, day),
		}

		/*
			log.Infof("Day: %v", senkouPlotDay)
			log.Infof("Cloud: %v", cloudMap[senkouPlotDay])
		*/

		if cloudMap[day] != nil {
			//senkouA := cloudMap[day].SenkouA
			log.Infof("SenkouA: %v", cloudMap[day].SenkouA)

			log.Infof("SenkouB: %v", cloudMap[day].SenkouB)
			cloudColor := getCloudColor(cloudMap[day].SenkouA, cloudMap[day].SenkouB)
			log.Infof("CloudColor: %v", cloudColor)
		}
		log.Info()
	}
}

func getTenkan(candles []bittrex.Candle, day int) decimal.Decimal {
	zero, _ := decimal.NewFromString("0")
	if day < TenkanPeriod || len(candles) < TenkanPeriod {
		return zero
	}

	var high decimal.Decimal = candles[day].High
	var low decimal.Decimal = candles[day].Low

	dayStop := day - (TenkanPeriod - 1)
	if dayStop < 0 {
		panic("SOMETHING WENT WRONG")
	}

	for i := day; i >= dayStop; i-- {
		candle := candles[i]
		//log.Info(i)

		if candle.High.GreaterThan(high) {
			high = candle.High
		}

		// Test Low
		if candle.Low.LessThan(low) {
			low = candle.Low
		}
	}

	two, _ := decimal.NewFromString("2")
	// 20 periods
	// Highest High + Lowest Low / 2
	return high.Add(low).Div(two)
}

func getKijun(candles []bittrex.Candle, day int) decimal.Decimal {
	zero, _ := decimal.NewFromString("0")
	if day < KijunPeriod || len(candles) < KijunPeriod {
		return zero
	}

	var high decimal.Decimal = candles[day].High
	var low decimal.Decimal = candles[day].Low

	dayStop := day - (KijunPeriod - 1)
	if dayStop < 0 {
		panic("SOMETHING WENT WRONG")
	}

	for i := day; i >= dayStop; i-- {
		candle := candles[i]
		//log.Info(i)

		if candle.High.GreaterThan(high) {
			high = candle.High
		}

		// Test Low
		if candle.Low.LessThan(low) {
			low = candle.Low
		}
	}

	two, _ := decimal.NewFromString("2")
	// 20 periods
	// Highest High + Lowest Low / 2
	return high.Add(low).Div(two)
}

// (Tenkan-sen + Kijun-sen) / 2
func getSenkouA(tenkan decimal.Decimal, kijun decimal.Decimal) decimal.Decimal {
	two, _ := decimal.NewFromString("2")
	return tenkan.Add(kijun).Div(two)
}

//(120-day high + 120-day low) / 2
func getSenkouB(candles []bittrex.Candle, day int) decimal.Decimal {

	zero, _ := decimal.NewFromString("0")
	if day < SenkouAPerioud || len(candles) < SenkouAPerioud {
		return zero
	}

	var high decimal.Decimal = candles[day].High
	var low decimal.Decimal = candles[day].Low

	dayStop := day - (SenkouAPerioud - 1)
	if dayStop < 0 {
		panic("SOMETHING WENT WRONG")
	}

	for i := day; i >= dayStop; i-- {
		candle := candles[i]
		//log.Info(i)

		if candle.High.GreaterThan(high) {
			high = candle.High
		}

		// Test Low
		if candle.Low.LessThan(low) {
			low = candle.Low
		}
	}

	two, _ := decimal.NewFromString("2")
	// 20 periods
	// Highest High + Lowest Low / 2
	return high.Add(low).Div(two)
}

func getCloudColor(senkouA decimal.Decimal, senkouB decimal.Decimal) string {
	zero, _ := decimal.NewFromString("0")
	if senkouB.Equals(zero) {
		return "N/A"
	}

	if senkouA.GreaterThan(senkouB) {
		return "green"
	}

	if senkouB.GreaterThan(senkouA) {
		return "red"
	}

	if senkouA.Equals(senkouB) {
		return "none"
	}

	return "N/A"
}
