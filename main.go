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
	if s != nil {
		log.Infof("New silences to be added to alert manager: %+v", s)
	}
}

func compare(a []AlertmanagerSilence, d []Record) ([]AlertmanagerSilence, error) {
	// do I need to sort these collections before comparing the names? worth testing
	newSilences := []AlertmanagerSilence{}

	// for each silence in dynamo collection
	// check if it's in alert manager
	// if not - post it - otherwise, continue

	for di, dv := range d {
		for ai, av := range a {
			if dv.Matchers[di].Name != av.Matchers[ai].Name {
				newSilence := AlertmanagerSilence{
					CreatedBy: "automated-tooling",
					Comment:   "silencing for regular maintenance",
					StartsAt:  dv.StartsAt,
					EndsAt:    dv.EndsAt,
					Matchers:  dv.Matchers,
				}
				newSilences = append(newSilences, newSilence)
			} else {
				if dv.Matchers[di].Value != av.Matchers[ai].Value {
					newSilence := AlertmanagerSilence{
						CreatedBy: "automated-tooling",
						Comment:   "silencing for regular maintenance",
						StartsAt:  dv.StartsAt,
						EndsAt:    dv.EndsAt,
						Matchers:  dv.Matchers,
					}
					newSilences = append(newSilences, newSilence)
				} else {
					continue
				}
			}
		}
	}
	return newSilences, nil
}
