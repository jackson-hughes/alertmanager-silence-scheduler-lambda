package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Status struct {
	State string `json:"state"`
}

type Matcher struct {
	IsRegex bool   `json:"isRegex"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

type AlertmanagerSilence struct {
	Id        string    `json:"id"`
	Status    Status    `json:"status"`
	Comment   string    `json:"comment"`
	CreatedBy string    `json:"createdBy"`
	StartsAt  time.Time `json:"startsAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EndsAt    time.Time `json:"endsAt"`
	Matchers  []Matcher `json:"matchers"`
}

func putSilence(s AlertmanagerSilence) error {
	apiUrl := alertmanagerBaseUrl + silencesApiUrl // dupe clean up later
	b, err := json.MarshalIndent(s, "", "    ")
	log.Debugf("posting new silence to alert manager:\n", string(b))
	resp, err := http.Post("http://"+apiUrl, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Error(err)
	}
	log.Info(resp)
	return nil
}

func getSilences() ([]AlertmanagerSilence, error) {
	apiUrl := alertmanagerBaseUrl + silencesApiUrl

	var allSilences []AlertmanagerSilence // existing silences includes all states (e.g. expired)

	resp, err := http.Get("http://" + apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &allSilences)
	if err != nil {
		return nil, err
	}
	return getActiveSilences(allSilences)
}

func getActiveSilences(silences []AlertmanagerSilence) ([]AlertmanagerSilence, error) {
	var activeSilences []AlertmanagerSilence
	for _, s := range silences {
		if s.Status.State == "active" || s.Status.State == "pending" {
			activeSilences = append(activeSilences, s)
		}
	}
	if len(activeSilences) > 0 {
		log.Debug("Found active silences in Alert Manager:")
		sPretty, err := json.MarshalIndent(activeSilences, "", "    ")
		if err != nil {
			return nil, err
		}
		log.Debug(string(sPretty))
	} else {
		log.Debug("There are no active silences")
	}
	return activeSilences, nil
}

func handleRecurringSilence() {}
