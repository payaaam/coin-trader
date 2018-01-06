package main

import (
	"database/sql"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"golang.org/x/net/context"
	"gopkg.in/volatiletech/null.v6"
	//log "github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
	//"github.com/toorop/go-bittrex"
	//"os"
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

	marketStore := db.NewMarketStore(postgres)
	tickStore := db.NewTickStore(postgres)

	var market *models.Market
	market, err = marketStore.GetMarket(ctx, "BTC-ETH", normalize("Bittrex"))
	if err == sql.ErrNoRows {
		market = &models.Market{
			ExchangeName:       normalize("Bittrex"),
			BaseCurrency:       "BTC",
			BaseCurrencyName:   null.StringFrom(normalize("Bitcoin")),
			MarketCurrency:     "ETH",
			MarketCurrencyName: null.StringFrom(normalize("Ethereum")),
			MarketKey:          "BTC-ETH",
		}

		err = marketStore.Save(ctx, market)
		if err != nil {
			panic(err)
		}
	}

	// Create Chart
	var chartID int = 1

	bittrexClient := exchanges.NewBittrexLocalClient()
	candles, err := bittrexClient.GetCandles(JSONPath, "BTC-ETH")
	if err != nil {
		panic(err)
	}

	for _, candle := range candles {
		tick := &models.Tick{
			ChartID:   chartID,
			Open:      candle.Open.String(),
			Close:     candle.Close.String(),
			High:      candle.High.String(),
			Low:       candle.Low.String(),
			Volume:    candle.Volume.String(),
			Timestamp: candle.TimeStamp.Unix(),
			Day:       candle.Day,
		}

		err := tickStore.Save(ctx, tick)
		if err != nil {
			panic(err)
		}
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
