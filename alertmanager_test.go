package main

import (
	"reflect"
	"testing"
)

func TestGetActiveSilences(t *testing.T) {
	testSilences := []AlertmanagerSilence{
		{
			Id: "1",
			Status: Status{
				"pending",
			},
		},
		{
			Id: "2",
			Status: Status{
				"active",
			},
		},
		{
			Id: "3",
			Status: Status{
				"expired",
			},
		},
	}

	want := []AlertmanagerSilence{
		{
			Id: "1",
			Status: Status{
				"pending",
			},
		},
		{
			Id: "2",
			Status: Status{
				"active",
			},
		},
	}

	got, err := getActiveSilences(testSilences)
	if err != nil {
		t.Error("error getting active silences: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", got, want)
	}
}
