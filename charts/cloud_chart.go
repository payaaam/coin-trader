package charts

import (
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	TenkanPeriod       = 20
	KijunPeriod        = 60
	SenkouBPeriod      = 120
	CloudLeadingPeriod = 30
)

type Candle struct {
	TimeStamp time.Time
	Day       int
	Open      decimal.Decimal
	Close     decimal.Decimal
	High      decimal.Decimal
	Low       decimal.Decimal
	Volume    decimal.Decimal
	Kijun     decimal.Decimal
	Tenkan    decimal.Decimal
}

type CloudPoint struct {
	SenkouA      decimal.Decimal
	SenkouB      decimal.Decimal
	Color        string
	Displacement decimal.Decimal
}

func NewCloudPoint(tenkan decimal.Decimal, kijun decimal.Decimal, senkouB decimal.Decimal) *CloudPoint {
	two, _ := decimal.NewFromString("2")
	senkouA := tenkan.Add(kijun).Div(two)
	displacement := senkouA.Sub(senkouB)

	return &CloudPoint{
		SenkouA:      senkouA,
		SenkouB:      senkouB,
		Displacement: displacement,
		Color:        getCloudColor(senkouA, senkouB),
	}
}

type CloudChart struct {
	Market               string
	Exchange             string
	candles              []*Candle
	KijunMovingAverage   *MovingAverage
	TenkanMovingAverage  *MovingAverage
	SenkouBMovingAverage *MovingAverage
	Cloud                map[int]*CloudPoint
}

func NewCloudChart(candles []*Candle, tradingPair string, exchange string) (*CloudChart, error) {
	chart := &CloudChart{
		candles:              candles,
		Market:               tradingPair,
		Exchange:             exchange,
		KijunMovingAverage:   NewMovingAverage(KijunPeriod),
		TenkanMovingAverage:  NewMovingAverage(TenkanPeriod),
		SenkouBMovingAverage: NewMovingAverage(SenkouBPeriod),
		Cloud:                make(map[int]*CloudPoint),
	}

	for day, _ := range chart.candles {
		candle := chart.candles[day]
		chart.KijunMovingAverage.Add(candle.High, candle.Low)
		chart.TenkanMovingAverage.Add(candle.High, candle.Low)
		chart.SenkouBMovingAverage.Add(candle.High, candle.Low)

		// Set Cloud
		candle.Kijun = chart.KijunMovingAverage.Avg()
		candle.Tenkan = chart.TenkanMovingAverage.Avg()
		senkouB := chart.SenkouBMovingAverage.Avg()
		cloud := NewCloudPoint(candle.Tenkan, candle.Kijun, senkouB)
		chart.SetCloudPoint(day, cloud)
	}

	return chart, nil
}

func (c *CloudChart) GetCandles() []*Candle {
	return c.candles
}

func (c *CloudChart) SetCloudPoint(day int, cloud *CloudPoint) {
	c.Cloud[day+CloudLeadingPeriod] = cloud
}

func (c *CloudChart) GetCloud(day int) (*CloudPoint, error) {
	if c.Cloud[day] != nil {
		return c.Cloud[day], nil
	}

	return nil, errors.New("No Cloud")
}

func (c *CloudChart) Print() {
	for day, candle := range c.candles {
		log.Infof("Day: %d -----", day)
		log.Infof("TimeStamp: %v", candle.TimeStamp)
		log.Infof("Open: %v", candle.Open)
		log.Infof("Close: %v", candle.Close)
		log.Infof("Tenkan: %v", candle.Tenkan)
		log.Infof("Kijun: %v", candle.Kijun)
		cloud, err := c.GetCloud(day)
		if err == nil {
			log.Infof("SenkouA: %v", cloud.SenkouA)
			log.Infof("SenkouB: %v", cloud.SenkouB)
			log.Infof("Displacement: %v", cloud.Displacement)
		}

		log.Info()
	}
}

func (c *CloudChart) GetLastCandle() *Candle {
	lastCandleIndex := len(c.candles) - 1
	return c.candles[lastCandleIndex]
}

func (c *CloudChart) PrintSummary() {
	log.Infof("Market: %s", c.Market)

	log.Info("--- Summary ----")

	candle := c.GetLastCandle()
	if candle.Tenkan.GreaterThan(candle.Kijun) {
		log.Info("Tenkan Over Kijun")
	} else {
		log.Info("Kijun Over Tenkan")
	}

	cloud, err := c.GetCloud(candle.Day)
	if err == nil {
		log.Infof("Cloud Color: %s", cloud.Color)
	}

	lastCross := FindLastTKCross(c)
	log.Infof("Last TK Cross Date: %v", lastCross)

	nextCross := FindNextTKCross(c)
	log.Infof("Next TK Cross: %v", nextCross)

	log.Info("--- Candle ----")
	log.Infof("TimeStamp: %v", candle.TimeStamp)
	log.Infof("Open: %v", candle.Open)
	log.Infof("Close: %v", candle.Close)
	log.Infof("Tenkan: %v", candle.Tenkan)
	log.Infof("Kijun: %v", candle.Kijun)
	log.Info()
}

func getCloudColor(senkouA decimal.Decimal, senkouB decimal.Decimal) string {
	zero, _ := decimal.NewFromString("0")
	if senkouB.Equals(zero) {
		return "N/A"
	}

	if senkouA.GreaterThan(senkouB) {
		return "GREEN"
	}

	if senkouB.GreaterThan(senkouA) {
		return "RED"
	}

	if senkouA.Equals(senkouB) {
		return "NONE"
	}

	return "N/A"
}
