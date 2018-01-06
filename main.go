package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	"gopkg.in/volatiletech/null.v6"
	"os"
	//	"time"
	"strings"
)

const PostgresConn = "dbname=coins user=postgres host=localhost port=5432 sslmode=disable"
const JSONPath = "/Users/payam/Development/go/src/github.com/payaaam/coin-trader/testdata/BTC-ETH.json"

const Interval = "day"

func normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func main() {
	ctx := context.Background()
	postgres, err := sql.Open("postgres", PostgresConn)
	if err != nil {
		panic(err)
	}

	// Get or Create Market
	marketStore := db.NewMarketStore(postgres)
	market := &models.Market{
		ExchangeName:       normalize("Bittrex"),
		BaseCurrency:       normalize("BTC"),
		BaseCurrencyName:   null.StringFrom(normalize("Bitcoin")),
		MarketCurrency:     normalize("ETH"),
		MarketCurrencyName: null.StringFrom(normalize("Ethereum")),
		MarketKey:          normalize("BTC-ETH"),
	}
	err = marketStore.Upsert(ctx, market)
	if err != nil {
		panic(err)
	}

	// Create Chart
	chartStore := db.NewChartStore(postgres)
	chart := &models.Chart{
		MarketID: market.ID,
		Interval: exchanges.Intervals[market.ExchangeName][Interval],
	}

	err = chartStore.Upsert(ctx, chart)
	if err != nil {
		panic(err)
	}

	// Ticks
	tickStore := db.NewTickStore(postgres)
	BittrexAPIKey := os.Getenv("BITTREX_API_KEY")
	BittrexAPISecret := os.Getenv("BITTREX_API_SECRET")
	bittrex := bittrex.New(BittrexAPIKey, BittrexAPISecret)

	log.Info("Getting Candles")

	bittrexClient := exchanges.NewClient(bittrex)
	candles, err := bittrexClient.GetCandles(market.MarketKey, Interval)
	if err != nil {
		panic(err)
	}

	log.Infof("Received %d candles", len(candles))

	for _, candle := range candles {
		tick := &models.Tick{
			ChartID:   chart.ID,
			Open:      candle.Open.String(),
			Close:     candle.Close.String(),
			High:      candle.High.String(),
			Low:       candle.Low.String(),
			Volume:    candle.Volume.String(),
			Timestamp: candle.TimeStamp.Unix(),
			Day:       candle.Day,
		}

		err := tickStore.Upsert(ctx, tick)
		if err != nil {
			panic(err)
		}
	}

	candle, err := bittrexClient.GetLatestCandle(market.MarketKey, Interval)
	if err != nil {
		panic(err)
	}

	log.Infof("Latest Candle: %v", candle)

	tick := &models.Tick{
		ChartID:   chart.ID,
		Open:      candle.Open.String(),
		Close:     candle.Close.String(),
		High:      candle.High.String(),
		Low:       candle.Low.String(),
		Volume:    candle.Volume.String(),
		Timestamp: candle.TimeStamp.Unix(),
		Day:       candle.Day,
	}

	err = tickStore.Upsert(ctx, tick)
	if err != nil {
		panic(err)
	}
}

/*
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
*/
