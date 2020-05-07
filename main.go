package main

import (
	"sort"
	"strings"

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

	for _, v := range s {
		if err := putSilence(v); err != nil {
			log.Error(err)
		}
	}
}

func compare(a []AlertmanagerSilence, d []Record) ([]AlertmanagerSilence, error) {
	// do I need to sort these collections before comparing the names? worth testing
	newSilences := []AlertmanagerSilence{}

	// for each silence in dynamo collection
	//	 check if it's in alert manager
	//	 if not - post it - otherwise, continue

	for _, dv := range d {
		found := false
		for _, v := range a {
			if matchersCompare(dv.Matchers, v.Matchers) {
				log.Info("Alertmanager and DynamoDB are synchronised. Nothing to do.")
				found = true
			}
		}
		if !found {
			log.Debug("Didn't find existing silence - creating a new silence for: ", dv.Matchers)
			s := AlertmanagerSilence{
				Comment:   "Silencing for regular maintenance window",
				CreatedBy: "MSO Automated Tooling",
				StartsAt:  dv.StartsAt,
				EndsAt:    dv.EndsAt,
				Matchers:  dv.Matchers,
			}
			newSilences = append(newSilences, s)
		}
	}

	return newSilences, nil
}

// takes matcher and returns it matcher with sorted key values
func sortMatchers(m []Matcher) {
	sort.SliceStable(m,
		func(i, j int) bool {
			result := strings.Compare(m[i].Name, m[j].Name)
			if result == -1 {
				return false
			}
			return true
		})
}

// true means matchers are equal - false means they are not
func matchersCompare(a, d []Matcher) bool {

	if len(a) != len(d) {
		return false
	}

	sortMatchers(a)
	sortMatchers(d)

	for i, _ := range a {
		if a[i].IsRegex != d[i].IsRegex {
			return false
		}
		if a[i].Name != d[i].Name {
			return false
		}
		if a[i].Value != d[i].Value {
			return false
		}
	}
	return true
}
