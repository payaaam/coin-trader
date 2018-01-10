package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/shopspring/decimal"
	//"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	//"gopkg.in/volatiletech/null.v6"
	"time"
)

const TickerSeconds = 10

type TraderCommand struct {
	config      *Config
	marketStore *db.MarketStore
	chartStore  *db.ChartStore
	tickStore   *db.TickStore
}

func NewTraderCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore) *TraderCommand {
	return &TraderCommand{
		config:      config,
		marketStore: marketStore,
		chartStore:  chartStore,
		tickStore:   tickStore,
	}
}

func (t *TraderCommand) Run(exchange string, interval string) {
	ctx := context.Background()
	log.Infof("Starting Automated Trader %s", exchange)

	if t.config.Bittrex == nil {
		panic("No Bittrex Config Found")
	}

	bittrex := bittrex.New(t.config.Bittrex.ApiKey, t.config.Bittrex.ApiSecret)
	bittrexClient := exchanges.NewClient(bittrex)
	clientInterval := exchanges.Intervals["bittrex"][interval]

	// Get Open Positions
	state := NewTraderState()
	state.SetBalance("btc-neo", utils.StringToDecimal("20.02"))

	candles, err := t.tickStore.GetChartCandles(ctx, "btc-neo", exchange, interval)
	if err != nil {
		panic(err)
	}

	chart, err := charts.NewCloudChart(candles, "btc-neo", "bittrex")
	if err != nil {
		panic(err)
	}

	state.SetChart("btc-neo", chart)

	// Will ping
	ticker := time.NewTicker(time.Second * TickerSeconds)
	go func() {
		for _ = range ticker.C {

			candle, err := bittrexClient.GetLatestCandle("btc-neo", clientInterval)
			if err != nil {
				panic(err)
			}

			neoChart := state.GetChart("btc-neo")
			neoChart.AddCandle(candle)
			err = t.tickStore.Upsert(ctx, 126, candle)
			if err != nil {
				panic(err)
			}

			balance := state.GetBalance("btc-neo")

			log.Infof("Balance: %v", balance)

			neoChart.PrintSummary()

			if balance != decimal.New(0, 0) {
				/*
					chart := state.GetChart("btc-neo")
					if shouldSell(chart) == true {
						order := executeSell("neo", balance)
						addOrderToActiveOrders()
						updateBalance()
						updateDatabase()
						log.Infof("Selling NEO")
						log.Infof("Amount: %v", amount)
						log.Infof("Price: %v", price)
					}
				*/
				//return
			}

			if balance == decimal.New(0, 0) {
				/*
					chart := state.GetChart("btc-neo")
					// check for TK cross
					//  - Cloud Color
					//  - Price in cloud
					if shouldBuy(chart) == true {
						order := executeBuy("neo", amount)
						addOrderToActiveOrders()
						updateBalance()
						updateDatabase()
						log.Infof("Purchased NEO")
						log.Infof("Amount: %v", amount)
						log.Infof("Price: %v", price)
					}
				*/
				//return
			}
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for _ = range signalChan {
		log.Warn("received SIGINT or SIGTERM")
		break
	}
	log.Info("Shutting down")
}
