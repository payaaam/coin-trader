package charts

import (
	"testing"

	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestSMAExactPeriods(t *testing.T) {
	ma := NewSMA(2, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	log.Info(ma.Avg())
	assert.Equal(t, ma.Avg().String(), "2")
}

func TestSMAExtraPeriods(t *testing.T) {
	ma := NewSMA(2, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	ma.Add(utils.StringToDecimal("5"))
	log.Info(ma.Avg())
	assert.Equal(t, ma.Avg().String(), "4")
}

func TestSMANotEnoughData(t *testing.T) {
	ma := NewSMA(5, false)

	ma.Add(utils.StringToDecimal("1"))
	ma.Add(utils.StringToDecimal("3"))
	log.Info(ma.Avg())
	assert.Equal(t, ma.Avg().String(), NotEnoughData.String())
}
