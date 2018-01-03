package charts

import (
	"github.com/shopspring/decimal"
)

type MovingAverage struct {
	values []*HighLow
	period int
}

type HighLow struct {
	high decimal.Decimal
	low  decimal.Decimal
}

func NewMovingAverage(period int) *MovingAverage {
	return &MovingAverage{
		period: period,
		values: []*HighLow{},
	}
}

func (m *MovingAverage) Add(high decimal.Decimal, low decimal.Decimal) {
	if len(m.values) == m.period {
		m.values = append(m.values[:0], m.values[1:]...)
	}

	// Add value
	m.values = append(m.values, &HighLow{
		high: high,
		low:  low,
	})
}

func (m *MovingAverage) Avg() decimal.Decimal {
	zero, _ := decimal.NewFromString("0")
	if len(m.values) != m.period {
		return zero
	}

	var high = m.values[0].high
	var low = m.values[0].low
	for _, value := range m.values {

		if value.high.GreaterThan(high) {
			high = value.high
		}

		if value.low.LessThan(low) {
			low = value.low
		}
	}

	two, _ := decimal.NewFromString("2")
	return high.Add(low).Div(two)
}
