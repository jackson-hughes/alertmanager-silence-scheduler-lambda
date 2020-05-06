package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		TimestampFormat:        "02-01-2006 15:04:05",
	})
	envLogLevel, set := os.LookupEnv("LOG_LEVEL")
	if set {
		envLogLevelString, err := log.ParseLevel(envLogLevel)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(envLogLevelString)
		log.Debug("log level environment variable found: ", envLogLevel)
	}

	envAlertmanagerUrl, set := os.LookupEnv("ALERTMANAGER_URL")
	if set {
		alertmanagerBaseUrl = envAlertmanagerUrl
		log.Debug("alertmanager url environment variable found: ", envAlertmanagerUrl)
	}
}
