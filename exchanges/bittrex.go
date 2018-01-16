package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/toorop/go-bittrex"
)

type BittrexClient struct {
	client *bittrex.Bittrex
}

func NewBittrexClient(client *bittrex.Bittrex) Exchange {
	return &BittrexClient{
		client: client,
	}
}

func (b *BittrexClient) GetCandles(tradingPair string, chartInterval string) ([]*charts.Candle, error) {
	candles, err := b.client.GetTicks(tradingPair, chartInterval)
	if err != nil {
		return nil, err
	}

	var chartCandles []*charts.Candle
	for _, candle := range candles {
		chartCandles = append(chartCandles, &charts.Candle{
			TimeStamp: candle.TimeStamp.Time.Unix(),
			Open:      candle.Open,
			Close:     candle.Close,
			High:      candle.High,
			Low:       candle.Low,
			Volume:    candle.Volume,
		})
	}

	return chartCandles, nil
}

func (b *BittrexClient) GetLatestCandle(tradingPair string, chartInterval string) (*charts.Candle, error) {
	candles, err := b.client.GetLatestTick(tradingPair, chartInterval)
	if err != nil {
		return nil, err
	}

	var chartCandles []*charts.Candle
	for _, candle := range candles {
		chartCandles = append(chartCandles, &charts.Candle{
			TimeStamp: candle.TimeStamp.Time.Unix(),
			Open:      candle.Open,
			Close:     candle.Close,
			High:      candle.High,
			Low:       candle.Low,
			Volume:    candle.Volume,
		})
	}

	return chartCandles[0], nil
}

func (b *BittrexClient) GetBitcoinMarkets() ([]*Market, error) {
	markets, err := b.client.GetMarkets()
	if err != nil {
		return nil, err
	}

	var bittrexMarkets []*Market
	for _, market := range markets {
		if market.IsActive == false {
			continue
		}
		if market.BaseCurrency == "BTC" {
			bittrexMarkets = append(bittrexMarkets, &Market{
				MarketKey:          market.MarketName,
				BaseCurrency:       market.BaseCurrency,
				MarketCurrency:     market.MarketCurrency,
				BaseCurrencyName:   market.BaseCurrencyLong,
				MarketCurrencyName: market.MarketCurrencyLong,
			})
		}
	}

	return bittrexMarkets, nil
}

func (b *BittrexClient) ExecuteLimitBuy(tradingPair string, price string, quantity string) (string, error) {
	return "", nil
}
