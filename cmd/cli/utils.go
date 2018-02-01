package main

import (
	"database/sql"
	"errors"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	"gopkg.in/volatiletech/null.v6"
	"time"
)

var ErrExchangeNotSetup = "exchange not setup"

func getExchangeClient(config *Config, exchange string) (exchanges.Exchange, error) {

	if exchange == "bittrex" {
		if config.Bittrex == nil {
			return nil, errors.New(ErrExchangeNotSetup)
		}
		bittrex := bittrex.New(config.Bittrex.ApiKey, config.Bittrex.ApiSecret)
		return exchanges.NewBittrexClient(bittrex), nil
	}

	/*
		if exchange == "binance" {
			if config.Binance == nil {
				return nil, errors.New(ErrExchangeNotSetup)
			}
		}
	*/

	return nil, errors.New(ErrExchangeNotSetup)
}

// Determines if the CLI should fetch all ticks, or just latest
func shouldFetchAllTicks(ctx context.Context, tickStore db.TickStoreInterface, chartID int, interval string) (bool, error) {
	latestCandle, err := tickStore.GetLatestChartCandle(ctx, chartID)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return true, nil
	}

	// Determine if we need to fetch all candles, or just the latest
	lastTimestamp := getLastTimestamp(interval)
	intervalMS := intervalMilliseconds(interval)

	if latestCandle.TimeStamp-lastTimestamp > intervalMS {
		return true, nil
	}
	return false, nil
}

func loadChart(ctx context.Context, chartStore db.ChartStoreInterface, marketID int, interval string) (*models.Chart, error) {
	chart := &models.Chart{
		MarketID: marketID,
		Interval: interval,
	}

	err := chartStore.Upsert(ctx, chart)
	if err != nil {
		return nil, err
	}

	return chart, nil
}

func loadMarkets(ctx context.Context, marketStore db.MarketStoreInterface, client exchanges.Exchange, exchange string) error {
	log.Debug("Loading BTC Markets")
	markets, err := client.GetBitcoinMarkets()
	if err != nil {
		return err
	}

	for _, m := range markets {

		market := &models.Market{
			ExchangeName:       exchange,
			BaseCurrency:       utils.Normalize(m.BaseCurrency),
			BaseCurrencyName:   null.StringFrom(utils.Normalize(m.BaseCurrencyName)),
			MarketCurrency:     utils.Normalize(m.MarketCurrency),
			MarketCurrencyName: null.StringFrom(utils.Normalize(m.MarketCurrencyName)),
			MarketKey:          utils.Normalize(m.MarketKey),
		}
		err = marketStore.Upsert(ctx, market)
		if err != nil {
			return err
		}
	}

	return nil
}

// Loads all ticks exchange
func loadTicks(ctx context.Context, tickStore db.TickStoreInterface, client exchanges.Exchange, chartID int, marketKey string, interval string) error {
	logInfo(marketKey, interval, "Fetched All Ticks")
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candles, err := client.GetCandles(marketKey, clientInterval)
	if err != nil {
		return err
	}

	for _, candle := range candles {
		err := tickStore.Upsert(ctx, chartID, candle)
		if err != nil {
			return err
		}
	}
	return nil
}

// Loads the latest tick for a marketKey. This API is inexpensive
func loadLatestTick(ctx context.Context, tickStore db.TickStoreInterface, client exchanges.Exchange, chartID int, marketKey string, interval string) error {
	logInfo(marketKey, interval, "Fetched Latest Ticks")
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candle, err := client.GetLatestCandle(marketKey, clientInterval)
	if err != nil {
		return err
	}

	err = tickStore.Upsert(ctx, chartID, candle)
	if err != nil {
		return err
	}

	return nil
}

// Gets the last timestamp for the chart given an interval.
func getLastTimestamp(interval string) int64 {
	ts := time.Now().UTC()

	if interval == charts.OneDayInterval {
		rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)
		return rounded.Unix()
	}

	if interval == charts.OneHourInterval {
		rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 0, 0, 0, time.UTC)
		return rounded.Unix()
	}

	return 0
}

// Gets the timestamp range for the previous period. This is used to create daily candles from
// hourly candles
func getPreviousPeriodRange(interval string) (int64, int64) {
	ts := time.Now().UTC()

	if interval == db.OneDayInterval {
		start := time.Date(ts.Year(), ts.Month(), ts.Day()-1, 0, 0, 0, 0, time.UTC)
		end := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

		log.Info()

		return start.Unix(), end.Unix()
	}

	if interval == db.OneHourInterval {
		start := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour()-1, 0, 0, 0, time.UTC)
		end := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 0, 0, 0, time.UTC)
		return start.Unix(), end.Unix()
	}

	return 0, 0
}

// Converts a interval to milliseconds. 1h = 60 minutes * 60 seconds
func intervalMilliseconds(interval string) int64 {
	return int64(charts.IntervalMilliseconds[interval])
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
