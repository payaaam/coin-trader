package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	"gopkg.in/volatiletech/null.v6"
)

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
	log.Infof("Setting up %s", exchange)

	if s.config.Bittrex == nil {
		panic("No Bittrex Config Found")
	}

	bittrex := bittrex.New(s.config.Bittrex.ApiKey, s.config.Bittrex.ApiSecret)
	bittrexClient := exchanges.NewClient(bittrex)

	err := s.loadMarkets(ctx, exchange, bittrexClient)
	if err != nil {
		log.Error(err)
	}

	err = s.loadTradingPairsByInterval(ctx, exchange, interval, bittrexClient)
	if err != nil {
		log.Error(err)
	}
}

func (s *SetupCommand) loadMarkets(ctx context.Context, exchange string, client *exchanges.BittrexClient) error {
	log.Debug("Loading Markets")
	markets, err := client.GetMarkets()
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

func (s *SetupCommand) loadTradingPairsByInterval(ctx context.Context, exchange string, interval string, client *exchanges.BittrexClient) error {
	markets, err := s.marketStore.GetMarkets(ctx, exchange)
	if err != nil {
		return err
	}

	for _, m := range markets {

		if m.MarketKey != "btc-neo" {
			continue
		}

		log.Infof("Loading Ticks: %s", m.MarketKey)
		chart := &models.Chart{
			MarketID: m.ID,
			Interval: interval,
		}

		err = s.chartStore.Upsert(ctx, chart)
		if err != nil {
			return err
		}

		err = s.loadTicks(ctx, chart.ID, m.MarketKey, interval, client)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *SetupCommand) loadTicks(ctx context.Context, chartID int, marketKey string, interval string, client *exchanges.BittrexClient) error {
	clientInterval := exchanges.Intervals["bittrex"][interval]
	candles, err := client.GetCandles(marketKey, clientInterval)
	if err != nil {
		return err
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

		err := s.tickStore.Upsert(ctx, tick)
		if err != nil {
			return err
		}
	}
	return nil
}
