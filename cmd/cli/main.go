package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/charts"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	"github.com/payaaam/coin-trader/orders"
	"github.com/payaaam/coin-trader/slack"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "coins-cli"
	app.Version = "0.1.0"
	app.Usage = ""

	var exchange string
	exchangeFlag := cli.StringFlag{
		Name:        "exchange",
		Value:       "bittrex",
		Usage:       "the `exchange` you want to pull ticker data from",
		Destination: &exchange,
	}
	var interval string
	intervalFlag := cli.StringFlag{
		Name:        "interval",
		Value:       "1D",
		Usage:       "the `time interval` of the tickers (1D, 4h, 1h, 30)",
		Destination: &interval,
	}

	var marketKey string
	marketKeyFlag := cli.StringFlag{
		Name:        "marketKey",
		Value:       "btc-eth",
		Usage:       "the `market pair` you want to test (btc-eth, btc-gnt)",
		Destination: &marketKey,
	}

	var isSimulation bool
	simulateFlag := cli.BoolFlag{
		Name:        "simulate",
		Usage:       "prevents order execution",
		Destination: &isSimulation,
	}

	app.Commands = []cli.Command{
		{
			Name:  "ticker",
			Usage: "fetches ticker information from exchange",
			Flags: []cli.Flag{exchangeFlag},
			Action: func(c *cli.Context) error {
				if exchanges.ValidExchanges[exchange] != true {
					logFatal(fmt.Errorf("%s not a valid exchange", exchange))
				}

				config := NewConfig()
				initLogging(config.LogLevel, true)

				exchangeClient, err := getExchangeClient(config, exchange)
				if err != nil {
					log.Fatal(err)
				}

				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					panic(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)

				tickerCommand := NewTickerCommand(config, marketStore, chartStore, tickStore, exchangeClient)
				tickerCommand.Run(exchange)
				return nil
			},
		},

		{
			Name:  "trader",
			Usage: "makes you money $$$$",
			Flags: []cli.Flag{exchangeFlag, intervalFlag, simulateFlag},
			Action: func(c *cli.Context) error {
				config := NewConfig()
				initLogging(config.LogLevel, true)

				if exchanges.ValidExchanges[exchange] != true {
					logFatal(fmt.Errorf("%s not a valid exchange", exchange))
				}

				if charts.ValidIntervals[interval] != true {
					logFatal(fmt.Errorf("%s not a valid interval", interval))
				}

				exchangeClient, err := getExchangeClient(config, exchange)
				if err != nil {
					log.Fatal(err)
				}

				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					panic(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)
				orderStore := db.NewOrderStore(postgres)

				orderUpdateChannel := make(chan *orders.OpenOrder)
				orderMonitor := orders.NewMonitor(exchangeClient, 10)
				orderManager := orders.NewManager(orderMonitor, orderUpdateChannel, exchangeClient, orderStore, marketStore)
				slackLogger := slack.NewSlackLogger(config.SlackToken)

				traderCommand := NewTraderCommand(config, marketStore, chartStore, tickStore, exchangeClient, orderManager, slackLogger)

				traderCommand.Run(exchange, interval, isSimulation)
				return nil
			},
		},
		{
			Name:  "backtest",
			Usage: "test a strat on all the data",
			Flags: []cli.Flag{exchangeFlag, intervalFlag, marketKeyFlag},
			Action: func(c *cli.Context) error {
				config := NewConfig()
				initLogging(config.LogLevel, false)

				if exchanges.ValidExchanges[exchange] != true {
					logFatal(fmt.Errorf("%s not a valid exchange", exchange))
				}

				if charts.ValidIntervals[interval] != true {
					logFatal(fmt.Errorf("%s not a valid interval", interval))
				}

				exchangeClient, err := getExchangeClient(config, exchange)
				if err != nil {
					log.Fatal(err)
				}

				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					log.Fatal(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)

				backTestCommand := NewBackTestCommand(config, marketStore, chartStore, tickStore, exchangeClient)
				backTestCommand.Run(exchange, interval, marketKey)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
