package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

var AllTicksTimeout time.Duration = 3 * time.Second
var LatestTickTimeout time.Duration = 1 * time.Second

// All Ticks Fetch Times
var AllTicksFetchTimes = []int{1}

// LatestTickFetchTimes
var LatestTickFetchTimes = []int{20, 30, 40, 50}

// Daily Candle Generation
var HourStartOfNewDay = 1
var MinuteStartOfNewDay = 10

type TickerCommand struct {
	config         *Config
	marketStore    db.MarketStoreInterface
	chartStore     db.ChartStoreInterface
	tickStore      db.TickStoreInterface
	exchangeClient exchanges.Exchange
}

func NewTickerCommand(config *Config, marketStore db.MarketStoreInterface, chartStore db.ChartStoreInterface, tickStore db.TickStoreInterface, client exchanges.Exchange) *TickerCommand {
	return &TickerCommand{
		config:         config,
		marketStore:    marketStore,
		chartStore:     chartStore,
		tickStore:      tickStore,
		exchangeClient: client,
	}
}

// Starts the ticker
func (t *TickerCommand) Run(exchange string) {
	ctx := context.Background()
	log.WithFields(log.Fields{
		"type":     "start",
		"exchange": exchange,
	}).Info()

	err := loadMarkets(ctx, t.marketStore, t.exchangeClient, exchange)
	if err != nil {
		log.Error(err)
	}

	// setup hourly on 1 minute
	// setup latest tick on 20, 30, 40, 50
	// set up daily on 10 minute

	// Fetch latest tick every 10 minutes
	// 20 30 40 50
	go t.setupTenMinuteTicker(ctx, exchange, t.exchangeClient)

	// Fetch Hourly at the start of the hour
	// 01
	go t.setupHourlyTicker(ctx, exchange, t.exchangeClient)

	// Generates Daily Candles
	// 10
	go t.setupDailyTicker(ctx, exchange)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for _ = range signalChan {
		log.Warn("received SIGINT or SIGTERM")
		break
	}
	log.WithFields(log.Fields{
		"type":     "stop",
		"exchange": exchange,
	}).Info()
}

// Fetches Latest Tick every 20, 30, 40, 50 minute
func (t *TickerCommand) setupTenMinuteTicker(ctx context.Context, exchange string, client exchanges.Exchange) {
	ticker := time.NewTicker(time.Minute * 1)
	for ticker := range ticker.C {
		if contains(LatestTickFetchTimes, ticker.Minute()) == true {
			logEvent("process_start", charts.OneHourInterval)
			err := t.fetchLatestTick(ctx, exchange, charts.OneHourInterval, client)
			if err != nil {
				log.Error(err)
			}
			logEvent("process_finish", charts.OneHourInterval)
		}
	}
}

// Fetches Latest Tick every 01 minute
func (t *TickerCommand) setupHourlyTicker(ctx context.Context, exchange string, client exchanges.Exchange) {
	ticker := time.NewTicker(time.Minute * 1)
	for ticker := range ticker.C {
		if contains(AllTicksFetchTimes, ticker.Minute()) == true {
			logEvent("process_start", charts.OneHourInterval)
			err := t.fetchAllTicks(ctx, exchange, charts.OneHourInterval, client)
			if err != nil {
				log.Error(err)
			}
			logEvent("process_finish", charts.OneHourInterval)
		}
	}
}

// Processes hourly candles to generate daily candle
func (t *TickerCommand) setupDailyTicker(ctx context.Context, exchange string) {
	ticker := time.NewTicker(time.Minute * 1)
	for ticker := range ticker.C {
		if ticker.Hour() == HourStartOfNewDay && ticker.Minute() == MinuteStartOfNewDay {
			logEvent("process_start", charts.OneDayInterval)
			err := t.generateDailyCandles(ctx, exchange, charts.OneHourInterval, charts.OneDayInterval)
			if err != nil {
				log.Error(err)
			}
			logEvent("process_finish", charts.OneDayInterval)
		}
	}
}

// Function to Fetch Latest Tick from exchange
func (t *TickerCommand) fetchLatestTick(ctx context.Context, exchange string, interval string, client exchanges.Exchange) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		chart, err := loadChart(ctx, t.chartStore, m.ID, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
			continue
		}

		err = loadLatestTick(ctx, t.tickStore, client, chart.ID, m.MarketKey, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
		}
		time.Sleep(LatestTickTimeout)

	}

	return nil
}

// Helper function to fetch all ticks from exchange
func (t *TickerCommand) fetchAllTicks(ctx context.Context, exchange string, interval string, client exchanges.Exchange) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		chart, err := loadChart(ctx, t.chartStore, m.ID, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
			continue
		}

		err = loadTicks(ctx, t.tickStore, client, chart.ID, m.MarketKey, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
		}
		time.Sleep(AllTicksTimeout)
	}

	return nil
}

// Helper function to process tick time range and generate daily candles
func (t *TickerCommand) generateDailyCandles(ctx context.Context, exchange string, baseInterval string, newCandleInterval string) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		logInfo(m.MarketKey, newCandleInterval, "Generating Daily Candle")
		chart, err := loadChart(ctx, t.chartStore, m.ID, baseInterval)
		if err != nil {
			logError(m.MarketKey, newCandleInterval, err)
			continue
		}

		startTime, endTime := getPreviousPeriodRange(newCandleInterval)
		candles, err := t.tickStore.GetCandlesFromRange(ctx, chart.ID, startTime, endTime)
		if err != nil {
			logError(m.MarketKey, newCandleInterval, err)
			continue
		}

		if len(candles) == 0 {
			log.WithFields(log.Fields{
				"tradingPair": m.MarketKey,
			}).Warn("Candles do not exist in this timeframe.")
			continue
		}

		var high = candles[0].High
		var low = candles[0].Low
		var close = candles[0].Close
		var open = candles[len(candles)-1].Open
		var volume = utils.ZeroDecimal()

		for _, c := range candles {
			volume = volume.Add(c.Volume)
			if c.High.GreaterThan(high) {
				high = c.High
			}
			if c.Low.LessThan(low) {
				low = c.Low
			}
		}

		dailyCandle := &charts.Candle{
			High:      high,
			Low:       low,
			Open:      open,
			Close:     close,
			Volume:    volume,
			TimeStamp: startTime,
		}

		dailyChart, err := loadChart(ctx, t.chartStore, m.ID, newCandleInterval)
		if err != nil {
			logError(m.MarketKey, newCandleInterval, err)
			continue
		}

		err = t.tickStore.Upsert(ctx, dailyChart.ID, dailyCandle)
		if err != nil {
			logError(m.MarketKey, newCandleInterval, err)
			continue
		}
	}
	return nil
}
