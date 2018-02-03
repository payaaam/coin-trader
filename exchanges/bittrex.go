package exchanges

import (
	"fmt"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	"github.com/toorop/go-bittrex"
	"time"
)

var orderTypeMap = map[string]string{
	"LIMIT_BUY":  "buy",
	"LIMIT_SELL": "sell",
}

var openStatus = "open"
var filledStatus = "filled"

var BittrexTimestamp = "2006-01-02T15:04:05.00"

type BittrexClient struct {
	client *bittrex.Bittrex
}

func NewBittrexClient(client *bittrex.Bittrex) Exchange {
	return &BittrexClient{
		client: client,
	}
}

func (b *BittrexClient) GetTicker(tradingPair string) (*Ticker, error) {
	ticker, err := b.client.GetTicker(tradingPair)
	if err != nil {
		return nil, err
	}

	return &Ticker{
		Bid:  ticker.Bid,
		Ask:  ticker.Ask,
		Last: ticker.Last,
	}, nil
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

func (b *BittrexClient) ExecuteLimitBuy(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error) {
	return "", nil
}

func (b *BittrexClient) ExecuteLimitSell(tradingPair string, price decimal.Decimal, quantity decimal.Decimal) (string, error) {
	return "", nil
}

func (b *BittrexClient) GetBalances() ([]*Balance, error) {
	return nil, nil
}

func (b *BittrexClient) GetMarketKey(base string, market string) string {
	return fmt.Sprintf("%s-%s", base, market)
}

func (b *BittrexClient) GetOrder(orderID string) (*Order, error) {
	order, err := b.client.GetOrder(orderID)
	if err != nil {
		return nil, err
	}

	return &Order{
		Type:           orderTypeMap[order.Type],
		MarketKey:      utils.Normalize(order.Exchange),
		OpenTimestamp:  convertTime(order.Opened),
		CloseTimestamp: convertTime(order.Closed),
		Quantity:       order.Quantity,
		QuantityFilled: order.QuantityRemaining,
		Limit:          order.Limit,
		TradePrice:     order.Price,
	}, nil
}

func (b *BittrexClient) CancelOrder(orderID string) error {
	return nil
}

func convertTime(timestamp string) int64 {
	if timestamp == "" {
		return 0
	}
	t, _ := time.Parse(BittrexTimestamp, timestamp)
	return t.UTC().Unix()
}
