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
	var interval string
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "exchange",
			Value:       "bittrex",
			Usage:       "the `exchange` you want to pull ticker data from",
			Destination: &exchange,
		},
		cli.StringFlag{
			Name:        "interval",
			Value:       "1D",
			Usage:       "the `time interval` of the tickers (1D, 4h, 1h, 30)",
			Destination: &interval,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "setup",
			Usage: "Initalizes the database",
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

				setupCommand := NewSetupCommand(config, marketStore, chartStore, tickStore)
				setupCommand.Run(exchange, interval)
				return nil
			},
		},
		{
			Name:  "trader",
			Usage: "makes you money $$$$",
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
	}

	app.Run(os.Args)
}
