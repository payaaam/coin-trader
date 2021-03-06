package orders

import (
	"github.com/payaaam/coin-trader/db"
	"github.com/satori/go.uuid"
	//"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"time"
)

type Manager struct {
	Balances     map[string]*Balance
	monitor      OrderMonitor
	client       exchanges.Exchange
	orderStore   db.OrderStoreInterface
	marketStore  db.MarketStoreInterface
	orderUpdates chan *OpenOrder
	isSimulation bool
}

func NewManager(monitor OrderMonitor, orderUpdates chan *OpenOrder, client exchanges.Exchange, os db.OrderStoreInterface, ms db.MarketStoreInterface) OrderManager {
	manager := &Manager{
		Balances:     make(map[string]*Balance),
		monitor:      monitor,
		client:       client,
		orderStore:   os,
		marketStore:  ms,
		orderUpdates: orderUpdates,
		isSimulation: false,
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
	go m.monitor.Start(m.orderUpdates)

	// Start listening for order updates
	go m.startOrderListener()
	return nil
}

func (m *Manager) SetupSimulation(balances map[string]*Balance) error {
	m.isSimulation = true
	for currency, balance := range balances {
		m.setBalance(currency, balance)
	}
	return nil
}

func (m *Manager) GetBalances() map[string]*Balance {
	return m.Balances
}

func (m *Manager) ExecuteLimitBuy(ctx context.Context, order *LimitOrder) error {
	log.Debug("ExecuteLimitBuy")
	balance := m.GetBalance(order.BaseCurrency).Available
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromLimitOrder(BuyOrder, order)
	if err != nil {
		return err
	}

	// Do not execute trade if this is simulation
	if m.isSimulation == true {
		return m.simulateOpenBuyOrder(ctx, order)
	}
	return m.createOpenBuyOrder(ctx, order)
}

func (m *Manager) ExecuteLimitSell(ctx context.Context, order *LimitOrder) error {
	log.Debug("ExecuteLimitSell")
	balance := m.GetBalance(order.MarketCurrency).Available
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromLimitOrder(SellOrder, order)
	if err != nil {
		return err
	}

	// Do not execute trade if this is simulation
	if m.isSimulation == true {
		return m.simulateOpenSellOrder(ctx, order)
	}
	return m.createOpenSellOrder(ctx, order)
}

// Creates and executes open buy order
func (m *Manager) createOpenBuyOrder(ctx context.Context, order *LimitOrder) error {
	log.Debug("createOpenBuyOrder")
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	// Create Open Order and add it to Open Orders
	openOrder := &OpenOrder{
		OpenTimestamp:  time.Now().Unix(),
		Type:           BuyOrder,
		Status:         OpenOrderStatus,
		BaseCurrency:   order.BaseCurrency,
		MarketCurrency: order.MarketCurrency,
		MarketKey:      marketKey,
		Limit:          order.Limit,
		Quantity:       order.Quantity,
	}

	orderID, err := m.monitor.Execute(openOrder)
	if err != nil {
		// Rollback Balance if order execution fails
		openOrder.CancelOrder()
		m.updateBalanceFromOpenOrder(openOrder)
		return err
	}

	openOrder.ID = orderID

	// Save Order to the database
	err = m.saveOpenOrder(ctx, openOrder)
	if err != nil {
		return err
	}

	return nil
}

// Simulates executing an open buy order. This will close it immediately
func (m *Manager) simulateOpenBuyOrder(ctx context.Context, order *LimitOrder) error {
	log.Debug("simulateOpenBuyOrder")
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	openOrder := &OpenOrder{
		ID:             uuid.NewV4().String(),
		OpenTimestamp:  time.Now().Unix(),
		CloseTimestamp: time.Now().Unix() + 1,
		Type:           BuyOrder,
		Status:         FilledOrderStatus,
		BaseCurrency:   order.BaseCurrency,
		MarketCurrency: order.MarketCurrency,
		MarketKey:      marketKey,
		Limit:          order.Limit,
		Quantity:       order.Quantity,
		QuantityFilled: order.Quantity,
		TradePrice:     order.Limit,
	}

	return m.processOrderUpdate(openOrder)
}

// Creates and executes an open sell order
func (m *Manager) createOpenSellOrder(ctx context.Context, order *LimitOrder) error {
	log.Debug("createOpenSellOrder")
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	// Create Open Order
	openOrder := &OpenOrder{
		OpenTimestamp:  time.Now().Unix(),
		Type:           SellOrder,
		Status:         OpenOrderStatus,
		BaseCurrency:   order.BaseCurrency,
		MarketCurrency: order.MarketCurrency,
		MarketKey:      marketKey,
		Limit:          order.Limit,
		Quantity:       order.Quantity,
	}

	orderID, err := m.monitor.Execute(openOrder)
	if err != nil {
		// Rollback Balance if order execution fails
		openOrder.CancelOrder()
		m.updateBalanceFromOpenOrder(openOrder)
		return err
	}
	openOrder.ID = orderID

	// Save Order to the database
	err = m.saveOpenOrder(ctx, openOrder)
	if err != nil {
		return err
	}

	return nil
}

// Simulates creating an open sell order. This will close it immediately
func (m *Manager) simulateOpenSellOrder(ctx context.Context, order *LimitOrder) error {
	log.Debug("simulateOpenSellOrder")
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	// Create Open Order
	openOrder := &OpenOrder{
		ID:             uuid.NewV4().String(),
		OpenTimestamp:  time.Now().Unix(),
		CloseTimestamp: time.Now().Unix() + 1,
		Type:           SellOrder,
		Status:         FilledOrderStatus,
		BaseCurrency:   order.BaseCurrency,
		MarketCurrency: order.MarketCurrency,
		MarketKey:      marketKey,
		Limit:          order.Limit,
		Quantity:       order.Quantity,
		QuantityFilled: order.Quantity,
		TradePrice:     order.Limit,
	}

	// Save Order to the database
	return m.processOrderUpdate(openOrder)
}

// This function updates the internal balance from a limit order
func (m *Manager) updateBalanceFromOpenOrder(order *OpenOrder) {
	log.Debug("updateBalanceFromOpenOrder")
	baseBalance := m.GetBalance(order.BaseCurrency)
	marketBalance := m.GetBalance(order.MarketCurrency)
	originalCost := order.Limit.Mul(order.Quantity)
	actualCost := order.TradePrice.Mul(order.QuantityFilled)

	if order.Type == BuyOrder {
		// Update Base Currency Balance (BTC)
		baseBalance.Available = baseBalance.Available.Add(originalCost)
		baseBalance.Available = baseBalance.Available.Sub(actualCost)
		baseBalance.Total = baseBalance.Total.Sub(actualCost)

		// Update Market Currency (ALT)
		marketBalance.Available = marketBalance.Available.Add(order.QuantityFilled)
		marketBalance.Total = marketBalance.Total.Add(order.QuantityFilled)
	}

	if order.Type == SellOrder {
		// Update Market Currency Balance (ALT)
		marketBalance.Available = marketBalance.Available.Add(order.Quantity)
		marketBalance.Available = marketBalance.Available.Sub(order.QuantityFilled)
		marketBalance.Total = marketBalance.Total.Sub(order.QuantityFilled)

		// Update Base Currency (BTC)
		baseBalance.Available = baseBalance.Available.Add(actualCost)
		baseBalance.Total = baseBalance.Total.Add(actualCost)
	}

	m.setBalance(order.MarketCurrency, marketBalance)
	m.setBalance(order.BaseCurrency, baseBalance)
}

// Receives updated open order and updates balances accordingly
func (m *Manager) processOrderUpdate(order *OpenOrder) error {
	log.Debug("processOrderUpdate")
	switch order.Status {
	case OpenOrderStatus:
		return nil
	case FilledOrderStatus:
		// Update manager balances
		m.updateBalanceFromOpenOrder(order)

		// Update Order in Database
		err := m.saveOpenOrder(context.Background(), order)
		if err != nil {
			return err
		}
		return nil
	case PartiallyFieldOrderStatus:
		return nil
	case CancelledOrderStatus:
		return nil
	default:
		return ErrInvalidOrderStatus
	}
}

// Updates the balance for currency after initial order placement
func (m *Manager) updateBalanceFromLimitOrder(orderType string, order *LimitOrder) error {
	log.Debug("updateBalanceFromLimitOrder")
	if orderType == BuyOrder {
		baseCurrencyBalance := m.GetBalance(order.BaseCurrency).Available
		orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := baseCurrencyBalance.Sub(orderCost)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.BaseCurrency, newBalance)
	}

	if orderType == SellOrder {
		marketCurrencyBalance := m.GetBalance(order.MarketCurrency).Available
		//orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := marketCurrencyBalance.Sub(order.Quantity)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.MarketCurrency, newBalance)
	}
	return nil
}

// Starts listening on the monitorChannel
func (m *Manager) startOrderListener() {
	for {
		updatedOrder := <-m.orderUpdates
		err := m.processOrderUpdate(updatedOrder)
		if err != nil {
			log.Error(err)
		}
	}
}

// Fetches balances from Exchnage
func (m *Manager) loadBalances() error {
	// Get your balances from exchange
	balances, err := m.client.GetBalances()
	if err != nil {
		return err
	}

	for _, balance := range balances {
		b := &Balance{
			Total:     balance.Total,
			Available: balance.Available,
		}
		m.setBalance(balance.BaseCurrency, b)
	}

	return nil
}

// Sets available and total balance
func (m *Manager) setBalance(marketKey string, balance *Balance) {
	m.Balances[utils.Normalize(marketKey)] = balance
}

// Sets available balance
func (m *Manager) setAvailableBalance(marketKey string, available decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)].Available = available
}

// Get Available Balance
func (m *Manager) GetBalance(marketKey string) *Balance {
	b := m.Balances[utils.Normalize(marketKey)]
	if b != nil {
		return b
	}
	return &Balance{
		Available: utils.ZeroDecimal(),
		Total:     utils.ZeroDecimal(),
	}
}

// Save an Order to the database
func (m *Manager) saveOpenOrder(ctx context.Context, openOrder *OpenOrder) error {
	// Save Order to the database
	orderModel := convertToOrderModel(openOrder)
	market, err := m.marketStore.GetMarket(ctx, "bittrex", openOrder.MarketKey)
	if err != nil {
		return err
	}

	orderModel.MarketID = market.ID
	err = m.orderStore.Upsert(ctx, orderModel)
	if err != nil {
		return err
	}

	return nil
}
