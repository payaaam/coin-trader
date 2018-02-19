#!/bin/bash

# DB
mockgen -source=./db/db.go -destination=./mocks/db.go -package=mocks

# Exchanges
mockgen -source=./exchanges/exchanges.go -destination=./mocks/exchanges.go -package=mocks -imports .=github.com/payaaam/coin-trader/exchanges

# Manager
mockgen -source=./db/db.go -destination=./mocks/db.go -package=mocks

# Strategy
mockgen -source=./strategies/strategy.go -destination=./mocks/strategy.go -package=mocks

# Slack
mockgen -source=./slack/slack.go -destination=./mocks/slack.go -package=mocks

echo Generated Mocks