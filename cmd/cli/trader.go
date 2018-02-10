package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/strategies"
	"github.com/shopspring/decimal"
	"io/ioutil"
	"strings"
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

const SimulationBalanceFile = "./balance.json"
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
		bMap, err := loadBalancesFromFile()
		if err != nil {
			log.Fatal(err)
		}
		// Setup Order Simulation with Balances
		t.orderManager.SetupSimulation(bMap)

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
				log.Infof("BTC Balance: %v", btcBalance)
				log.Infof("%s Balance: %v", strings.ToUpper(market.MarketCurrency), altBalance)

				// Generate Chart
				chart, err := t.getChart(ctx, market.MarketKey, exchange, interval)
				if err != nil {
					log.Error(err)
					continue
				}

				// Sell
				if hasBalance(altBalance) {
					if ichimokuCloudStrategy.ShouldSell(chart) == true {
						log.Infof("Executed Sell: %s to %s", market.MarketCurrency, market.BaseCurrency)
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

						log.Infof("Quantity: %v", altBalance)
						log.Infof("Price: %v", limit)
						err = t.orderManager.ExecuteLimitSell(ctx, newSellOrder)
						if err != nil {
							logError(market.MarketKey, interval, err)
						}
						log.Infof("%s Balance: %v", strings.ToUpper(market.MarketCurrency), altBalance)
					}

					continue
				}

				if hasBalance(btcBalance) && hasZeroBalance(altBalance) {
					if ichimokuCloudStrategy.ShouldBuy(chart) == true {

						log.Infof("Executed Buy: %s to %s", market.BaseCurrency, market.MarketCurrency)

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
						log.Infof("Quantity: %v", quantity)
						log.Infof("Price: %v", limit)
						err = t.orderManager.ExecuteLimitBuy(ctx, newBuyOrder)
						if err != nil {
							logError(market.MarketKey, interval, err)
						}
						log.Infof("%s Balance: %v", strings.ToUpper(market.MarketCurrency), altBalance)
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
		balanceMap := t.orderManager.GetBalances()
		err := writeBalancesToFile(balanceMap)
		if err != nil {
			log.Errorf("Error writing balance file: %v", err)
		}

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

func loadBalancesFromFile() (map[string]*orders.Balance, error) {
	var balanceMap map[string]*orders.Balance

	balanceData, err := ioutil.ReadFile(SimulationBalanceFile)
	if err != nil {
		if os.IsNotExist(err) == false {
			log.Error(err)
			return nil, err
		}

		// Create new balance file
		bMap := defaultBalance()
		err = writeBalancesToFile(bMap)
		if err != nil {
			return nil, err
		}
		return bMap, nil
	}

	err = json.Unmarshal(balanceData, &balanceMap)
	if err != nil {
		return nil, err
	}

	return balanceMap, nil
}

func defaultBalance() map[string]*orders.Balance {
	balanceMap := make(map[string]*orders.Balance)
	balanceMap["btc"] = &orders.Balance{
		Available: utils.StringToDecimal("1"),
		Total:     utils.StringToDecimal("1"),
	}
	return balanceMap
}

func writeBalancesToFile(bMap map[string]*orders.Balance) error {
	data, err := json.Marshal(bMap)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(SimulationBalanceFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
