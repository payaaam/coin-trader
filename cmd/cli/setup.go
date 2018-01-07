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
	"strings"
)

const PostgresConn = "dbname=coins user=postgres host=localhost port=5432 sslmode=disable"

func SetupExchange(exchange string, interval string) {
	ctx := context.Background()
	log.Infof("Downloading information from %s", exchange)

	// Setup Database Libs
	postgres, err := sql.Open("postgres", PostgresConn)
	if err != nil {
		panic(err)
	}
	marketStore := db.NewMarketStore(postgres)
	chartStore := db.NewChartStore(postgres)
	tickStore := db.NewTickStore(postgres)

	// Setup Exchange Client
	if exchange != "bittrex" {
		panic("Not a valid exchange")
	}
	BittrexAPIKey := os.Getenv("BITTREX_API_KEY")
	BittrexAPISecret := os.Getenv("BITTREX_API_SECRET")
	bittrex := bittrex.New(BittrexAPIKey, BittrexAPISecret)
	bittrexClient := exchanges.NewClient(bittrex)

	err = loadMarkets(ctx, exchange, bittrexClient, marketStore)
	if err != nil {
		log.Error(err)

	}

	err = loadTradingPairsByInterval(ctx, exchange, bittrexClient, marketStore, chartStore, tickStore)
	if err != nil {
		log.Error(err)
	}
}

func normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func loadMarkets(ctx context.Context, exchange string, client *exchanges.BittrexClient, marketStore *db.MarketStore) error {

	markets, err := client.GetMarkets()
	if err != nil {
		return err
	}

	for _, m := range markets {
		market := &models.Market{
			ExchangeName:       exchange,
			BaseCurrency:       normalize(m.BaseCurrency),
			BaseCurrencyName:   null.StringFrom(normalize(m.BaseCurrencyName)),
			MarketCurrency:     normalize(m.MarketCurrency),
			MarketCurrencyName: null.StringFrom(normalize(m.MarketCurrencyName)),
			MarketKey:          normalize(m.MarketKey),
		}
		err = marketStore.Upsert(ctx, market)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadTradingPairsByInterval(ctx context.Context, exchange string, client *exchanges.BittrexClient, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore) error {
	markets, err := marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		log.Info(m.MarketKey)
		// Get or Create Chart
		// Get Tickers for all Markets
	}

	return nil
}
