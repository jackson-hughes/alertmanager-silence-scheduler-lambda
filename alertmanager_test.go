package main

import (
	"reflect"
	"testing"
)

func TestFilterAlertManagerSilences(t *testing.T) {
	testSilences := []alertmanagerSilence{
		{ID: "1", Status: status{"pending"}},
		{ID: "2", Status: status{"active"}},
		{ID: "3", Status: status{"expired"}},
	}

	want := []alertmanagerSilence{
		{ID: "1", Status: status{"pending"}},
		{ID: "2", Status: status{"active"}},
	}

	got, err := filterAlertManagerSilences(testSilences, "active", "pending")
	if err != nil {
		t.Error("error getting active silences: ", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v got %v", got, want)
	}
}
