package main

import (
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/strategies"
	"github.com/shopspring/decimal"
	//"github.com/payaaam/coin-trader/db/models"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"os"
	"os/signal"
	//"gopkg.in/volatiletech/null.v6"
	"time"
)

const DefaultPricePadding = "1.01"
const DefaultLimit = "0.01"
const DefaultQuantity = "1"
const EveryTenMinutes = 10

type TraderCommand struct {
	config         *Config
	marketStore    db.MarketStoreInterface
	chartStore     db.ChartStoreInterface
	tickStore      db.TickStoreInterface
	exchangeClient exchanges.Exchange
	orderManager   orders.OrderManager
}

func NewTraderCommand(config *Config, marketStore db.MarketStoreInterface, chartStore db.ChartStoreInterface, tickStore db.TickStoreInterface, client exchanges.Exchange, orderManager orders.OrderManager) *TraderCommand {
	return &TraderCommand{
		config:         config,
		marketStore:    marketStore,
		chartStore:     chartStore,
		tickStore:      tickStore,
		exchangeClient: client,
		orderManager:   orderManager,
	}
}

func (t *TraderCommand) Run(exchange string, interval string, isSimulation bool) {
	log.Infof("Starting Automated Trader %s", exchange)
	ctx := context.Background()

	if isSimulation == true {
		// Setup Order Simulation with bitcoin value
		t.orderManager.SetupSimulation(&orders.Balance{
			Total:     utils.StringToDecimal("1.0"),
			Available: utils.StringToDecimal("1.0"),
		})

		log.Info("Simulation Mode")
	} else {
		log.Fatal("Must run a simulation")
		//t.orderManager.Setup()
	}

	ichimokuCloudStrategy := strategies.NewCloudStrategy()

	// Will ping
	ticker := time.NewTicker(time.Minute * EveryTenMinutes)
	go func() {
		for _ = range ticker.C {

			// Get all Markets
			markets, err := t.marketStore.GetMarkets(ctx, exchange)
			if err != nil {
				log.Error(err)
				continue
			}

			btcBalance := t.getBalance("btc")
			log.Infof("Start BTC Balance: %v", btcBalance)

			for _, market := range markets {
				// Get Balances
				btcBalance := t.getBalance("btc")
				altBalance := t.getBalance(market.MarketCurrency)

				// Generate Chart
				chart, err := t.getChart(ctx, market.MarketKey, exchange, interval)
				if err != nil {
					log.Error(err)
					continue
				}

				if hasBalance(altBalance) {
					if ichimokuCloudStrategy.ShouldSell(chart) == true {

						log.Infof("Executed Sell: %s to %s", market.MarketCurrency, market.BaseCurrency)
						return

						ticker, err := t.getLatestPrice(ctx, market.MarketKey)
						if err != nil {
							log.Error(err)
							continue
						}
						limit := getOrderPrice(orders.SellOrder, ticker)

						newSellOrder := &orders.LimitOrder{
							Limit:          limit,
							Quantity:       altBalance,
							MarketCurrency: market.MarketCurrency,
							BaseCurrency:   market.BaseCurrency,
						}

						log.Infof("Executed Sell: %s to %s", market.MarketCurrency, market.BaseCurrency)
						log.Infof("Quantity: %v", altBalance)
						log.Infof("Price: %v", limit)
						err = t.orderManager.ExecuteLimitBuy(ctx, newSellOrder)
						if err != nil {
							logError(market.MarketKey, interval, err)
						}
					}
					continue
				}

				if hasBalance(btcBalance) && hasZeroBalance(altBalance) {
					if ichimokuCloudStrategy.ShouldBuy(chart) == true {

						log.Infof("Executed Buy: %s to %s", market.BaseCurrency, market.MarketCurrency)
						return

						ticker, err := t.getLatestPrice(ctx, market.MarketKey)
						if err != nil {
							log.Error(err)
							continue
						}
						limit := getOrderPrice(orders.BuyOrder, ticker)
						quantity := getOrderQuantity(ticker, getBTCLimit())

						newBuyOrder := &orders.LimitOrder{
							Limit:          limit,
							Quantity:       quantity,
							MarketCurrency: market.MarketCurrency,
							BaseCurrency:   market.BaseCurrency,
						}
						log.Infof("Executed Buy: %s to %s", market.BaseCurrency, market.MarketCurrency)
						log.Infof("Quantity: %v", quantity)
						log.Infof("Price: %v", limit)
						err = t.orderManager.ExecuteLimitBuy(ctx, newBuyOrder)
						if err != nil {
							logError(market.MarketKey, interval, err)
						}

					}
				}
				continue
			}
			btcBalance = t.getBalance("btc")
			log.Infof("End BTC Balance: %v", btcBalance)

		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	for _ = range signalChan {
		log.Warn("received SIGINT or SIGTERM")
		break
	}
	log.Info("Shutting down")
}

func (t *TraderCommand) getChart(ctx context.Context, marketKey string, exchange string, interval string) (*charts.CloudChart, error) {
	candles, err := t.tickStore.GetChartCandles(ctx, marketKey, exchange, interval)
	if err != nil {
		return nil, err
	}

	return charts.NewCloudChartFromCandles(candles, marketKey, exchange, interval)
}

func (t *TraderCommand) getLatestPrice(ctx context.Context, marketKey string) (*exchanges.Ticker, error) {
	return t.exchangeClient.GetTicker(marketKey)
}

func (t *TraderCommand) getBalance(currency string) decimal.Decimal {
	balance := t.orderManager.GetBalance(currency)
	if balance != nil {
		return balance.Available
	}
	return utils.ZeroDecimal()
}
func hasBalance(b decimal.Decimal) bool {
	return b.Equal(utils.ZeroDecimal()) == false
}
func hasZeroBalance(b decimal.Decimal) bool {
	return b.Equal(utils.ZeroDecimal())
}

func getBTCLimit() decimal.Decimal {
	return utils.StringToDecimal(DefaultLimit)
}

func getDefaultQuantity() decimal.Decimal {
	return utils.StringToDecimal(DefaultQuantity)
}

func getOrderPrice(orderType string, ticker *exchanges.Ticker) decimal.Decimal {
	pricePadding := utils.StringToDecimal(DefaultPricePadding)
	var limit decimal.Decimal

	if orderType == orders.BuyOrder {
		limit = ticker.Ask.Mul(pricePadding)
	} else if orderType == orders.SellOrder {
		limit = ticker.Bid.Mul(pricePadding)
	}
	return limit
}
func getOrderQuantity(ticker *exchanges.Ticker, btcMax decimal.Decimal) decimal.Decimal {
	return btcMax.Div(ticker.Ask)
}
