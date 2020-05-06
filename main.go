package main

import (
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	alertmanagerBaseUrl string = "localhost:9093"
	silencesApiUrl      string = "/api/v2/silences/"
)

func main() {
	silences, err := getSilences()
	if err != nil {
		log.Error(err)
	}

	if silences != nil {
		silencesJsonPretty, err := json.MarshalIndent(silences, "", "    ")
		if err != nil {
			log.Error(err)
		}
		log.Debug(string(silencesJsonPretty))
	}
	// getScheduledSilences()
	// lambdaHandler()
	next, err := parseCronSchedule()
	if err != nil {
		log.Error(err)
	} else {
		log.Info(next.Format("2006-01-02T15:04:05Z"))
	}
	log.Debug("Time now: ", time.Now().Format("2006-01-02T15:04:05Z"))
}

func lambdaHandler() {}
