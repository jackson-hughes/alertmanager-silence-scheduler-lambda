package main

import (
	log "github.com/sirupsen/logrus"
)

var (
	alertmanagerBaseUrl string = "localhost:9093"
	silencesApiUrl      string = "/api/v2/silences/"
)

func lambdaHandler() {}

func main() {
	alertManagerSilences, err := getSilences()
	if err != nil {
		log.Error(err)
	}

	scheduledSilences, err := getScheduledSilences()
	if err != nil {
		log.Error(err)
	}
	if scheduledSilences == nil {
		log.Info("No scheduled silences found in database table")
		return
	}
	s, err := compare(alertManagerSilences, scheduledSilences)
	if err != nil {
		log.Error(err)
	}
	if len(s) == 0 {
		log.Info("no new silences to be added to alert manager")
		return
	} else {
		log.Infof("New silences to be added to alert manager: %+v", s)
	}

	for _, v := range s {
		if err := putSilence(v); err != nil {
			log.Error(err)
		}
	}
}
