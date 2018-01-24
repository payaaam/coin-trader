package main

import (
	"fmt"
	"github.com/fatih/color"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/strategies"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type BackTestCommand struct {
	config         *Config
	marketStore    db.MarketStoreInterface
	chartStore     db.ChartStoreInterface
	tickStore      db.TickStoreInterface
	exchangeClient exchanges.Exchange
}

func NewBackTestCommand(config *Config, marketStore db.MarketStoreInterface, chartStore db.ChartStoreInterface, tickStore db.TickStoreInterface, client exchanges.Exchange) *BackTestCommand {
	return &BackTestCommand{
		config:         config,
		marketStore:    marketStore,
		chartStore:     chartStore,
		tickStore:      tickStore,
		exchangeClient: client,
	}
}

// UPDATE HERE when testing a new strategy
func GetStrategy() strategies.Strategy {
	return strategies.NewCloudStrategy()
}

func (b *BackTestCommand) Run(exchange string, interval string, marketKey string) {
	ctx := context.Background()

	if b.config.Bittrex == nil {
		logFatal(fmt.Errorf("bittrex config not found"))
	}

	err := loadMarkets(ctx, b.marketStore, b.exchangeClient, exchange)
	if err != nil {
		logError(marketKey, interval, err)
		return
	}

	logInfo(marketKey, interval, "Running backtest")

	market, err := b.marketStore.GetMarket(ctx, exchange, marketKey)
	if err != nil {
		logError(marketKey, interval, err)
		return
	}

	chart, err := loadChart(ctx, b.chartStore, market.ID, interval)
	if err != nil {
		logError(marketKey, interval, err)
		return
	}

	err = loadTicks(ctx, b.tickStore, b.exchangeClient, chart.ID, market.MarketKey, interval)
	if err != nil {
		logError(marketKey, interval, err)
		return
	}

	candles, err := b.tickStore.GetAllChartCandles(ctx, marketKey, exchange, interval)
	if err != nil {
		logError(marketKey, interval, err)
		return
	}

	b.Test(candles, exchange, marketKey, interval)
}

func (b *BackTestCommand) Test(candles []*charts.Candle, exchange string, marketKey string, interval string) {
	ichimokuCloudStrategy := GetStrategy()
	chart := charts.NewCloudChart(marketKey, exchange, interval)

	originalPrice := candles[0].Open

	balance := utils.StringToDecimal("1")
	var inPlay decimal.Decimal
	var hasMoneyInPlay = false
	for _, candle := range candles {
		chart.AddCandle(candle)
		if ichimokuCloudStrategy.ShouldBuy(chart) == true && hasMoneyInPlay == false {
			balance, inPlay = buy(balance, candle)
			hasMoneyInPlay = true

		}

		if ichimokuCloudStrategy.ShouldSell(chart) == true && hasMoneyInPlay == true {
			balance, inPlay = sell(balance, candle)
			hasMoneyInPlay = false
		}
	}

	calculateWinnings(originalPrice, balance.Add(inPlay))
}

func sell(balance decimal.Decimal, candle *charts.Candle) (decimal.Decimal, decimal.Decimal) {
	red := color.New(color.FgRed).SprintFunc()
	balance = balance.Add(candle.Close)
	log.Infof("SELL: %v - TimeStamp: %v", red(candle.Close), time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
	return balance, utils.ZeroDecimal()
}

func buy(balance decimal.Decimal, candle *charts.Candle) (decimal.Decimal, decimal.Decimal) {
	green := color.New(color.FgGreen).SprintFunc()
	balance = balance.Sub(candle.Close)
	log.Printf("BUY:  %v - TimeStamp: %v", green(candle.Close), time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
	return balance, candle.Close
}

func calculateWinnings(original decimal.Decimal, endAmount decimal.Decimal) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	winnings := endAmount.Sub(utils.StringToDecimal("1")).Sub(original)
	percentChange := winnings.Div(original).Mul(utils.StringToDecimal("100")).Round(2)

	if percentChange.Sign() == 1 {
		log.Infof("Summary - Percent Change: %v", green(percentChange))
	} else {
		log.Infof("Summary - Percent Change: %v", red(percentChange))
	}
}
