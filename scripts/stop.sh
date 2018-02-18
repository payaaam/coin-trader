#!/bin/bash

TICKER_PID_FILE=/tmp/ticker.pid
TRADER_PID_FILE=/tmp/trader.pid

TRADER_PID=$(cat ${TRADER_PID_FILE})
TICKER_PID=$(cat ${TICKER_PID_FILE})

echo Killing trader: $TRADER_PID 
kill -INT $TRADER_PID

echo Killing ticker: $TICKER_PID 
kill -INT $TICKER_PID
exit