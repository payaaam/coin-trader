package charts

import (
	"errors"
	//"fmt"
	//"github.com/fatih/color"
	//"github.com/shopspring/decimal"
	"sort"
	"time"

	log "github.com/sirupsen/logrus"
)

type CloudChart struct {
	Market               string
	Exchange             string
	Interval             string
	candles              []*Candle
	KijunMovingAverage   *MovingAverage
	TenkanMovingAverage  *MovingAverage
	SenkouBMovingAverage *MovingAverage
	SMA200               *SMA
	Cloud                map[int64]*CloudPoint
	Test                 bool
}

func NewCloudChartFromCandles(candles []*Candle, tradingPair string, exchange string, interval string) (*CloudChart, error) {

	sort.Slice(candles, func(i, j int) bool {
		return candles[i].TimeStamp < candles[j].TimeStamp
	})

	chart := &CloudChart{
		candles:              candles,
		Market:               tradingPair,
		Exchange:             exchange,
		Interval:             interval,
		KijunMovingAverage:   NewMovingAverage(KijunPeriod),
		TenkanMovingAverage:  NewMovingAverage(TenkanPeriod),
		SenkouBMovingAverage: NewMovingAverage(SenkouBPeriod),
		SMA200:               NewSMA(200, false),
		Cloud:                make(map[int64]*CloudPoint),
		Test:                 false,
	}

	for index, _ := range chart.candles {
		candle := chart.candles[index]
		chart.KijunMovingAverage.Add(candle.High, candle.Low)
		chart.TenkanMovingAverage.Add(candle.High, candle.Low)
		chart.SenkouBMovingAverage.Add(candle.High, candle.Low)
		chart.SMA200.Add(candle.Close)

		// Set Cloud
		candle.Kijun = chart.KijunMovingAverage.Avg()
		candle.Tenkan = chart.TenkanMovingAverage.Avg()
		candle.SMA200 = chart.SMA200.Avg()
		senkouB := chart.SenkouBMovingAverage.Avg()
		cloud := NewCloudPoint(candle.Tenkan, candle.Kijun, senkouB)
		chart.SetCloudPoint(candle.TimeStamp, cloud)
	}

	return chart, nil
}

func NewCloudChart(tradingPair string, exchange string, interval string) *CloudChart {
	chart := &CloudChart{
		candles:              []*Candle{},
		Market:               tradingPair,
		Exchange:             exchange,
		Interval:             interval,
		KijunMovingAverage:   NewMovingAverage(KijunPeriod),
		TenkanMovingAverage:  NewMovingAverage(TenkanPeriod),
		SenkouBMovingAverage: NewMovingAverage(SenkouBPeriod),
		SMA200:               NewSMA(200, false),
		Cloud:                make(map[int64]*CloudPoint),
		Test:                 false,
	}
	return chart
}

func (c *CloudChart) SetCandles(candles []*Candle) {
	c.candles = candles
}

func (c *CloudChart) AddCandle(candle *Candle) {
	lastCandle := c.GetLastCandle()
	if lastCandle != nil && lastCandle.TimeStamp == candle.TimeStamp {
		c.KijunMovingAverage.RemoveLast()
		c.TenkanMovingAverage.RemoveLast()
		c.SenkouBMovingAverage.RemoveLast()
		c.SMA200.RemoveLast()
		c.candles = c.candles[:len(c.candles)-1]
	}

	c.candles = append(c.candles, candle)
	c.KijunMovingAverage.Add(candle.High, candle.Low)
	c.TenkanMovingAverage.Add(candle.High, candle.Low)
	c.SenkouBMovingAverage.Add(candle.High, candle.Low)
	c.SMA200.Add(candle.Close)

	// Set Cloud
	candle.Kijun = c.KijunMovingAverage.Avg()
	candle.Tenkan = c.TenkanMovingAverage.Avg()
	candle.SMA200 = c.SMA200.Avg()
	senkouB := c.SenkouBMovingAverage.Avg()
	cloud := NewCloudPoint(candle.Tenkan, candle.Kijun, senkouB)
	c.SetCloudPoint(candle.TimeStamp, cloud)
}

func (c *CloudChart) GetCandles() []*Candle {
	return c.candles
}

func (c *CloudChart) SetCloudPoint(timestamp int64, cloud *CloudPoint) {
	leadingPeriodMs := int64(IntervalMilliseconds[c.Interval] * 30)
	c.Cloud[timestamp+leadingPeriodMs] = cloud
}

func (c *CloudChart) GetCloud(timestamp int64) (*CloudPoint, error) {
	if c.Cloud[timestamp] != nil {
		return c.Cloud[timestamp], nil
	}

	return nil, errors.New("No Cloud")
}

func (c *CloudChart) Print() {
	for _, candle := range c.candles {
		log.Infof("TimeStamp: %v", time.Unix(candle.TimeStamp, 0).UTC().Format("2006-01-02"))
		log.Infof("TimeStamp (unix): %v", candle.TimeStamp)

		log.Infof("Open: %v", candle.Open)
		log.Infof("Close: %v", candle.Close)
		log.Infof("Tenkan: %v", candle.Tenkan)
		log.Infof("Kijun: %v", candle.Kijun)
		log.Infof("200 SMA: %v", candle.SMA200)
		cloud, err := c.GetCloud(candle.TimeStamp)
		if err == nil {
			log.Infof("SenkouA: %v", cloud.SenkouA)
			log.Infof("SenkouB: %v", cloud.SenkouB)
			log.Infof("Displacement: %v", cloud.Displacement)
		}

		log.Info()
	}

	log.Infof("CLOUD LENGTH: %v", len(c.Cloud))
}

func (c *CloudChart) GetCandleCount() int {
	return len(c.candles)
}

func (c *CloudChart) GetLastCandle() *Candle {
	if len(c.candles) == 0 {
		return nil
	}
	lastCandleIndex := len(c.candles) - 1
	return c.candles[lastCandleIndex]
}

/*
func (c *CloudChart) PrintSummary() {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	three, _ := decimal.NewFromString("3")
	negThree, _ := decimal.NewFromString("-3")
	hundred, _ := decimal.NewFromString("100")

	boldBlue := color.New(color.FgBlue, color.Bold)
	boldBlue.Printf("Market: %s\n", c.Market)

	color.White("--- Summary ----")

	candle := c.GetLastCandle()
	var tkOrientation string
	if candle.Tenkan.GreaterThan(candle.Kijun) {
		tkOrientation = green("Tenkan Over Kijun")
	} else {
		tkOrientation = red("Kijun Over Tenkan")
	}
	fmt.Printf("TK Orientation:        %s\n", tkOrientation)

	cloud, err := c.GetCloud(candle.Day)
	if err == nil {
		var cc string
		if cloud.Color == "RED" {
			cc = red(cloud.Color)
		} else if cloud.Color == "GREEN" {
			cc = green(cloud.Color)
		} else {
			cc = yellow(cloud.Color)
		}
		fmt.Printf("Cloud Color:           %s\n", cc)
	}

	lastCross := FindLastTKCross(c)
	daysSinceLastCross := getDaysSinceLastCross(lastCross)

	if daysSinceLastCross < 4 {
		fmt.Printf("Last TK Cross Date:    %v\n", green(lastCross))
		fmt.Printf("Days Since Last Cross: %v\n", green(daysSinceLastCross))
	} else if daysSinceLastCross < 6 {
		fmt.Printf("Last TK Cross Date:    %v\n", yellow(lastCross))
		fmt.Printf("Days Since Last Cross: %v\n", yellow(daysSinceLastCross))
	} else {
		fmt.Printf("Last TK Cross Date:    %v\n", red(lastCross))
		fmt.Printf("Days Since Last Cross: %v\n", red(daysSinceLastCross))
	}

	nextCross := FindNextTKCross(c)
	var nc string
	if nextCross.LessThan(three) && nextCross.GreaterThan(negThree) {
		nc = green(nextCross)
	} else {
		nc = red(nextCross)
	}
	fmt.Printf("Next TK Cross:         %v\n", nc)

	dayChange := candle.Close.Sub(candle.Open)
	percentChange := dayChange.Div(candle.Open).Mul(hundred).Round(2)
	pcString := fmt.Sprintf("%s%%", percentChange)
	var dc string
	var pc string
	if dayChange.Sign() == 1 {
		dc = green(dayChange)
		pc = green(pcString)
	} else {
		dc = red(dayChange)
		pc = red(pcString)
	}
	fmt.Printf("Change:                %s\n", pc)
	fmt.Printf("Price Change:          %s\n", dc)

	fmt.Println("\n--- Candle ----")
	fmt.Printf("TimeStamp: %v\n", candle.TimeStamp)
	fmt.Printf("Open: %v\n", candle.Open)
	fmt.Printf("Close: %v\n", candle.Close)
	fmt.Printf("Tenkan: %v\n", candle.Tenkan)
	fmt.Printf("Kijun: %v\n", candle.Kijun)
	fmt.Println()
}
*/
