#!/bin/bash

set -eu

EXCHANGE="bittrex"
INTERVAL="1h"
TICKER_LOG_PATH="ticker.log"
TRADER_LOG_PATH="trader.log"
TICKER_PID_FILE="/tmp/ticker.pid"
TRADER_PID_FILE="/tmp/trader.pid"

# Source the environment
source /etc/cli.env

# Start the ticker
./coins-cli ticker --exchange=$EXCHANGE &>> $TICKER_LOG_PATH & echo $! > $TICKER_PID_FILE

# Start the trader
./coins-cli trader --exchange=$EXCHANGE --interval=1h --simulate &>> $TRADER_LOG_PATH & echo $! > $TRADER_PID_FILE