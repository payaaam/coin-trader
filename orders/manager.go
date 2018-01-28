package orders

import (
	"github.com/payaaam/coin-trader/db"
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
}

func NewManager(monitor OrderMonitor, orderUpdates chan *OpenOrder, client exchanges.Exchange, os db.OrderStoreInterface, ms db.MarketStoreInterface) OrderManager {
	manager := &Manager{
		Balances:     make(map[string]*Balance),
		monitor:      monitor,
		client:       client,
		orderStore:   os,
		marketStore:  ms,
		orderUpdates: orderUpdates,
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

func (m *Manager) GetBalances() map[string]*Balance {
	return m.Balances
}

func (m *Manager) ExecuteLimitBuy(ctx context.Context, order *LimitOrder) error {
	balance := m.getBalance(order.BaseCurrency).Available
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromInitialOrder(BuyOrder, order)
	if err != nil {
		return err
	}
	err = m.createOpenBuyOrder(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (m *Manager) ExecuteLimitSell(ctx context.Context, order *LimitOrder) error {
	balance := m.getBalance(order.MarketCurrency).Available
	if hasAvailableFunds(balance, order) == false {
		return ErrNotEnoughFunds
	}

	err := m.updateBalanceFromInitialOrder(SellOrder, order)
	if err != nil {
		return err
	}
	err = m.createOpenSellOrder(ctx, order)
	if err != nil {
		// If Error, process and roll back balance updates
		return err
	}

	return nil
}

func (m *Manager) createOpenBuyOrder(ctx context.Context, order *LimitOrder) error {
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	// Create Open Order and add it to Open Orders
	openOrder := &OpenOrder{
		OrderPlacedTimestamp: time.Now().Unix(),
		Type:                 BuyOrder,
		Status:               OpenOrderStatus,
		BaseCurrency:         order.BaseCurrency,
		MarketCurrency:       order.MarketCurrency,
		MarketKey:            marketKey,
		Limit:                order.Limit,
		Quantity:             order.Quantity,
	}

	orderID, err := m.monitor.Execute(openOrder)
	if err != nil {
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

func (m *Manager) createOpenSellOrder(ctx context.Context, order *LimitOrder) error {
	marketKey := m.client.GetMarketKey(order.BaseCurrency, order.MarketCurrency)

	// Create Open Order
	openOrder := &OpenOrder{
		OrderPlacedTimestamp: time.Now().Unix(),
		Type:                 SellOrder,
		Status:               OpenOrderStatus,
		BaseCurrency:         order.BaseCurrency,
		MarketCurrency:       order.MarketCurrency,
		MarketKey:            marketKey,
		Limit:                order.Limit,
		Quantity:             order.Quantity,
	}

	orderID, err := m.monitor.Execute(openOrder)
	if err != nil {
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

func (m *Manager) updateBalanceFromOpenOrder(order *OpenOrder) error {
	switch order.Status {
	case OpenOrderStatus:
		log.Info("INSIDE OPEN ORDER STATUS WOOO")
		// Check order placed time
		// If longer than timeout, cancel order
		return nil
	case FilledOrderStatus:
		if order.Type == BuyOrder {
			baseBalance := m.getBalance(order.BaseCurrency)
			marketBalance := m.getBalance(order.MarketCurrency)
			originalCost := order.Limit.Mul(order.Quantity)
			actualCost := order.TradePrice.Mul(order.Quantity)

			// Update Base Currency Balance (BTC)
			baseBalance.Available = baseBalance.Available.Add(originalCost)
			baseBalance.Available = baseBalance.Available.Sub(actualCost)
			baseBalance.Total = baseBalance.Total.Sub(actualCost)

			// Update Market Currency
			marketBalance.Available = marketBalance.Available.Add(order.Quantity)
			marketBalance.Total = marketBalance.Total.Add(order.Quantity)

			// Update Order in Database
			err := m.saveOpenOrder(context.Background(), order)
			if err != nil {
				return err
			}
		}

		// Add the limit * quantity back to Available
		// Calcualte total cost and remove from Available
		// remove from Total
		// Add Quantity to Market or CBase depending on buy or sell
		log.Info("FILLED ORDER WOOHOO")
		return nil
	case PartiallyFieldOrderStatus:
		return nil
	default:
		return ErrInvalidOrderStatus
	}
}

// Updates the balance for currency after initial order placement
func (m *Manager) updateBalanceFromInitialOrder(orderType string, order *LimitOrder) error {
	switch orderType {
	case BuyOrder:
		baseCurrencyBalance := m.getBalance(order.BaseCurrency).Available
		orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := baseCurrencyBalance.Sub(orderCost)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.BaseCurrency, newBalance)
	case SellOrder:

		marketCurrencyBalance := m.getBalance(order.MarketCurrency).Available
		orderCost := order.Limit.Mul(order.Quantity)

		// Debit Base Currency Available
		newBalance := marketCurrencyBalance.Sub(orderCost)

		if newBalance.Sign() == -1 {
			return ErrNotEnoughFunds
		}
		m.setAvailableBalance(order.MarketCurrency, newBalance)
	default:
		return ErrInvalidOrderType
	}

	return nil
}

// Starts listening on the monitorChannel
func (m *Manager) startOrderListener() {
	for {
		updatedOrder := <-m.orderUpdates
		m.updateBalanceFromOpenOrder(updatedOrder)
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
		m.setBalance(balance.BaseCurrency, balance.Total, balance.Available)
	}

	return nil
}

// Sets available and total balance
func (m *Manager) setBalance(marketKey string, total decimal.Decimal, available decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)] = &Balance{
		Total:     total,
		Available: available,
	}
}

// Sets available balance
func (m *Manager) setAvailableBalance(marketKey string, available decimal.Decimal) {
	m.Balances[utils.Normalize(marketKey)].Available = available
}

// Get Available Balance
func (m *Manager) getBalance(marketKey string) *Balance {
	return m.Balances[utils.Normalize(marketKey)]
}

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
