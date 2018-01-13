package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	//"github.com/shopspring/decimal"
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
var HourStartOfNewDay = 0

type TickerCommand struct {
	config      *Config
	marketStore *db.MarketStore
	chartStore  *db.ChartStore
	tickStore   *db.TickStore
}

func NewTickerCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore) *TickerCommand {
	return &TickerCommand{
		config:      config,
		marketStore: marketStore,
		chartStore:  chartStore,
		tickStore:   tickStore,
	}
}

func (t *TickerCommand) Run(exchange string) {
	ctx := context.Background()
	log.Infof("Starting Ticker %s", exchange)

	if t.config.Bittrex == nil {
		panic("No Bittrex Config Found")
	}

	bittrex := bittrex.New(t.config.Bittrex.ApiKey, t.config.Bittrex.ApiSecret)
	bittrexClient := exchanges.NewBittrexClient(bittrex)

	err := t.loadMarkets(ctx, exchange, bittrexClient)
	if err != nil {
		log.Error(err)
	}

	t.addDailyCandles(ctx, exchange, db.OneHourInterval, db.OneDayInterval)

	/*

		// Fetch Daily Once
		go t.fetchOnce(ctx, exchange, db.OneDayInterval, bittrexClient)

		// Fetch Hourly
		ticker := time.NewTicker(time.Minute * 1)
		go t.fetchInterval(ctx, exchange, db.OneHourInterval, bittrexClient, ticker)
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

func (t *TickerCommand) loadMarkets(ctx context.Context, exchange string, client exchanges.Exchange) error {
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
		err = t.marketStore.Upsert(ctx, market)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *TickerCommand) loadTradingPairsByInterval(ctx context.Context, exchange string, interval string, client exchanges.Exchange) error {
	markets, err := t.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {
		log.Infof("--- %s ---", m.MarketKey)
		chartID, err := t.getChartID(ctx, m.ID, interval)
		if err != nil {
			return err
		}

		shouldFetchLatest, err := t.shouldFetchAllTicks(ctx, chartID, interval)
		if err != nil {
			return err
		}

		// Checks if we should fetch all ticks, or just the latest
		if shouldFetchLatest == true {
			err = t.loadAllTicks(ctx, chartID, m.MarketKey, interval, client)
			if err != nil {
				return err
			}
			time.Sleep(AllTicksTimeout)
			continue
		}

		err = t.loadLatestTick(ctx, chartID, m.MarketKey, interval, client)
		if err != nil {
			return err
		}
		time.Sleep(LatestTickTimeout)
	}

	return nil
}

// Load all the ticks for a graph. This API is expensive
func (t *TickerCommand) loadAllTicks(ctx context.Context, chartID int, marketKey string, interval string, client exchanges.Exchange) error {
	log.Infof("Fetched All Ticks: %s", interval)
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candles, err := client.GetCandles(marketKey, clientInterval)
	if err != nil {
		return err
	}

	for _, candle := range candles {
		err := t.tickStore.Upsert(ctx, chartID, candle)
		if err != nil {
			return err
		}
	}
	return nil
}

// Loads the latest tick for a marketKey. This API is inexpensive
func (t *TickerCommand) loadLatestTick(ctx context.Context, chartID int, marketKey string, interval string, client exchanges.Exchange) error {
	log.Infof("Fetched Latest Ticks: %s", interval)
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candle, err := client.GetLatestCandle(marketKey, clientInterval)
	if err != nil {
		return err
	}

	err = t.tickStore.Upsert(ctx, chartID, candle)
	if err != nil {
		return err
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
		chartID, err := t.getChartID(ctx, m.ID, baseInterval)
		if err != nil {
			return err
		}

		log.Infof("ChartID: %v", chartID)
		startTime, endTime := getPreviousTimestampRange(newCandleInterval)
		log.Infof("start: %v - end: %v", startTime, endTime)
		candles, err := t.tickStore.GetCandlesFromRange(ctx, chartID, startTime, endTime)
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

// Get ChartID from Market
func (t *TickerCommand) getChartID(ctx context.Context, marketID int, interval string) (int, error) {
	chart := &models.Chart{
		MarketID: marketID,
		Interval: interval,
	}

	err := t.chartStore.Upsert(ctx, chart)
	if err != nil {
		return 0, err
	}

	return chart.ID, nil
}

// Determines if the CLI should fetch all ticks, or just latest
func (t *TickerCommand) shouldFetchAllTicks(ctx context.Context, chartID int, interval string) (bool, error) {
	latestCandle, err := t.tickStore.GetLatestChartCandle(ctx, chartID)
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

func getPreviousTimestampRange(interval string) (int64, int64) {
	ts := time.Now().UTC()

	if interval == db.OneDayInterval {
		start := time.Date(ts.Year(), ts.Month(), ts.Day()-1, 0, 0, 0, 0, time.UTC)
		end := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.UTC)

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
	return int64(db.IntervalMilliseconds[interval])
}
