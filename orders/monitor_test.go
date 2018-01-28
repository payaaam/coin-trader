package orders

import (
	"testing"
)

func TestExecuteBuy(t *testing.T) {
	mockConfig := newMockDependencies(t)
	monitor := NewMonitor(mockConfig.Exchange, mockConfig.OrderUpdateChannel, 1)
	monitor.Start()

	// Check calls for client
	// Check orders array
}

func TestExecuteBuyError(t *testing.T) {

}

func TestExecuteSell(t *testing.T) {

}

func TestExecuteSellError(t *testing.T) {

}

func TestOrderUpdateFilled(t *testing.T) {

}

func TestOrderUpdatePartiallyFilled(t *testing.T) {

}

func TestOrderTimeout(t *testing.T) {

}
