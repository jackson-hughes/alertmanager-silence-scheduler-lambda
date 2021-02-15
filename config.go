package main

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Specification struct {
	LogLevel                  string `default:"info"`
	AlertManagerUrl           string `default:"localhost"`
	AlertManagerTcpPort       string `default:"9093"`
	AlertManagerSilenceApiUrl string `default:"/api/v2/silences/"`
}

var (
	appConfig Specification
)

func init() {
	// logger config
	log.SetFormatter(&log.TextFormatter{
		DisableLevelTruncation: true,
		FullTimestamp:          true,
		TimestampFormat:        "2006-02-01 15:04:05",
	})

	err := envconfig.Process("amsilencer", &appConfig)
	if err != nil {
		log.Fatal("error initialising application configuration: ", err)
	}

	LogLevelType, err := log.ParseLevel(appConfig.LogLevel)
	if err != nil {
		log.Error("error parsing log level: ", err)
	}
	log.SetLevel(LogLevelType)
	log.Debug("log level setting: ", appConfig.LogLevel)
}
