package charts

import (
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
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
	var ma decimal.Decimal
	total, _ := decimal.NewFromString("0")
	if len(m.values) < m.period {
		return NotEnoughData
	}

	if m.exponential {
		// Exponential Moving Average
		periodpp := utils.IntToDecimal(m.period + 1)
		mult := utils.StringToDecimal("2").Div(periodpp)

		var ema []decimal.Decimal

		// First, get the simple moving average for the first N periods
		for _, value := range m.values[0:m.period] {
			total = total.Add(value)
		}

		log.Info(total.String())

		// Add the initial SMA to the EMA array
		ema = append(ema, total.Div(decimal.NewFromFloat(float64(m.period))))

		// Calculate the full EMA
		for i := m.period; i < len(m.values); i++ {
			close := m.values[i]
			prevEma := ema[len(ema)-1:][0]
			newEma := close.Sub(prevEma).Mul(mult).Add(prevEma)
			ema = append(ema, newEma)
		}

		// Grab the last EMA value
		ma = ema[len(ema)-1:][0]
	} else {
		// Simple Moving Average
		for _, value := range m.values[len(m.values)-m.period:] {
			total = total.Add(value)
		}

		ma = total.Div(decimal.NewFromFloat(float64(m.period)))
	}
	return ma
}
