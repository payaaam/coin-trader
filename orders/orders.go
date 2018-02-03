package orders

import (
	"errors"
	// "github.com/payaaam/coin-trader/db"
	//"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"time"
)

type OrderManager interface {
	Setup() error
	SetupSimulation(*Balance) error
	GetBalance(marketKey string) *Balance
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

var OrderOpenTimeoutMS = int64(10000)

// OrderType
var BuyOrder = "buy"
var SellOrder = "sell"

// OrderStatus
var OpenOrderStatus = "open"
var CancelledOrderStatus = "cancelled"
var FilledOrderStatus = "filled"
var PartiallyFieldOrderStatus = "partially-filled"

var ErrNotEnoughFunds = errors.New("not enough coins")
var ErrInvalidOrderType = errors.New("invalid order type")
var ErrInvalidOrderStatus = errors.New("invalid order status")
var ErrExecuteFailure = errors.New("error occured during trade execution")
var ErrRollbackFailure = errors.New("error occurd on rollback of execution")

// LimitOrder object to execute on an exchange
type LimitOrder struct {
	BaseCurrency   string
	MarketCurrency string
	Limit          decimal.Decimal
	Quantity       decimal.Decimal
}

type OpenOrder struct {
	Type           string
	MarketKey      string
	BaseCurrency   string
	MarketCurrency string
	OpenTimestamp  int64
	CloseTimestamp int64
	Quantity       decimal.Decimal
	Limit          decimal.Decimal
	TradePrice     decimal.Decimal
	QuantityFilled decimal.Decimal
	Status         string
	ID             string
}

func (o *OpenOrder) hasReachedTimeout() bool {
	timeDiff := time.Now().Unix() - o.OpenTimestamp
	if timeDiff >= OrderOpenTimeoutMS {
		return true
	}

	return false
}

func (o *OpenOrder) CancelOrder() {
	o.Status = FilledOrderStatus
	o.CloseTimestamp = time.Now().Unix()
	o.QuantityFilled = utils.ZeroDecimal()
	o.TradePrice = utils.ZeroDecimal()
}

type Balance struct {
	Total     decimal.Decimal
	Available decimal.Decimal
}
