package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	"gopkg.in/volatiletech/null.v6"
	"os"
	"os/signal"
	"time"
)

var MarketKey = "btc-ebst"

var AllTicksTimeout time.Duration = 5 * time.Second
var LatestTickTimeout time.Duration = 1 * time.Second
var OnTheFiftyMinuteMark = 50

type SetupCommand struct {
	config      *Config
	marketStore *db.MarketStore
	chartStore  *db.ChartStore
	tickStore   *db.TickStore
}

func NewSetupCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore) *SetupCommand {
	return &SetupCommand{
		config:      config,
		marketStore: marketStore,
		chartStore:  chartStore,
		tickStore:   tickStore,
	}
}

func (s *SetupCommand) Run(exchange string, interval string) {
	ctx := context.Background()
	log.Infof("Starting Ticker %s", exchange)

	if s.config.Bittrex == nil {
		panic("No Bittrex Config Found")
	}

	bittrex := bittrex.New(s.config.Bittrex.ApiKey, s.config.Bittrex.ApiSecret)
	bittrexClient := exchanges.NewBittrexClient(bittrex)

	err := s.loadMarkets(ctx, exchange, bittrexClient)
	if err != nil {
		log.Error(err)
	}

	// Fetch Daily Once
	go s.fetchOnce(ctx, exchange, db.OneDayInterval, bittrexClient)

	// Fetch Hourly
	ticker := time.NewTicker(time.Minute * 1)
	go s.fetchInterval(ctx, exchange, db.OneHourInterval, bittrexClient, ticker)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for _ = range signalChan {
		log.Warn("received SIGINT or SIGTERM")
		break
	}
	log.Info("Shutting down")
}

func (s *SetupCommand) fetchOnce(ctx context.Context, exchange string, interval string, client exchanges.Exchange) {
	err := s.loadTradingPairsByInterval(ctx, exchange, interval, client)
	if err != nil {
		log.Error(err)
	}
}

func (s *SetupCommand) fetchInterval(ctx context.Context, exchange string, interval string, client exchanges.Exchange, ticker *time.Ticker) {
	for t := range ticker.C {
		if t.Minute() == OnTheFiftyMinuteMark {
			err := s.loadTradingPairsByInterval(ctx, exchange, interval, client)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (s *SetupCommand) loadMarkets(ctx context.Context, exchange string, client exchanges.Exchange) error {
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
		err = s.marketStore.Upsert(ctx, market)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SetupCommand) loadTradingPairsByInterval(ctx context.Context, exchange string, interval string, client exchanges.Exchange) error {
	markets, err := s.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		log.Infof("--- %s ---", m.MarketKey)
		chartID, err := s.getChartID(ctx, m.ID, interval)
		if err != nil {
			return err
		}

		shouldFetchLatest, err := s.shouldFetchAllTicks(ctx, chartID, interval)
		if err != nil {
			return err
		}

		// Checks if we should fetch all ticks, or just the latest
		if shouldFetchLatest == true {
			err = s.loadAllTicks(ctx, chartID, m.MarketKey, interval, client)
			if err != nil {
				return err
			}
			time.Sleep(AllTicksTimeout)
			continue
		}

		err = s.loadLatestTick(ctx, chartID, m.MarketKey, interval, client)
		if err != nil {
			return err
		}
		time.Sleep(LatestTickTimeout)
	}

	return nil
}

// Load all the ticks for a graph. This API is expensive
func (s *SetupCommand) loadAllTicks(ctx context.Context, chartID int, marketKey string, interval string, client exchanges.Exchange) error {
	log.Infof("Fetched All Ticks: %s", interval)
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candles, err := client.GetCandles(marketKey, clientInterval)
	if err != nil {
		return err
	}

	for _, candle := range candles {
		err := s.tickStore.Upsert(ctx, chartID, candle)
		if err != nil {
			return err
		}
	}
	return nil
}

// Loads the latest tick for a marketKey. This API is inexpensive
func (s *SetupCommand) loadLatestTick(ctx context.Context, chartID int, marketKey string, interval string, client exchanges.Exchange) error {
	log.Infof("Fetched Latest Ticks: %s", interval)
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candle, err := client.GetLatestCandle(marketKey, clientInterval)
	if err != nil {
		return err
	}

	err = s.tickStore.Upsert(ctx, chartID, candle)
	if err != nil {
		return err
	}

	return nil
}

func (s *SetupCommand) addDailyCandles(ctx context.Context, exchange string, baseInterval string, newCandleInterval string) error {

	markets, err := s.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		log.Infof("--- %s ---", m.MarketKey)
		chartID, err := s.getChartID(ctx, m.ID, baseInterval)
		if err != nil {
			return err
		}

		log.Info(chartID)

		// Get Last 24 Candles
		// high -> Highest over 24
		// Low -> Lowest over 24
		// Open -> First candle.open
		// Close -> Last candle.close
		// Volume (All candles added togather?)
		// Timestamp = last timstamp
		// Day???

		/*
			chartID, err := s.getChartID(ctx, m.ID, newCandleInterval)
			if err != nil {
				return err
			}
		*/

		// Get 1D chart
		// Upsert Candle to chart.ID
	}
	return nil
}

// Get ChartID from Market
func (s *SetupCommand) getChartID(ctx context.Context, marketID int, interval string) (int, error) {
	chart := &models.Chart{
		MarketID: marketID,
		Interval: interval,
	}

	err := s.chartStore.Upsert(ctx, chart)
	if err != nil {
		return 0, err
	}

	return chart.ID, nil
}

// Determines if the CLI should fetch all ticks, or just latest
func (s *SetupCommand) shouldFetchAllTicks(ctx context.Context, chartID int, interval string) (bool, error) {
	latestCandle, err := s.tickStore.GetLatestChartCandle(ctx, chartID)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		latestCandle = &charts.Candle{}
	}

	// Determine if we need to fetch all candles, or just the latest
	lastTimestamp := getLastTimestamp(interval)
	intervalMS := intervalMilliseconds(interval)

	if lastTimestamp-latestCandle.TimeStamp > intervalMS {
		return true, nil
	}
	return false, nil
}

// Gets the last timestamp for the chart given an interval.
func getLastTimestamp(interval string) int64 {
	ts := time.Now().UTC()

	if interval == db.OneDayInterval {
		rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)
		return rounded.Unix()
	}

	if interval == db.OneHourInterval {
		rounded := time.Date(ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 0, 0, 0, time.UTC)
		return rounded.Unix()
	}

	return 0
}

// Converts a interval to milliseconds. 1h = 60 minutes * 60 seconds
func intervalMilliseconds(interval string) int64 {
	return int64(db.IntervalMilliseconds[interval])
}
