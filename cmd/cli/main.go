package main

import (
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
				SetupExchange(exchange, interval)
				return nil
			},
		},
	}

	app.Run(os.Args)
}
