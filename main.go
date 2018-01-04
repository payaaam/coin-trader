package main

import (
	"github.com/payaaam/coin-trader/exchanges"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"os"
	"time"
)

const JSONPath = "/Users/payam/Development/go/src/github.com/payaaam/coin-trader/testdata/BTC-ETH.json"

const Interval = "day"

func main() {

	BittrexAPIKey := os.Getenv("BITTREX_API_KEY")
	BittrexAPISecret := os.Getenv("BITTREX_API_SECRET")
	bittrex := bittrex.New(BittrexAPIKey, BittrexAPISecret)

	bittrexClient := exchanges.NewClient(bittrex)

	markets, err := bittrexClient.GetMarkets()
	if err != nil {
		panic(err)
	}

	for _, markets := range markets {
		chart, err := bittrexClient.GetCandles(markets.TradingPair, Interval)
		if err != nil {
			log.Error(err)
		}

		chart.PrintSummary()

		time.Sleep(1000 * time.Millisecond)
	}
}
