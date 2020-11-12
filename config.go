package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	silencesApiUrl      string = "/api/v2/silences/"
	alertmanagerBaseUrl string = "localhost:9093"
	alertManagerUrl     string = alertmanagerBaseUrl + silencesApiUrl
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

	envAlertmanagerUrl, set := os.LookupEnv("ALERTMANAGER_URL")
	if set {
		alertmanagerBaseUrl = envAlertmanagerUrl
		log.Debug("alertmanager url environment variable found: ", alertmanagerBaseUrl)
	}

	envAlertManagerSilencesApiUrl, set := os.LookupEnv("ALERTMANAGER_SILENCE_API_URL")
	if set {
		silencesApiUrl = envAlertManagerSilencesApiUrl
		log.Debug("silence api url environment variable found: ", silencesApiUrl)
	}
}
