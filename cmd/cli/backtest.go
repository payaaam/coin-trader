package main

import (
	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/shopspring/decimal"
	//"github.com/payaaam/coin-trader/db/models"
	//"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/strategies"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	//"github.com/toorop/go-bittrex"
	"golang.org/x/net/context"
	//"gopkg.in/volatiletech/null.v6"
	"time"
)

type BackTestCommand struct {
	config      *Config
	marketStore *db.MarketStore
	chartStore  *db.ChartStore
	tickStore   *db.TickStore
}

func NewBackTestCommand(config *Config, marketStore *db.MarketStore, chartStore *db.ChartStore, tickStore *db.TickStore) *BackTestCommand {
	return &BackTestCommand{
		config:      config,
		marketStore: marketStore,
		chartStore:  chartStore,
		tickStore:   tickStore,
	}
}

func (s *BackTestCommand) Run(exchange string, interval string) {
	ctx := context.Background()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	// Get Market and Chart
	// Get Ticks
	// For Each tick, addCandle to the chart
	// At Each, test should by and should sell

	candles, err := s.tickStore.GetAllChartCandles(ctx, MarketKey, exchange, interval)
	if err != nil {
		log.Error(err)
	}

	ichimokuCloudStrategy := strategies.NewCloudStrategy()
	chart := charts.NewCloudChart(MarketKey, exchange)

	balance := utils.StringToDecimal("1")
	var inPlay decimal.Decimal
	for _, candle := range candles {
		//log.Infof("INSIDE CANDLE: %d", candle.Day)
		chart.AddCandle(candle)
		if ichimokuCloudStrategy.ShouldBuy(chart) == true {
			balance = balance.Sub(candle.Close)
			inPlay = candle.Close
			log.Infof("BUY: %v", green(candle.Close))
			log.Printf("TimeStamp: %v", time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
			log.Info()
		}

		if ichimokuCloudStrategy.ShouldSell(chart) == true {
			log.Infof("SELL: %v", red(candle.Close))
			balance = balance.Add(candle.Close)
			inPlay = utils.ZeroDecimal()
			log.Printf("TimeStamp: %v", time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
			/*
				log.Println("--- Candle ----")
				log.Printf("TimeStamp: %v", time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
				log.Printf("Open: %v", candle.Open)
				log.Printf("Close: %v", candle.Close)
				log.Printf("Tenkan: %v", candle.Tenkan)
				log.Printf("Kijun: %v", candle.Kijun)
			*/
			log.Info()
		}
	}

	log.Infof("NET: %v", balance.Add(inPlay))

	//chart.Print()
}
