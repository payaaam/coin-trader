package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/strategies"
	//"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	//"gopkg.in/volatiletech/null.v6"
	"time"
)

const TickerSeconds = 10

type TraderCommand struct {
	config         *Config
	marketStore    *db.MarketStore
	chartStore     *db.ChartStore
	tickStore      *db.TickStore
	exchangeClient exchanges.Exchange
}

func NewTraderCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore, client exchanges.Exchange) *TraderCommand {
	return &TraderCommand{
		config:         config,
		marketStore:    marketStore,
		chartStore:     chartStore,
		tickStore:      tickStore,
		exchangeClient: client,
	}
}

func (t *TraderCommand) Run(exchange string, interval string) {
	ctx := context.Background()
	log.Infof("Starting Automated Trader %s", exchange)

	clientInterval := exchanges.Intervals["bittrex"][interval]

	// Get Open Positions
	state := NewTraderState()
	state.SetBalance("btc-neo", utils.StringToDecimal("20.02"))

	candles, err := t.tickStore.GetChartCandles(ctx, "btc-neo", exchange, interval)
	if err != nil {
		panic(err)
	}

	chart, err := charts.NewCloudChartFromCandles(candles, "btc-neo", "bittrex", interval)
	if err != nil {
		panic(err)
	}

	state.SetChart("btc-neo", chart)

	ichimokuCloudStrategy := strategies.NewCloudStrategy()

	// Will ping
	ticker := time.NewTicker(time.Second * TickerSeconds)
	go func() {
		for _ = range ticker.C {

			candle, err := t.exchangeClient.GetLatestCandle("btc-neo", clientInterval)
			if err != nil {
				panic(err)
			}

			neoChart := state.GetChart("btc-neo")
			neoChart.AddCandle(candle)
			err = t.tickStore.Upsert(ctx, 1, candle)
			if err != nil {
				panic(err)
			}

			balance := state.GetBalance("btc-neo")

			log.Infof("Balance: %v", balance)

			//neoChart.PrintSummary()

			if balance != utils.ZeroDecimal() {

				if ichimokuCloudStrategy.ShouldSell(neoChart) == true {
					/*
						order := executeSell("neo", balance)
						addOrderToActiveOrders()
						updateBalance()
						updateDatabase()
						log.Infof("Selling NEO")
						log.Infof("Amount: %v", amount)
						log.Infof("Price: %v", price)
					*/
				}
			}

			if balance == utils.ZeroDecimal() {
				if ichimokuCloudStrategy.ShouldBuy(chart) == true {
					/*
						// check for TK cross
						//  - Cloud Color
						//  - Price in cloud

							order := executeBuy("neo", amount)
							addOrderToActiveOrders()
							updateBalance()
							updateDatabase()
							log.Infof("Purchased NEO")
							log.Infof("Amount: %v", amount)
							log.Infof("Price: %v", price)

					*/
				}
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
