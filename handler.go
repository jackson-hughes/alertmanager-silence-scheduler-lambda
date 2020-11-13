package main

import (
	"context"
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"

	log "github.com/sirupsen/logrus"
)

func getSilencesFromInputEvent(jsonSilencesInput []byte) ([]scheduledSilence, error) {
	silences := []scheduledSilence{}
	if err := json.Unmarshal(jsonSilencesInput, &silences); err != nil {
		return silences, err
	}
	return silences, nil
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) {
	// get user submitted silences from event input
	scheduledSilences, err := getSilencesFromInputEvent(event.Detail)
	log.Debug("scheduled silences: ", string(event.Detail))
	if err != nil {
		log.Error(err)
	}
	if scheduledSilences == nil {
		log.Info("no scheduled silences from event")
		return
	}

	// parse cron schedules into next invocation times and append to struct field
	for i := 0; i < len(scheduledSilences); i++ {
		k := &scheduledSilences[i]
		startTime, err := parseCronSchedule(k.StartScheduleCron, time.Now().UTC())
		if err != nil {
			log.Error(err)
		}
		k.StartsAt = startTime
		endTime, err := parseCronSchedule(k.EndScheduleCron, time.Now().UTC())
		if err != nil {
			log.Error(err)
		}
		k.EndsAt = endTime
	}

	// get existing silences from alertmanager
	alertManagerSilences, err := getAlertManagerSilences(alertManagerURL)
	if err != nil {
		log.Error(err)
	}

	// compare event silences and existing silences
	s := compareSilences(alertManagerSilences, scheduledSilences)
	if len(s) == 0 {
		log.Info("no new silences to be added to alert manager")
		return
	}
	log.Debugf("New silences to be added to alert manager: %+v", s)

	// post any new silences to alertmanager
	for _, v := range s {
		if err := putAlertManagerSilence(alertManagerURL, v); err != nil {
			log.Error(err)
		}
	}
}
