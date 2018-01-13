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
var HourStartOfNewDay = 0

type TickerCommand struct {
	config         *Config
	marketStore    *db.MarketStore
	chartStore     *db.ChartStore
	tickStore      *db.TickStore
	exchangeClient exchanges.Exchange
}

func NewTickerCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore, client exchanges.Exchange) *TickerCommand {
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
	log.Infof("Starting Ticker %s", exchange)

	err := loadMarkets(ctx, t.marketStore, t.exchangeClient, exchange)
	if err != nil {
		log.Error(err)
	}

	t.addDailyCandles(ctx, exchange, db.OneHourInterval, db.OneDayInterval)

	/*

		// Fetch Daily Once
		go t.fetchOnce(ctx, exchange, db.OneDayInterval, t.exchangeClient)

		// Fetch Hourly
		ticker := time.NewTicker(time.Minute * 1)
		go t.fetchInterval(ctx, exchange, db.OneHourInterval, t.exchangeClient, ticker)
	*/

	// Update Daily Chart

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for _ = range signalChan {
		log.Warn("received SIGINT or SIGTERM")
		break
	}
	log.Info("Shutting down")
}

func (t *TickerCommand) fetchOnce(ctx context.Context, exchange string, interval string, client exchanges.Exchange) {
	err := t.loadTradingPairsByInterval(ctx, exchange, interval, client)
	if err != nil {
		log.Error(err)
	}
}

func (t *TickerCommand) fetchInterval(ctx context.Context, exchange string, interval string, client exchanges.Exchange, ticker *time.Ticker) {
	for ticker := range ticker.C {
		if ticker.Minute() == OnTheFiftyMinuteMark {
			err := t.loadTradingPairsByInterval(ctx, exchange, interval, client)
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
		log.Infof("--- %s ---", m.MarketKey)
		chart, err := loadChart(ctx, t.chartStore, m.ID, interval)
		if err != nil {
			return err
		}

		shouldFetchLatest, err := shouldFetchAllTicks(ctx, t.tickStore, chart.ID, interval)
		if err != nil {
			return err
		}

		// Checks if we should fetch all ticks, or just the latest
		if shouldFetchLatest == true {
			err = loadTicks(ctx, t.tickStore, client, chart.ID, m.MarketKey, interval)
			if err != nil {
				return err
			}
			time.Sleep(AllTicksTimeout)
			continue
		}

		err = loadLatestTick(ctx, t.tickStore, client, chart.ID, m.MarketKey, interval)
		if err != nil {
			return err
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
		if m.MarketKey != "btc-eth" {
			continue
		}

		log.Infof("--- %s ---", m.MarketKey)
		chart, err := loadChart(ctx, t.chartStore, m.ID, baseInterval)
		if err != nil {
			return err
		}

		log.Infof("ChartID: %v", chart.ID)
		startTime, endTime := getPreviousPeriodRange(newCandleInterval)
		log.Infof("start: %v - end: %v", startTime, endTime)
		candles, err := t.tickStore.GetCandlesFromRange(ctx, chart.ID, startTime, endTime)
		if err != nil {
			return err
		}

		log.Infof("# CANDLES: %v", len(candles))

		var high = candles[0].High
		var low = candles[0].Low
		var close = candles[0].Close
		var open = candles[len(candles)-1].Open
		var volume = utils.ZeroDecimal()
		var newDay = candles[len(candles)-1].Day + 1

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
			Day:       newDay,
		}

		dailyCandle.Print()
		/*

			dailyChartID, err := t.getChartID(ctx, m.ID, newCandleInterval)
			if err != nil {
				return err
			}

			err = t.tickStore.Upsert(ctx, dailyChartID, dailyCandle)
			if err != nil {
				return err
			}
		*/

		// Get 1D chart
		// Upsert Candle to chart.ID

	}
	return nil
}
