# Coin Trader

# Setup

**Install Go**

```bash
$ brew update
$ brew install golang
```

**Update your .bashrc / .bash_profile file**

Add these lines to your `.bashrc` file.

```bash
export GOPATH=$HOME/{{ GO_WORKSPACE_DIR }} # Change this to your go workspace.
export GOROOT=/usr/local/bin/go # Location of your go installation
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
```

**Source your updated bash profile**

```bash
$ source ~/.bash_profile
```

**Install Postgres**

1. Download [Postgres](https://postgresapp.com/)
2. Download [Postico](https://eggerapps.at/postico/)
3. Open Postgres App and `Initialize`
4. Open Postico and connect with these settings...

```
Host: localhost
Password: _empty_
Port: 5432
Database: postgres
```

**Running Tests**

```bash
$ make test
```

## Database

**Download Mirgation Tool**
`$ go get github.com/rubenv/sql-migrate/...`

**Upload Schema to Database**
`$ make setup-db`

## Bittrex

**Create an API Key**

1. Create `./bittrex.env`
2. Add the following lines...
```bash
export BITTREX_API_KEY="test"
export BITTREX_API_SECRET="test"
```
3. Source the environment variables
```bash
$ source bittrex.env
```


Navigate to [Bittrex Settings](https://bittrex.com/Manage#sectionApi)

## Build

**Download all the dependencies**
`$ govendor sync`

**Build Executable**
`$ make cli`


**Run Setup**

`$ bin/cli setup`

**Run Trading Bot**

`$ bin/cli trader`



## Backtesting

Backtesting involves running a strategy over existing data to determine if it will be profitable. To backtest, you need to create a new strategy in the strategies folder. It must abide by the `Strategy interface{}` and have `ShouldBuy()` and `ShouldSell()`.

First, update the `GetStrategy` function in `cmd/cli/backtest.go` to point to your new strategy. Then build and run.

```bash
# Build CLI
$ make cli

# Run Backtest with 1h chart and btc-eth market
$ bin/cli backtest --interval=1h --marketKey=btc-eth
```

You will see a summary and the percent change on your investment.

## Ticker

The ticker application can will run and collect hourly ticker information for all of the markets on an exchange. 

### Bittrex

```bash
# Build CLI
$ make cli

# Run ticker
$ bin/cli ticker --exchange=bittrex
```

We collect hourly ticker information from bittrex due to API Rate Limit Restrictions. Then, every morning at 00:10 UTC, we process the last 24 hourly ticks and generate a 24H candle. This data is all written to RDS.


## Deploy

You must have AWS access and have your public key setup on the server to deploy. 

To deploy...

```bash
# If your local username is different than your server username
$ DEPLOY_USER=chris make deploy

# If your local username is the same as the server
$ make deploy
```