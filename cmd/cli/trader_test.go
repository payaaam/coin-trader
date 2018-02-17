package main

// TODO: Finish out for errors

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/mocks"
	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

type TraderCommandTestConfig struct {
	config       *Config
	ChartStore   *mocks.MockChartStoreInterface
	MarketStore  *mocks.MockMarketStoreInterface
	TickStore    *mocks.MockTickStoreInterface
	Exchange     *mocks.MockExchange
	OrderManager *orders.MockOrderManager
	Strategy     *mocks.MockStrategy
}

func newTraderTestConfig(t *testing.T) *TraderCommandTestConfig {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	return &TraderCommandTestConfig{
		Exchange:     mocks.NewMockExchange(mockCtrl),
		ChartStore:   mocks.NewMockChartStoreInterface(mockCtrl),
		MarketStore:  mocks.NewMockMarketStoreInterface(mockCtrl),
		TickStore:    mocks.NewMockTickStoreInterface(mockCtrl),
		OrderManager: orders.NewMockOrderManager(mockCtrl),
		Strategy:     mocks.NewMockStrategy(mockCtrl),
	}
}

var testInterval string = "1h"
var testExchange string = "bittrex"
var timestampOne = int64(1)
var timestampTwo = int64(2)

func TestTradeBuySuccess(t *testing.T) {
	ctx := context.Background()
	testConfig := newTraderTestConfig(t)
	toTest := NewTraderCommand(nil, testConfig.MarketStore, testConfig.ChartStore, testConfig.TickStore, testConfig.Exchange, testConfig.OrderManager)

	m := &models.Market{
		MarketCurrency: "ltc",
		BaseCurrency:   "btc",
		MarketKey:      "btc-ltc",
	}

	buyChart := getBullishTestChart()
	testConfig.TickStore.EXPECT().GetChartCandles(ctx, m.MarketKey, testExchange, testInterval).Return(buyChart.GetCandles(), nil)
	testConfig.OrderManager.EXPECT().GetBalance(m.MarketCurrency).Return(nil)
	testConfig.OrderManager.EXPECT().GetBalance(m.BaseCurrency).Return(&orders.Balance{
		Available: utils.StringToDecimal("1"),
		Total:     utils.StringToDecimal("1"),
	})
	testConfig.Strategy.EXPECT().ShouldBuy(gomock.Any()).Return(true)
	ticker := &exchanges.Ticker{
		Bid:  utils.StringToDecimal("1"),
		Ask:  utils.StringToDecimal("1"),
		Last: utils.StringToDecimal("1"),
	}
	testConfig.Exchange.EXPECT().GetTicker(m.MarketKey).Return(ticker, nil)

	sellOrder := &orders.LimitOrderMatcher{
		Limit:          utils.StringToDecimal("1.01"),
		Quantity:       utils.StringToDecimal("0.01"),
		MarketCurrency: m.MarketCurrency,
		BaseCurrency:   m.BaseCurrency,
	}
	testConfig.OrderManager.EXPECT().ExecuteLimitBuy(ctx, sellOrder).Return(nil)

	err := toTest.trade(ctx, m, testConfig.Strategy, testExchange, testInterval)
	assert.Nil(t, err, "should not error on decide call")

}
func TestTradeBuyFail(t *testing.T) {
	ctx := context.Background()
	testConfig := newTraderTestConfig(t)
	toTest := NewTraderCommand(nil, testConfig.MarketStore, testConfig.ChartStore, testConfig.TickStore, testConfig.Exchange, testConfig.OrderManager)

	m := &models.Market{
		MarketCurrency: "ltc",
		BaseCurrency:   "btc",
		MarketKey:      "btc-ltc",
	}

	buyChart := getBullishTestChart()
	testConfig.TickStore.EXPECT().GetChartCandles(ctx, m.MarketKey, testExchange, testInterval).Return(buyChart.GetCandles(), nil)
	testConfig.OrderManager.EXPECT().GetBalance(m.MarketCurrency).Return(nil)
	testConfig.OrderManager.EXPECT().GetBalance(m.BaseCurrency).Return(&orders.Balance{
		Available: utils.StringToDecimal("1"),
		Total:     utils.StringToDecimal("1"),
	})
	testConfig.Strategy.EXPECT().ShouldBuy(gomock.Any()).Return(true)
	ticker := &exchanges.Ticker{
		Bid:  utils.StringToDecimal("1"),
		Ask:  utils.StringToDecimal("1"),
		Last: utils.StringToDecimal("1"),
	}
	testConfig.Exchange.EXPECT().GetTicker(m.MarketKey).Return(ticker, nil)

	sellOrder := &orders.LimitOrderMatcher{
		Limit:          utils.StringToDecimal("1.01"),
		Quantity:       utils.StringToDecimal("0.01"),
		MarketCurrency: m.MarketCurrency,
		BaseCurrency:   m.BaseCurrency,
	}
	testConfig.OrderManager.EXPECT().ExecuteLimitBuy(ctx, sellOrder).Return(errors.New("some-error"))

	err := toTest.trade(ctx, m, testConfig.Strategy, testExchange, testInterval)
	assert.Equal(t, err, errors.New("some-error"), "should not error on decide call")
}
func TestTradeSellSuccess(t *testing.T) {}
func TestTradeSellFail(t *testing.T)    {}

func getBullishTestChart() *charts.CloudChart {
	var candles []*charts.Candle
	// Second to last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampOne,
		High:      utils.StringToDecimal("5"),
		Low:       utils.StringToDecimal("5"),
		Open:      utils.StringToDecimal("5"),
		Close:     utils.StringToDecimal("5"),
		Kijun:     utils.StringToDecimal("1"),
		Tenkan:    utils.StringToDecimal("0"),
	})
	// Last Candle
	candles = append(candles, &charts.Candle{
		TimeStamp: timestampTwo,
		High:      utils.StringToDecimal("6"),
		Low:       utils.StringToDecimal("6"),
		Open:      utils.StringToDecimal("6"),
		Close:     utils.StringToDecimal("6"),
		Kijun:     utils.StringToDecimal("0"),
		Tenkan:    utils.StringToDecimal("1"),
	})

	chartToTest := &charts.CloudChart{
		Test:  true,
		Cloud: make(map[int64]*charts.CloudPoint),
	}
	chartToTest.SetCandles(candles)

	return chartToTest
}
