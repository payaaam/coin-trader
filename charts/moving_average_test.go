package charts

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

var Highs = []string{"100", "200", "300", "400"}
var Lows = []string{"1", "2", "3", "4"}
var Expected2 = []string{"0", "100.5", "151", "201.5"}
var Expected3 = []string{"0", "0", "150.5", "202"}
var Expected4 = []string{"0", "0", "0", "200.5"}

func TestMovingAverageSize2(t *testing.T) {
	ma := NewMovingAverage(2)

	for i := 0; i < len(Highs)-1; i++ {
		h, _ := decimal.NewFromString(Highs[i])
		l, _ := decimal.NewFromString(Lows[i])
		ma.Add(h, l)
		assert.Equal(t, Expected2[i], ma.Avg().String())
	}
}

func TestMovingAverageSize3(t *testing.T) {
	ma := NewMovingAverage(3)

	for i := 0; i < len(Highs)-1; i++ {
		h, _ := decimal.NewFromString(Highs[i])
		l, _ := decimal.NewFromString(Lows[i])
		ma.Add(h, l)
		assert.Equal(t, Expected3[i], ma.Avg().String())
	}
}

func TestMovingAverageSize4(t *testing.T) {
	ma := NewMovingAverage(4)

	for i := 0; i < len(Highs)-1; i++ {
		h, _ := decimal.NewFromString(Highs[i])
		l, _ := decimal.NewFromString(Lows[i])
		ma.Add(h, l)
		assert.Equal(t, Expected4[i], ma.Avg().String())
	}
}
