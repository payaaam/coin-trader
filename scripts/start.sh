#!/bin/bash

set -eu

EXCHANGE="bittrex"
INTERVAL="1h"
BASE_PATH="/var/app"
TICKER_LOG_PATH="${BASE_PATH}/ticker.log"
TRADER_LOG_PATH="${BASE_PATH}/trader.log"
TICKER_PID_FILE="/tmp/ticker.pid"
TRADER_PID_FILE="/tmp/trader.pid"


# Source the environment
source /etc/cli.env

# Start the ticker
${BASE_PATH}/coins-cli ticker --exchange=$EXCHANGE &>> $TICKER_LOG_PATH & echo $! > $TICKER_PID_FILE

# Start the trader
${BASE_PATH}/coins-cli trader --exchange=$EXCHANGE --interval=1h --simulate &>> $TRADER_LOG_PATH & echo $! > $TRADER_PID_FILE