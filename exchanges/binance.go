package exchanges

import (
	"github.com/payaaam/coin-trader/charts"
	"github.com/adshao/go-binance"
	"golang.org/x/net/context"
	"strings"
	"fmt"
)

type BinanceClient struct {
	client *binance.Client
}

func NewBinanceClient(client *binance.Client) *BinanceClient {
	return &BinanceClient{
		client: client,
	}
}

func (b *BinanceClient) GetCandles(tradingPair string, chartInterval string) ([]*charts.Candle, error) {
	return []*charts.Candle{}, nil
}

func (b *BinanceClient) GetLatestCandle(tradingPair string, chartInterval string) (*charts.Candle, error) {
	return &charts.Candle{}, nil
}

func (b *BinanceClient) GetBitcoinMarkets() ([]*Market, error) {
	ctx := context.Background()
	markets, err := b.client.NewListPricesService().Do(ctx)
	if err != nil {
		fmt.Println(err);
		return nil, err
	}

	var binanceMarkets []*Market
	for _, market := range markets {
		if strings.HasSuffix(market.Symbol, "BTC") {
			fmt.Println(market.Symbol)
			/*binanceMarkets = append(binanceMarkets, &Market{
				MarketKey:          market.MarketName,
				BaseCurrency:       market.BaseCurrency,
				MarketCurrency:     market.MarketCurrency,
				BaseCurrencyName:   market.BaseCurrencyLong,
				MarketCurrencyName: market.MarketCurrencyLong,
			})*/
		}
	}

	return binanceMarkets, nil
}

func (b *BinanceClient) ExecuteLimitBuy(tradingPair string, price string, quantity string) (string, error) {
	return "", nil
}
