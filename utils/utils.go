package utils

import (
	"github.com/shopspring/decimal"
	"strings"
)

func Normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func StringToDecimal(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

func ZeroDecimal() decimal.Decimal {
	return decimal.New(0, 0)
}
