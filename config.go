package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	silencesAPIURL      = "/api/v2/silences/"
	alertManagerBaseURL = "localhost:9093"
	alertManagerTcpPort = "9093"
	alertManagerURL     = alertManagerBaseURL + alertManagerTcpPort + silencesAPIURL
)

func init() {
	// logger config
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-02-01 15:04:05",
	})

	// environment variables for app configuration
	envLogLevel, set := os.LookupEnv("LOG_LEVEL")
	if set {
		envLogLevelString, err := log.ParseLevel(envLogLevel)
		if err != nil {
			log.Error(err)
		}
		log.SetLevel(envLogLevelString)
		log.Debug("log level environment variable found: ", envLogLevel)
	}

	envAlertManagerURL, set := os.LookupEnv("ALERTMANAGER_URL")
	if set {
		alertManagerBaseURL = envAlertManagerURL
		alertManagerURL = alertManagerBaseURL + alertManagerTcpPort + silencesAPIURL
		log.Debug("alertmanager url environment variable found: ", alertManagerBaseURL)
	}

	envAlertManagerSilencesAPIURL, set := os.LookupEnv("ALERTMANAGER_SILENCE_API_URL")
	if set {
		silencesAPIURL = envAlertManagerSilencesAPIURL
		log.Debug("silence api url environment variable found: ", silencesAPIURL)
	}
}
