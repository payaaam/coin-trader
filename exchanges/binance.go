package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
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

func (b *BinanceClient) ExecuteLimitBuy(tradingPair string, price string, quantity string) (string, error) {
	return "", nil
}
