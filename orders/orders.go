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
	GetBalances() map[string]*Balance
	ExecuteLimitSell(ctx context.Context, order *LimitOrder) error
	ExecuteLimitBuy(ctx context.Context, order *LimitOrder) error
}

type OrderMonitor interface {
	Start(chan *OpenOrder)
	process()
	GetOrders() []*OpenOrder
	Execute(order *OpenOrder) (string, error)
}

// OrderType
var BuyOrder = "buy"
var SellOrder = "sell"

// OrderSTatus
var OpenOrderStatus = "open"
var FilledOrderStatus = "filled"
var PartiallyFieldOrderStatus = "partially-filled"

var ErrNotEnoughFunds = errors.New("not enough coins")
var ErrInvalidOrderType = errors.New("invalid order type")
var ErrInvalidOrderStatus = errors.New("invalid order status")

// LimitOrder object to execute on an exchange
type LimitOrder struct {
	BaseCurrency   string
	MarketCurrency string
	Limit          decimal.Decimal
	Quantity       decimal.Decimal
}

type OpenOrder struct {
	Type                 string
	MarketKey            string
	BaseCurrency         string
	MarketCurrency       string
	OrderPlacedTimestamp int64
	OrderFilledTimestamp int64
	Quantity             decimal.Decimal
	Limit                decimal.Decimal
	TradePrice           decimal.Decimal
	Status               string
	ID                   string
}

type Balance struct {
	Total     decimal.Decimal
	Available decimal.Decimal
}
