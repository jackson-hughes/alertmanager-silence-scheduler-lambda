package main

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestGetActiveSilences(t *testing.T) {
	testSilencesTwo := `[{
      "id": "143df16d-4b0a-4b16-ac23-5b695e2b72a8",
      "matchers": [
        {
          "name": "alertname",
          "value": "mytestalert",
          "isRegex": false
        }
      ],
      "startsAt": "2020-11-12T10:46:55.0262225Z",
      "endsAt": "2020-11-12T12:46:43.106Z",
      "updatedAt": "2020-11-12T10:46:55.0262225Z",
      "createdBy": "Jackson Hughes",
      "comment": "this is a test alert",
      "status": {
        "state": "active"
      }
    },
    {
      "id": "31782406-28c0-44df-9fb1-105044134d6b",
      "matchers": [
        {
          "name": "alertname",
          "value": "myFakeAlert2",
          "isRegex": false
        }
      ],
      "startsAt": "2020-11-12T10:52:41.6230511Z",
      "endsAt": "2020-11-12T12:52:21.98Z",
      "updatedAt": "2020-11-12T10:52:41.6230511Z",
      "createdBy": "Jackson Hughes",
      "comment": "Comment",
      "status": {
        "state": "active"
      }
    }]`
	log.Trace(testSilencesTwo)

	testSilences := []alertmanagerSilence{
		{ID: "1", Status: status{"pending"}},
		{ID: "2", Status: status{"active"}},
		{ID: "3", Status: status{"expired"}},
	}

	want := []alertmanagerSilence{
		{ID: "1", Status: status{"pending"}},
		{ID: "2", Status: status{"active"}},
	}

	got, err := getActiveSilences(testSilences)
	if err != nil {
		t.Error("error getting active silences: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", got, want)
	}
}
