package main

import (
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

// takes matcher and returns matcher with sorted key values
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

// compares dynamo db records to alert manager records
func compare(a []AlertmanagerSilence, d []Record) ([]AlertmanagerSilence, error) {
	newSilences := []AlertmanagerSilence{}

	// for each silence in dynamo collection
	// check if it's in alert manager
	// if not - add to newSilences list - otherwise, continue

	for _, dv := range d {
		found := false
		for _, v := range a {
			if matchersCompare(dv.Matchers, v.Matchers) {
				log.Debugf("silence for: %+v found in Alert Manager... continuing", dv.Matchers)
				found = true
			}
		}
		if !found {
			log.Debug("Didn't find existing silence - creating a new silence for: ", dv.Matchers)
			newSilences = append(newSilences, AlertmanagerSilence{
				Comment:   "Silencing for regular maintenance window",
				CreatedBy: "MSO Automated Tooling",
				StartsAt:  dv.StartsAt,
				EndsAt:    dv.EndsAt,
				Matchers:  dv.Matchers,
			})
		}
	}
	return newSilences, nil
}

// true means matchers are equal - false means they are not
func matchersCompare(a, d []Matcher) bool {

	if len(a) != len(d) {
		return false
	}

	log.Trace("Alertmanager matcher prior to sort:\n", a)
	sortMatchers(a)
	log.Trace("Alertmanager matcher after sort:\n", a)

	log.Trace("DynamoDB matcher prior to sort:\n", a)
	sortMatchers(d)
	log.Trace("DynamoDB matcher after sort:\n", a)

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
