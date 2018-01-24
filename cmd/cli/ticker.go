package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	//"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	"time"
)

var AllTicksTimeout time.Duration = 5 * time.Second
var LatestTickTimeout time.Duration = 1 * time.Second
var OnTheFiftyMinuteMark = 50
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

	go t.loadDailyTickers(ctx, exchange, t.exchangeClient)

	go t.setupHourlyTickerInterval(ctx, exchange, t.exchangeClient)

	go t.setupDailyCandleTickerInterval(ctx, exchange)

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

func (t *TickerCommand) loadDailyTickers(ctx context.Context, exchange string, client exchanges.Exchange) {
	err := t.loadTradingPairsByInterval(ctx, exchange, charts.OneDayInterval, client)
	if err != nil {
		log.Error(err)
	}
}

func (t *TickerCommand) setupHourlyTickerInterval(ctx context.Context, exchange string, client exchanges.Exchange) {
	ticker := time.NewTicker(time.Minute * 1)
	for ticker := range ticker.C {
		if ticker.Minute() == OnTheFiftyMinuteMark {
			log.WithFields(log.Fields{
				"exchange": exchange,
			}).Info("Fetching Hourly Candles")
			err := t.loadTradingPairsByInterval(ctx, exchange, charts.OneHourInterval, client)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (t *TickerCommand) setupDailyCandleTickerInterval(ctx context.Context, exchange string) {
	ticker := time.NewTicker(time.Minute * 1)
	for ticker := range ticker.C {
		if ticker.Hour() == HourStartOfNewDay && ticker.Minute() == MinuteStartOfNewDay {
			log.WithFields(log.Fields{
				"exchange": exchange,
			}).Info("Processing 24H Candles")
			err := t.addDailyCandles(ctx, exchange, charts.OneHourInterval, charts.OneDayInterval)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (t *TickerCommand) loadTradingPairsByInterval(ctx context.Context, exchange string, interval string, client exchanges.Exchange) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		logInfo(m.MarketKey, interval, "Processing")
		chart, err := loadChart(ctx, t.chartStore, m.ID, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
			continue
		}

		shouldFetchAllTicks, err := shouldFetchAllTicks(ctx, t.tickStore, chart.ID, interval)
		if err != nil {
			logError(m.MarketKey, interval, err)
			continue
		}

		// Checks if we should fetch all ticks, or just the latest
		if shouldFetchAllTicks == true {
			err = loadTicks(ctx, t.tickStore, client, chart.ID, m.MarketKey, interval)
			if err != nil {
				logError(m.MarketKey, interval, err)
			}
			time.Sleep(AllTicksTimeout)
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

func (t *TickerCommand) addDailyCandles(ctx context.Context, exchange string, baseInterval string, newCandleInterval string) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		logInfo(m.MarketKey, newCandleInterval, "Generating 24H Candle")
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
			log.Warn("Candles do not exist in this timeframe.")
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
