package main

import (
	"github.com/payaaam/coin-trader/exchanges"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"os"
)

func main() {
	BittrexApiKey := os.Getenv("BITTREX_API_KEY")
	BittrexApiSecret := os.Getenv("BITTREX_API_SECRET")

	// Bittrex client
	bittrex := bittrex.New(BittrexApiKey, BittrexApiSecret)

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
