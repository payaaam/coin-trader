package exchanges

import (
	"fmt"
	"github.com/payaaam/coin-trader/charts"
	"github.com/shopspring/decimal"
)

// Potential Libraries
// https://github.com/adshao/go-binance
// https://github.com/pdepip/go-binance

type BinanceClient struct {
}

func NewBinanceClient() Exchange {
	return &BinanceClient{}
}

func (b *BinanceClient) GetCandles(tradingPair string, chartInterval string) ([]*charts.Candle, error) {
	return []*charts.Candle{}, nil
}

func (b *BinanceClient) GetLatestCandle(tradingPair string, chartInterval string) (*charts.Candle, error) {
	return &charts.Candle{}, nil
}

func (b *BinanceClient) GetBitcoinMarkets() ([]*Market, error) {

	return []*Market{}, nil
}

func (b *BinanceClient) ExecuteLimitBuy(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error) {
	return "", nil
}

func (b *BinanceClient) ExecuteLimitSell(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error) {
	return "", nil
}
func (b *BinanceClient) GetBalances() ([]*Balance, error) {
	return nil, nil
}

func (b *BinanceClient) GetMarketKey(base string, market string) string {
	return fmt.Sprintf("%s%s", base, market)
}
