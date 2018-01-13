package main

import (
	"github.com/adshao/go-binance"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	//"github.com/toorop/go-bittrex"
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

	//bittrex := bittrex.New(s.config.Bittrex.ApiKey, s.config.Bittrex.ApiSecret)
	//bittrexClient := exchanges.NewBittrexClient(bittrex)
	binance := binance.NewClient(s.config.Binance.ApiKey, s.config.Binance.ApiSecret)
	binanceClient := exchanges.NewBinanceClient(binance)
	err := s.loadMarkets(ctx, exchange, binanceClient)
	if err != nil {
		log.Error(err)
	}

	/*
		err = s.loadTradingPairsByInterval(ctx, exchange, interval, bittrexClient)
		if err != nil {
			log.Error(err)
		}
	*/
}

func (s *SetupCommand) loadMarkets(ctx context.Context, exchange string, client exchanges.Exchange) error {
	log.Debug("Loading Markets")
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

func (s *SetupCommand) loadTicks(ctx context.Context, chartID int, marketKey string, interval string, client exchanges.Exchange) error {
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
