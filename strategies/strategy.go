package strategies

import (
	"github.com/payaaam/coin-trader/charts"
)

type Strategy interface {
	ShouldBuy(chart *charts.CloudChart) bool
	ShouldSell(chart *charts.CloudChart) bool
}
