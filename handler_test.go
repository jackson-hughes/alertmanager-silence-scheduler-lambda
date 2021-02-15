// +build integration

package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestHandleRequest(t *testing.T) {
	alertManagerFullUrl := appConfig.AlertManagerUrl + ":" + appConfig.AlertManagerTcpPort + appConfig.AlertManagerSilenceApiUrl

	// test JSON input
	testJsonInput := `[
    {
        "Service": "ExampleService",
        "StartScheduleCron": "0 3 * * 0",
        "EndScheduleCron": "0 4 * * 0",
        "Matchers": [
            {
                "IsRegex": false,
                "Name": "environment",
                "Value": "test"
            },
            {
                "IsRegex": false,
                "Name": "alertname",
                "Value": "HostDown"
            }
        ]
    },
    {
        "Service": "ExampleServiceTwo",
        "StartScheduleCron": "0 8 * * 0",
        "EndScheduleCron": "0 9 * * 0",
        "Matchers": [
            {
                "IsRegex": false,
                "Name": "environment",
                "Value": "prod"
            },
            {
                "IsRegex": false,
                "Name": "alertname",
                "Value": "LambdaError"
            }
        ]
    }
]`

	// validate JSON input
	valid := json.Valid([]byte(testJsonInput))
	if !valid {
		t.Errorf("invalid JSON: %v", testJsonInput)
	}

	// create mock Cloudwatch event and insert test JSON data
	testEvent := events.CloudWatchEvent{
		Detail: []byte(testJsonInput),
	}

	// call handler func which will post the test silences to alertmanager
	handleRequest(context.TODO(), testEvent)

	// query alertmanager for silences
	resp, err := http.Get("http://" + alertManagerFullUrl)
	if err != nil {
		t.Error(err)
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}

	// unmarshall the alert manager response
	actualAlerts := []alertmanagerSilence{}
	err = json.Unmarshal(responseBody, &actualAlerts)
	if err != nil {
		t.Error(err)
	}

	// naive check just counting the number of silences matches what we posted
	expectedAlertCount := 2
	if expectedAlertCount != len(actualAlerts) {
		t.Errorf("want %v alerts but got %v alerts", expectedAlertCount, len(actualAlerts))
	}
}
