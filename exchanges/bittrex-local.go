package exchanges

import (
	"encoding/json"
	"github.com/payaaam/coin-trader/charts"
	"github.com/toorop/go-bittrex"
	"io/ioutil"
)

/*
	localBittrexClient := exchanges.NewLocalClient()
	chart, err := localBittrexClient.GetCandles(JSONPath, "BTC-ETH")
	if err != nil {
		panic(err)
	}
*/
type BittrexLocalClient struct{}

func NewBittrexLocalClient() *BittrexLocalClient {
	return &BittrexLocalClient{}
}

func (b *BittrexLocalClient) GetCandles(path string, tradingPair string) ([]*charts.Candle, error) {

	candles, err := readJSON(path)
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

	return chartCandles, nil
}

type jsonResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Result  json.RawMessage `json:"result"`
}

func readJSON(path string) ([]bittrex.Candle, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var response jsonResponse
	if err := json.Unmarshal(raw, &response); err != nil {
		return nil, err
	}

	var candles []bittrex.Candle
	if err := json.Unmarshal(response.Result, &candles); err != nil {
		return nil, err
	}

	return candles, nil
}
