package orders

import (
	"errors"
	"github.com/payaaam/coin-trader/db"
	//"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	"golang.org/x/net/context"
	"time"
)

type Manager struct {
	Balances   map[string]*Balance
	OpenOrders []*OpenOrder
	client     exchanges.Exchange
	orderStore db.OrderStoreInterface
}

func NewManager(client exchanges.Exchange, os db.OrderStoreInterface) OrderManager {
	manager := &Manager{
		Balances:   make(map[string]*Balance),
		OpenOrders: []*OpenOrder{},
		client:     client,
		orderStore: os,
	}

	return manager
}

// Loads balances and start the open order monitor
func (m *Manager) Setup() error {
	err := m.loadBalances()
	if err != nil {
		return err
	}

	// Start goroutine for processing open orders if they exist
	go m.startOpenOrderMonitor()
	return nil
}

func (m *Manager) startOpenOrderMonitor() {
	ticker := time.NewTicker(time.Second * 10)
	for _ = range ticker.C {
		if len(m.OpenOrders) == 0 {
			continue
		}

		m.processOpenOrders()
	}
}

func (m *Manager) loadBalances() error {
	// Get your balances from exchange
	balances, err := m.client.GetBalances()
	if err != nil {
		return err
	}

	for _, balance := range balances {
		m.setBalance(balance.BaseCurrency, balance.Total, balance.Available)
	}

	return nil
}

func (m *Manager) ExecuteLimitBuy(ctx context.Context, exchange string, order *LimitOrder) error {
	balance := m.getBalance(order.BaseCurrency)
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromOrder(BuyOrder, order)
	if err != nil {
		return err
	}
	m.createOpenBuyOrder(ctx, order)
	return nil
}

func (m *Manager) ExecuteLimitSell(ctx context.Context, exchange string, order *LimitOrder) error {
	balance := m.getBalance(order.MarketCurrency)
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromOrder(SellOrder, order)
	if err != nil {
		return err
	}
	m.createOpenSellOrder(ctx, order)

	return nil
}

func (m *Manager) processOpenOrders() {
	/*
		for _, openOrder := range m.OpenOrders {
			// Get Order Status from bittrex
			// If order is closed
			//  - Update Balance
			//  - Update Database
			//  - Check for Timeout
		}
	*/
}

func (m *Manager) createOpenBuyOrder(ctx context.Context, order *LimitOrder) error {
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	orderID, err := m.client.ExecuteLimitBuy(marketKey, order.Limit, order.Quantity)
	if err != nil {
		return err
	}

	// Create Open Order and add it to Open Orders
	openOrder := &OpenOrder{
		Type:      BuyOrder,
		MarketKey: marketKey,
		ID:        orderID,
		Limit:     order.Limit,
		Quantity:  order.Quantity,
	}
	m.addOpenOrder(openOrder)

	// Save Order to the database
	orderModel := convertToOrderModel(openOrder)
	err = m.orderStore.Save(ctx, orderModel)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) createOpenSellOrder(ctx context.Context, order *LimitOrder) error {
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)
	orderID, err := m.client.ExecuteLimitSell(marketKey, order.Limit, order.Quantity)
	if err != nil {
		return err
	}

	// Create Open Order and add it to Open Orders
	openOrder := &OpenOrder{
		Type:      SellOrder,
		MarketKey: marketKey,
		ID:        orderID,
		Limit:     order.Limit,
		Quantity:  order.Quantity,
	}
	m.addOpenOrder(openOrder)

	// Save Order to the database
	orderModel := convertToOrderModel(openOrder)
	err = m.orderStore.Save(ctx, orderModel)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) setBalance(marketKey string, total decimal.Decimal, available decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)] = &Balance{
		Total:     total,
		Available: available,
	}
}

func (m *Manager) setAvailableBalance(marketKey string, available decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)].Available = available
}

func (m *Manager) getBalance(marketKey string) decimal.Decimal {
	return m.Balances[utils.Normalize(marketKey)].Available
}

func (m *Manager) updateBalanceFromOrder(orderType string, order *LimitOrder) error {
	if orderType == BuyOrder {
		baseCurrencyBalance := m.getBalance(order.BaseCurrency)
		orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := baseCurrencyBalance.Sub(orderCost)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.BaseCurrency, newBalance)
		return nil
	}

	if orderType == SellOrder {

		marketCurrencyBalance := m.getBalance(order.MarketCurrency)
		orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := marketCurrencyBalance.Sub(orderCost)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.MarketCurrency, newBalance)
		return nil
	}

	return errors.New("invalid order type")
}

func (m *Manager) addOpenOrder(openOrder *OpenOrder) {
	m.OpenOrders = append(m.OpenOrders, openOrder)
}
