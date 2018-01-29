package charts

import (
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
)

var NotEnoughData = utils.StringToDecimal("-1")

type SMA struct {
	values      []decimal.Decimal
	period      int
	exponential bool
}

func NewSMA(period int, exponential bool) *SMA {
	return &SMA{
		period:      period,
		values:      []decimal.Decimal{},
		exponential: exponential,
	}
}

func (m *SMA) Add(close decimal.Decimal) {
	if len(m.values) == m.period {
		// m.values = append(m.values[:0], m.values[1:]...)
		m.values = m.values[1:]
	}

	// Add value
	m.values = append(m.values, close)
}

func (m *SMA) RemoveLast() {
	if len(m.values) == 0 {
		return
	}
	m.values = m.values[:len(m.values)-1]
}

func (m *SMA) Avg() decimal.Decimal {
	total, _ := decimal.NewFromString("0")
	if len(m.values) != m.period {
		return NotEnoughData
	}

	for _, value := range m.values {
		total = total.Add(value)
	}

	sma := total.Div(decimal.NewFromFloat(float64(m.period)))
	return sma
}
