package orders

import (
	"errors"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
)

type OrderManager interface {
	ExecuteLimitSell(exchange string, order *LimitOrder) error
	ExecuteLimitBuy(exchange string, order *LimitOrder) error
}

var ErrNotEnoughFunds = errors.New("not enough coins")

// LimitOrder object to execute on an exchange
type LimitOrder struct {
	BaseCurrency   string
	MarketCurrency string
	Limit          decimal.Decimal
	Amount         decimal.Decimal
}

type OpenOrder struct {
	Type      string
	MarketKey string
	Quantity  decimal.Decimal
	Limit     decimal.Decimal
	UUID      string
}

type Manager struct {
	Balances   map[string]decimal.Decimal
	OpenOrders []*OpenOrder
	client     exchanges.Exchange
	orderStore *db.OrderStore
}

func NewManager(client exchanges.Exchange, os *db.OrderStore) (OrderManager, error) {
	manager := &Manager{
		Balances:   make(map[string]decimal.Decimal),
		OpenOrders: []*OpenOrder{},
		client:     client,
		orderStore: os,
	}

	err := manager.loadBalances()
	if err != nil {
		return nil, err
	}

	return manager, nil
}

func (m *Manager) ExecuteLimitSell(exchange string, order *LimitOrder) error {
	balance := m.getBalance(order.MarketCurrency)
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}


	m.createOpenSellOrder()
	// Update Balance Available
	// Add Sell Order to DB

	return  nil
}


func (m *Manager) ExecuteLimitBuy(exchange string, order *LimitOrder) error {
	balance := m.getBalance(order.BaseCurrency)
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	m.createOpenBuyOrder()
	// Update Balance Available
	// Add Buy Order to DB


	return nil
}

func (m *Manager) loadBalances() error {
	// Get your balances from exchange
	balances, err := m.client.GetBalances()
	if err != nil {
		return err
	}

	for _, balance := range balances {
		m.setBalance(balance.BaseCurrency, balance.Amount)
	}

	return nil
}

func (m *Manager) setBalance(marketKey string, amount decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)] = amount
}

func (m *Manager) getBalance(marketKey string) decimal.Decimal {
	return m.Balances[[utils.Normalize(marketKey)]
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
