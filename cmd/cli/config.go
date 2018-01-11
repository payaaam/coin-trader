package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	PostgresConn string
	Bittrex      *BittrexConfig
	Binance      *BinanceConfig
	LogLevel     log.Level
}

type BittrexConfig struct {
	ApiKey    string
	ApiSecret string
}

type BinanceConfig struct {
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

	binanceConfig := loadBinanceConfig()

	return &Config{
		PostgresConn: postgresConn,
		Bittrex:      bittrexConfig,
		Binance:			binanceConfig,
		LogLevel:     logLevel,
	}
}


// @TODO: Make Generic
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

func loadBinanceConfig() *BinanceConfig {
	apiKey := os.Getenv("BINANCE_API_KEY")
	apiSecret := os.Getenv("BINANCE_API_SECRET")

	if apiKey != "" && apiSecret != "" {
		return &BinanceConfig{
			ApiKey:    apiKey,
			ApiSecret: apiSecret,
		}
	}
	return nil
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

	return fmt.Sprintf("dbname=%s user=%s host=%s port=%s sslmode=disable", postgresDBName, postgresUser, postgresHost, postgresPort)
}
