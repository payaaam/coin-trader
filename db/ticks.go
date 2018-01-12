package db

import (
	"database/sql"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/utils"
	//log "github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/net/context"
)

type TickStore struct {
	db *sql.DB
}

func NewTickStore(db *sql.DB) *TickStore {
	return &TickStore{
		db: db,
	}
}

func (t *TickStore) Upsert(ctx context.Context, chartID int, candle *charts.Candle) error {
	tick := &models.Tick{
		ChartID:   chartID,
		Open:      candle.Open.String(),
		Close:     candle.Close.String(),
		High:      candle.High.String(),
		Low:       candle.Low.String(),
		Volume:    candle.Volume.String(),
		Timestamp: candle.TimeStamp,
		Day:       candle.Day,
	}
	return tick.Upsert(t.db, true, []string{"chart_id", "timestamp"}, []string{"open", "close", "high", "low"})
}

func (t *TickStore) Save(ctx context.Context, tick *models.Tick) error {
	return tick.Insert(t.db)
}

func (t *TickStore) GetAllChartCandles(ctx context.Context, marketKey string, exchange string, interval string) ([]*charts.Candle, error) {
	ticks, err := models.Ticks(t.db,
		qm.Select("open", "close", "high", "low", "volume", "day", "timestamp"),
		qm.InnerJoin("chart ON chart.id = tick.chart_id"),
		qm.InnerJoin("market ON market.id = chart.market_id"),
		qm.Where("market.exchange_name = ?", exchange),
		qm.And("market.market_key = ?", marketKey),
		qm.And("chart.interval = ?", interval),
		qm.OrderBy("timestamp ASC"),
	).All()

	if err != nil {
		return nil, err
	}

	var candles []*charts.Candle
	for _, tick := range ticks {
		candles = append(candles, &charts.Candle{
			TimeStamp: tick.Timestamp,
			Day:       tick.Day,
			Open:      utils.StringToDecimal(tick.Open),
			Close:     utils.StringToDecimal(tick.Close),
			High:      utils.StringToDecimal(tick.High),
			Low:       utils.StringToDecimal(tick.Low),
			Volume:    utils.StringToDecimal(tick.Volume),
		})
	}

	return candles, nil
}

func (t *TickStore) GetChartCandles(ctx context.Context, marketKey string, exchange string, interval string) ([]*charts.Candle, error) {
	ticks, err := models.Ticks(t.db,
		qm.Select("open", "close", "high", "low", "volume", "day", "timestamp"),
		qm.InnerJoin("chart ON chart.id = tick.chart_id"),
		qm.InnerJoin("market ON market.id = chart.market_id"),
		qm.Where("market.exchange_name = ?", exchange),
		qm.And("market.market_key = ?", marketKey),
		qm.And("chart.interval = ?", interval),
		qm.OrderBy("timestamp DESC"),
		qm.Limit(200),
	).All()

	if err != nil {
		return nil, err
	}

	var candles []*charts.Candle
	for _, tick := range ticks {
		candles = append(candles, &charts.Candle{
			TimeStamp: tick.Timestamp,
			Day:       tick.Day,
			Open:      utils.StringToDecimal(tick.Open),
			Close:     utils.StringToDecimal(tick.Close),
			High:      utils.StringToDecimal(tick.High),
			Low:       utils.StringToDecimal(tick.Low),
			Volume:    utils.StringToDecimal(tick.Volume),
		})
	}

	return candles, nil
}
