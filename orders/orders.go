package orders

import (
	"errors"
	// "github.com/payaaam/coin-trader/db"
	//"github.com/payaaam/coin-trader/exchanges"
	//"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
)

type OrderManager interface {
	ExecuteLimitSell(ctx context.Context, change string, order *LimitOrder) error
	ExecuteLimitBuy(ctx context.Context, exchange string, order *LimitOrder) error
}

var BuyOrder = "buy"
var SellOrder = "sell"
var ErrNotEnoughFunds = errors.New("not enough coins")

// LimitOrder object to execute on an exchange
type LimitOrder struct {
	BaseCurrency   string
	MarketCurrency string
	Limit          decimal.Decimal
	Quantity       decimal.Decimal
}

type OpenOrder struct {
	Type      string
	MarketKey string
	Quantity  decimal.Decimal
	Limit     decimal.Decimal
	ID        string
}

type Balance struct {
	Total     decimal.Decimal
	Available decimal.Decimal
}

/*
type TraderState struct {
	Balances   map[string]decimal.Decimal
	OpenOrders []*OpenOrder
	Charts     map[string]*charts.CloudChart
}

func NewTraderState() *TraderState {
	return &TraderState{
		Balances:   make(map[string]decimal.Decimal),
		OpenOrders: []*OpenOrder{},
		Charts:     make(map[string]*charts.CloudChart),
	}
}

func (t *TraderState) SetBalance(marketKey string, amount decimal.Decimal) error {
	t.Balances[marketKey] = amount
	return nil
}

func (t *TraderState) GetBalance(marketKey string) decimal.Decimal {
	return t.Balances[marketKey]
}

func (t *TraderState) SetChart(marketKey string, chart *charts.CloudChart) error {
	t.Charts[marketKey] = chart
	return nil
}

func (t *TraderState) GetChart(marketKey string) *charts.CloudChart {
	return t.Charts[marketKey]
}
*/
