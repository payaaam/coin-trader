package charts

import (
	"errors"
	//"fmt"
	//"github.com/fatih/color"
	//"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"sort"
	//"time"
)

const (
	TenkanPeriod       = 20
	KijunPeriod        = 60
	SenkouBPeriod      = 120
	CloudLeadingPeriod = 30
)

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

	sort.Slice(candles, func(i, j int) bool {
		return candles[i].TimeStamp < candles[j].TimeStamp
	})

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

func (c *CloudChart) AddCandle(candle *Candle) {
	if c.GetLastCandle().TimeStamp == candle.TimeStamp {
		c.KijunMovingAverage.RemoveLast()
		c.TenkanMovingAverage.RemoveLast()
		c.SenkouBMovingAverage.RemoveLast()
	}
	c.KijunMovingAverage.Add(candle.High, candle.Low)
	c.TenkanMovingAverage.Add(candle.High, candle.Low)
	c.SenkouBMovingAverage.Add(candle.High, candle.Low)

	// Set Cloud
	candle.Kijun = c.KijunMovingAverage.Avg()
	candle.Tenkan = c.TenkanMovingAverage.Avg()
	senkouB := c.SenkouBMovingAverage.Avg()
	cloud := NewCloudPoint(candle.Tenkan, candle.Kijun, senkouB)
	c.SetCloudPoint(candle.Day, cloud)
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

/*
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

func getDaysSinceLastCross(lc int64) int {

	last := time.Unix(lc, 0)
	now := time.Now()
	diff := now.Sub(last)
	return int(diff.Hours() / 24)
}
*/
