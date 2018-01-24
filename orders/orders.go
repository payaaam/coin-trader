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
	Setup() error
	GetOpenOrders() []*OpenOrder
	GetBalances() map[string]*Balance
	ExecuteLimitSell(ctx context.Context, order *LimitOrder) error
	ExecuteLimitBuy(ctx context.Context, order *LimitOrder) error
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
