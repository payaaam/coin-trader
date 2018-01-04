package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/toorop/go-bittrex"
)

type BittrexClient struct {
	client *bittrex.Bittrex
}

func NewClient(client *bittrex.Bittrex) *BittrexClient {
	return &BittrexClient{
		client: client,
	}
}

func (b *BittrexClient) GetCandles(tradingPair string, chartInterval string) (*charts.CloudChart, error) {
	candles, err := b.client.GetTicks(tradingPair, chartInterval)
	if err != nil {
		return nil, err
	}

	var chartCandles []*charts.Candle
	for day, candle := range candles {
		chartCandles = append(chartCandles, &charts.Candle{
			TimeStamp: candle.TimeStamp.Time,
			Day:       day,
			Open:      candle.Open,
			Close:     candle.Close,
			High:      candle.High,
			Low:       candle.Low,
			Volume:    candle.Volume,
		})
	}

	return charts.NewCloudChart(chartCandles, tradingPair, "Bittrex")
}

func (b *BittrexClient) GetMarkets() ([]*Market, error) {
	markets, err := b.client.GetMarkets()
	if err != nil {
		return nil, err
	}

	var bittrexMarkets []*Market
	for _, market := range markets {
		if market.BaseCurrency == "BTC" {
			bittrexMarkets = append(bittrexMarkets, &Market{
				TradingPair:    market.MarketName,
				BaseCurrency:   market.BaseCurrency,
				MarketCurrency: market.MarketCurrency,
			})
		}
	}

	return bittrexMarkets, nil
}
