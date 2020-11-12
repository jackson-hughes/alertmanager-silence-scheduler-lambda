package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type status struct {
	State string `json:"state"`
}

type matcher struct {
	IsRegex bool   `json:"isRegex"`
	Name    string `json:"name"`
	Value   string `json:"value"`
}

type scheduledSilence struct {
	Service           string
	StartScheduleCron string
	EndScheduleCron   string
	Matchers          []matcher
	StartsAt          time.Time
	EndsAt            time.Time
}

type alertmanagerSilence struct {
	ID        string    `json:"id"`
	Status    status    `json:"status"`
	Comment   string    `json:"comment"`
	CreatedBy string    `json:"createdBy"`
	StartsAt  time.Time `json:"startsAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	EndsAt    time.Time `json:"endsAt"`
	Matchers  []matcher `json:"matchers"`
}

// getSilences retrieves all silences from AlertManager
func getSilences(alertManagerURL string) ([]alertmanagerSilence, error) {
	var allSilences []alertmanagerSilence // existing silences includes all states (e.g. expired)

	resp, err := http.Get("http://" + alertManagerURL)
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

// getActiveSilences sorts through a slice of AlertManager silences removing any that have expired
func getActiveSilences(silences []alertmanagerSilence) ([]alertmanagerSilence, error) {
	var activeSilences []alertmanagerSilence
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

// putSilence takes an AlertManager silence and PUTs it into Alertmanager over http
func putSilence(alertManagerURL string, s alertmanagerSilence) error {
	b, err := json.MarshalIndent(s, "", "    ")

	log.Debug("posting new silence to alert manager:\n", string(b))
	resp, err := http.Post("http://"+alertManagerURL, "application/json", bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	log.Debug("alertmanager response code: ", resp.Status)
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	log.Debug("alertmanager response body: ", buf.String())
	defer resp.Body.Close()
	return nil
}
