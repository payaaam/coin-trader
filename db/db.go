package db

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db/models"
	"golang.org/x/net/context"
)

var OpenOrderStatus = "open"
var CompletedOrderStatus = "closed"
var CancelledOrderStatus = "cancelled"

var OneDayInterval string = "1D"
var FourHourInterval string = "4h"
var OneHourInterval string = "1h"
var ThirtyMinuteInterval string = "30m"
var MinuteMillisecond = 60000

var ValidIntervals = map[string]bool{
	OneDayInterval:       true,
	FourHourInterval:     true,
	OneHourInterval:      true,
	ThirtyMinuteInterval: true,
}

var IntervalMilliseconds = map[string]int{
	OneDayInterval:       MinuteMillisecond * 60 * 24,
	FourHourInterval:     MinuteMillisecond * 60 * 4,
	OneHourInterval:      MinuteMillisecond * 60,
	ThirtyMinuteInterval: MinuteMillisecond * 30,
}

type OrderStoreInterface interface {
	Upsert(ctx context.Context, order *models.Order) error
	Save(ctx context.Context, order *models.Order) error
}

type ChartStoreInterface interface {
	Upsert(ctx context.Context, chart *models.Chart) error
	Save(ctx context.Context, chart *models.Chart) error
}

type MarketStoreInterface interface {
	Upsert(ctx context.Context, market *models.Market) error
	Save(ctx context.Context, market *models.Market) error
	GetMarkets(ctx context.Context, exchange string) ([]*models.Market, error)
	GetMarket(ctx context.Context, exchangeName string, marketKey string) (*models.Market, error)
}

type TickStoreInterface interface {
	Upsert(ctx context.Context, chartID int, candle *charts.Candle) error
	Save(ctx context.Context, tick *models.Tick) error
	GetAllChartCandles(ctx context.Context, marketKey string, exchange string, interval string) ([]*charts.Candle, error)
	GetChartCandles(ctx context.Context, marketKey string, exchange string, interval string) ([]*charts.Candle, error)
	GetLatestChartCandle(ctx context.Context, chartID int) (*charts.Candle, error)
	GetCandlesFromRange(ctx context.Context, chartID int, start int64, end int64) ([]*charts.Candle, error)
}
