package orders

import (
	"github.com/payaaam/coin-trader/exchanges"
	"time"
)

type Monitor struct {
	openOrders []*OpenOrder
	ticker     *time.Ticker
	client     exchanges.Exchange
}

func NewMonitor(client exchanges.Exchange, tickerIntervalSeconds int) OrderMonitor {
	return &Monitor{
		openOrders: []*OpenOrder{},
		ticker:     time.NewTicker(time.Second * time.Duration(tickerIntervalSeconds)),
		client:     client,
	}
}

/*
func NewMonitor(tickerIntervalSeconds int) OrderMonitor {
	var tickerValue *time.Ticker
	if tickerIntervalSeconds != 0 {
		tickerValue = time.NewTicker(time.Second * time.Duration(tickerIntervalSeconds))
	}

	return &Monitor{
		openOrders: []*OpenOrder{},
		ticker:     tickerValue,
	}
}
*/

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
		return "", err
	}

	order.ID = orderID
	m.addOrder(order)
	return orderID, nil
}

// Starts listening for open orders
func (m *Monitor) Start(orderChannel chan *OpenOrder) {
	if m.ticker == nil {
		return
	}
	for _ = range m.ticker.C {
		if len(m.openOrders) == 0 {
			continue
		}

		m.process(orderChannel)
	}
}

func (m *Monitor) process(orderChannel chan *OpenOrder) {
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

func (m *Monitor) GetOrders() []*OpenOrder {
	return m.openOrders
}

func (m *Monitor) addOrder(order *OpenOrder) {
	m.openOrders = append(m.openOrders, order)
}
