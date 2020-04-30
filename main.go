package main

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

var (
	alertmanagerBaseUrl string = "localhost:9093"
	silencesApiUrl      string = "/api/v2/silences/"
)

func main() {
	log.Debug(alertmanagerBaseUrl + silencesApiUrl)

	silences, err := getSilences()
	if err != nil {
		log.Error(err)
	}

	if silences != nil {
		silencesJsonPretty, err := json.MarshalIndent(silences, "", "    ")
		if err != nil {
			log.Error(err)
		}
		log.Info(string(silencesJsonPretty))
	}

	// lambdaHandler()
}

func lambdaHandler() {}
