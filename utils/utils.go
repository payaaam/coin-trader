package utils

import (
	"strings"

	"github.com/shopspring/decimal"
)

func Normalize(s string) string {
	return strings.TrimSpace(strings.ToLower(s))
}

func StringToDecimal(s string) decimal.Decimal {
	d, _ := decimal.NewFromString(s)
	return d
}

func IntToDecimal(i int) decimal.Decimal {
	d := decimal.New(int64(i), 0)
	return d
}

func ZeroDecimal() decimal.Decimal {
	return decimal.New(0, 0)
}
