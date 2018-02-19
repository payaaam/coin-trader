package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const BalanceFile = "balance.json"

type Config struct {
	PostgresConn string
	Bittrex      *BittrexConfig
	LogLevel     log.Level
	SlackToken   string
	BalancePath  string
}

type BittrexConfig struct {
	ApiKey    string
	ApiSecret string
}

func NewConfig() *Config {
	logLevel := log.InfoLevel
	if os.Getenv("DEBUG") == "true" {
		logLevel = log.DebugLevel
	}

	// Setup database
	postgresConn := loadPostgres()

	// Setup Bittrex Client
	bittrexConfig := loadBittrexConfig()

	// Slack
	slackToken := os.Getenv("SLACK_TOKEN")

	// Balance Path
	balancePath := loadBalancePath()

	return &Config{
		PostgresConn: postgresConn,
		Bittrex:      bittrexConfig,
		LogLevel:     logLevel,
		SlackToken:   slackToken,
		BalancePath:  balancePath,
	}
}

func loadBittrexConfig() *BittrexConfig {
	apiKey := os.Getenv("BITTREX_API_KEY")
	apiSecret := os.Getenv("BITTREX_API_SECRET")

	if apiKey != "" && apiSecret != "" {
		return &BittrexConfig{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		}
	}
	return nil
}

func loadBalancePath() string {
	balancePath := os.Getenv("BALANCE_PATH")
	if balancePath == "" {
		balancePath = "."
	}

	return fmt.Sprintf("%s/%s", balancePath, BalanceFile)
}
func loadPostgres() string {
	postgresConn := os.Getenv("POSTGRES_CONN")
	if postgresConn != "" {
		return postgresConn
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		postgresHost = "localhost"
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		postgresPort = "5432"
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		postgresUser = "postgres"
	}

	postgresDBName := os.Getenv("POSTGRES_DB")
	if postgresDBName == "" {
		postgresDBName = "coins"
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")

	return fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable", postgresDBName, postgresUser, postgresPassword, postgresHost, postgresPort)
}
