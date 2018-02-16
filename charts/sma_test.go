package charts

import (
	"testing"

	"github.com/payaaam/coin-trader/utils"
	"github.com/stretchr/testify/assert"
)

/* --- SMA TESTS --- */
func TestSMAExactPeriods(t *testing.T) {
	ma := NewSMA(2, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	assert.Equal(t, "2", ma.Avg().String())
}

func TestSMAExtraPeriods(t *testing.T) {
	ma := NewSMA(2, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	ma.Add(utils.StringToDecimal("5"))
	assert.Equal(t, "4", ma.Avg().String())
}

func TestSMANotEnoughData(t *testing.T) {
	ma := NewSMA(5, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	assert.Equal(t, NotEnoughData.String(), ma.Avg().String())
}

/* -- EMA TESTS --- */
func TestEMAExactPeriods(t *testing.T) {
	ma := NewSMA(2, true)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	assert.Equal(t, "2", ma.Avg().String())
}

func TestEMAExtraPeriods(t *testing.T) {
	ma := NewSMA(10, true)

	ma.Add(utils.StringToDecimal("22.27"))
	ma.Add(utils.StringToDecimal("22.19"))
	ma.Add(utils.StringToDecimal("22.08"))
	ma.Add(utils.StringToDecimal("22.17"))
	ma.Add(utils.StringToDecimal("22.18"))
	ma.Add(utils.StringToDecimal("22.13"))
	ma.Add(utils.StringToDecimal("22.23"))
	ma.Add(utils.StringToDecimal("22.43"))
	ma.Add(utils.StringToDecimal("22.24"))
	ma.Add(utils.StringToDecimal("22.29"))
	ma.Add(utils.StringToDecimal("22.15"))
	ma.Add(utils.StringToDecimal("22.39"))
	ma.Add(utils.StringToDecimal("22.38"))
	ma.Add(utils.StringToDecimal("22.61"))
	ma.Add(utils.StringToDecimal("23.36"))
	ma.Add(utils.StringToDecimal("24.05"))
	assert.Equal(t, "22.8", ma.Avg().Round(2).String())
}

func TestEMANotEnoughData(t *testing.T) {
	ma := NewSMA(5, true)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	assert.Equal(t, NotEnoughData.String(), ma.Avg().String())
}
