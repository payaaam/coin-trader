package main

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/payaaam/coin-trader/db"
	"github.com/payaaam/coin-trader/exchanges"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
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

	app.Commands = []cli.Command{
		{
			Name:  "ticker",
			Usage: "Initalizes the database",
			Flags: []cli.Flag{exchangeFlag},
			Action: func(c *cli.Context) error {
				if exchanges.ValidExchanges[exchange] != true {
					log.Error(errors.New("Not a valid exchange"))
				}

				if db.ValidIntervals[interval] != true {
					log.Error(errors.New("Not a valid exchange interval"))
				}

				config := NewConfig()
				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					panic(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)

				tickerCommand := NewTickerCommand(config, marketStore, chartStore, tickStore)
				tickerCommand.Run(exchange)
				return nil
			},
		},
		{
			Name:  "trader",
			Usage: "makes you money $$$$",
			Flags: []cli.Flag{exchangeFlag, intervalFlag},
			Action: func(c *cli.Context) error {
				if exchanges.ValidExchanges[exchange] != true {
					log.Error(errors.New("Not a valid exchange"))
				}

				if db.ValidIntervals[interval] != true {
					log.Error(errors.New("Not a valid exchange interval"))
				}

				config := NewConfig()
				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					panic(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)

				traderCommand := NewTraderCommand(config, marketStore, chartStore, tickStore)
				traderCommand.Run(exchange, interval)
				return nil
			},
		},
		{
			Name:  "backtest",
			Usage: "test a strat on all the data",
			Flags: []cli.Flag{exchangeFlag, intervalFlag, marketKeyFlag},
			Action: func(c *cli.Context) error {
				if exchanges.ValidExchanges[exchange] != true {
					log.Error(errors.New("Not a valid exchange"))
				}

				if db.ValidIntervals[interval] != true {
					log.Error(errors.New("Not a valid exchange interval"))
				}

				config := NewConfig()
				postgres, err := sql.Open("postgres", config.PostgresConn)
				if err != nil {
					panic(err)
				}

				marketStore := db.NewMarketStore(postgres)
				chartStore := db.NewChartStore(postgres)
				tickStore := db.NewTickStore(postgres)

				backTestCommand := NewBackTestCommand(config, marketStore, chartStore, tickStore)
				backTestCommand.Run(exchange, interval, marketKey)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
