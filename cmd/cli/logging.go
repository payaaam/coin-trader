package main

import (
	"github.com/sirupsen/logrus"
)

func initLogging(level logrus.Level, useJson bool) {
	logrus.SetLevel(level)
	if useJson == true {
		logrus.SetFormatter(&logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyMsg:  "message",
				logrus.FieldKeyTime: "timestamp",
			},
		})
	}
}

func logInfo(marketKey string, interval string, message string) {
	logrus.WithFields(logrus.Fields{
		"component":   "ticker",
		"tradingPair": marketKey,
		"interval":    interval,
		"exchange":    "bittrex",
	}).Info(message)
}

func logError(marketKey string, interval string, err error) {
	logrus.WithFields(logrus.Fields{
		"component":   "ticker",
		"tradingPair": marketKey,
		"interval":    interval,
	}).Error(err)
}

func logFatal(err error) {
	logrus.WithFields(logrus.Fields{
		"component": "ticker",
	}).Fatal(err)
}
