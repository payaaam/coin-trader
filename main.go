package main

import (
	//"github.com/payaaam/coin-trader/exchanges"
	//"github.com/payaaam/coin-trader/strategies"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"os"
)

const JSONPath = "/Users/payam/Development/go/src/github.com/payaaam/coin-trader/testdata/BTC-ETH.json"

func main() {

	BittrexAPIKey := os.Getenv("BITTREX_API_KEY")
	BittrexAPISecret := os.Getenv("BITTREX_API_SECRET")
	bittrex := bittrex.New(BittrexAPIKey, BittrexAPISecret)
	/*
		bittrexClient := exchanges.NewClient(bittrex)
		chart, err := bittrexClient.GetCandles("BTC-DCR", "hour")
		if err != nil {
			panic(err)
		}
		strategies.FindNextTKCross(chart)
	*/
	/*
		localBittrexClient := exchanges.NewLocalClient()
		chart, err := localBittrexClient.GetCandles(JSONPath, "BTC-ETH")
		if err != nil {
			panic(err)
		}
	*/

	/*
		crosses := strategies.FindTKCrosses(chart)
		for _, ts := range crosses {
			log.Info(ts)
		}

		strategies.GetTKSlopeLines(chart)
	*/

	printMarketPairs(bittrex)

	// Get markets

	//log.Infof("%d days", len(candles))
}

func printMarketPairs(bittrex *bittrex.Bittrex) {
	markets, err := bittrex.GetMarkets()
	if err != nil {
		panic(err)
	}
	log.Infof("Market Count: %v", len(markets))
	count := 0
	for _, market := range markets {
		if market.BaseCurrency == "BTC" {
			count++
			log.Infof("Market: %s", market.MarketCurrency)
		}

	}
	log.Infof("BTC Market Count: %v", count)

}
