package charts

import (
	"github.com/shopspring/decimal"
)

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
		//Color:        getCloudColor(senkouA, senkouB),
	}
}
