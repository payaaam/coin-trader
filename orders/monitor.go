package orders

import (
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"time"
)

type Monitor struct {
	openOrders   []*OpenOrder
	ticker       *time.Ticker
	client       exchanges.Exchange
	orderUpdates chan *OpenOrder
	// TODO ADD error channel
}

func NewMonitor(client exchanges.Exchange, tickerIntervalSeconds int) OrderMonitor {
	return &Monitor{
		openOrders: []*OpenOrder{},
		ticker:     time.NewTicker(time.Second * time.Duration(tickerIntervalSeconds)),
		client:     client,
	}
}

func NewTestMonitor(client exchanges.Exchange) OrderMonitor {
	return &Monitor{
		openOrders: []*OpenOrder{},
		client:     client,
	}
}

// Executes a buy or sell order
func (m *Monitor) Execute(order *OpenOrder) (string, error) {
	var orderID string
	var err error
	switch order.Type {
	case BuyOrder:
		orderID, err = m.client.ExecuteLimitBuy(order.MarketKey, order.Limit, order.Quantity)
	case SellOrder:
		orderID, err = m.client.ExecuteLimitSell(order.MarketKey, order.Limit, order.Quantity)
	default:
		return "", ErrInvalidOrderType
	}

	// Check buy and sell errors
	if err != nil {
		return "", errors.Wrap(err, ErrExecuteFailure.Error())
	}

	order.ID = orderID
	m.addOrder(order)
	return orderID, nil
}

// Starts listening for open orders
func (m *Monitor) Start(orderChannel chan *OpenOrder) {
	m.orderUpdates = orderChannel
	if m.ticker == nil {
		return
	}
	for _ = range m.ticker.C {
		if len(m.openOrders) == 0 {
			continue
		}
		m.process()
	}
}

func (m *Monitor) process() {
	var newOpenOrders []*OpenOrder
	for index, _ := range m.openOrders {
		order := m.openOrders[index]
		// Remove any items that have been filled
		if order.Status == FilledOrderStatus {
			continue
		}
		newOpenOrders = append(newOpenOrders, order)

		// Fetch Latest Order Information
		exchangeOrder, err := m.client.GetOrder(order.ID)
		if err != nil {
			log.Error(err)
			continue
		}

		// Update fields in order
		updateOrder(order, exchangeOrder)

		log.Info("HERE")

		// Determine if the order has timed out
		if order.hasReachedTimeout() && order.Status == OpenOrderStatus {
			log.Info("HERE TOO")
			err := m.cancelOrder(order)
			if err != nil {
				log.Error(err)
				continue
			}
		}

		// Send updated order to channel for Manager
		m.orderUpdates <- order
	}

	m.openOrders = newOpenOrders
}

func (m *Monitor) GetOrders() []*OpenOrder {
	return m.openOrders
}

func (m *Monitor) addOrder(order *OpenOrder) {
	m.openOrders = append(m.openOrders, order)
}

func updateOrder(order *OpenOrder, exchangeOrder *exchanges.Order) {
	order.QuantityFilled = exchangeOrder.QuantityFilled
	order.TradePrice = exchangeOrder.TradePrice
	if exchangeOrder.CloseTimestamp != 0 {
		order.Status = FilledOrderStatus
		order.CloseTimestamp = exchangeOrder.CloseTimestamp
	}
}

func (m *Monitor) cancelOrder(order *OpenOrder) error {
	err := m.client.CancelOrder(order.ID)
	if err != nil {
		return err
	}

	// Slight Delay for their system to update
	time.Sleep(time.Millisecond * 10)

	exchangeOrder, err := m.client.GetOrder(order.ID)
	if err != nil {
		return err
	}

	updateOrder(order, exchangeOrder)
	return nil
}
