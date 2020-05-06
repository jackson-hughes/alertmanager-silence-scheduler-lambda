package main

import (
	log "github.com/sirupsen/logrus"
)

var (
	alertmanagerBaseUrl string = "localhost:9093"
	silencesApiUrl      string = "/api/v2/silences/"
)

func main() {
	alertManagerSilences, err := getSilences()
	if err != nil {
		log.Error(err)
	}
	log.Info(alertManagerSilences)

	scheduledSilences, err := getScheduledSilences()
	if err != nil {
		log.Error(err)
	}
	if scheduledSilences == nil {
		log.Info("No scheduled silences found in database table")
		return
	}

}
