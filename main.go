package main

import (
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"os"
)

const (
	TenkanPeriod       = 20
	KijunPeriod        = 60
	SenkouBPeriod      = 120
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

	bittrexClient := exchanges.NewClient(bittrex)

	chart, err := bittrexClient.GetCandles("BTC-ETH", "day")
	if err != nil {
		panic(err)
	}
	chart.Print()

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
