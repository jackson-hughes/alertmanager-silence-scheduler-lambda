package main

import (
	"testing"
	"time"
)

func TestParseCronSchedule(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		timeFormat := time.RFC3339
		cronExpression := "0 10 * * 0" // 10am on Sunday

		expectedNextCronTriggerTime, err := time.Parse(timeFormat, "2020-08-09T10:00:00Z")
		if err != nil {
			t.Error("error parsing time type from expectedNextCronTriggerTime: ", expectedNextCronTriggerTime)
		}

		startTime, err := time.Parse(timeFormat, "2020-08-06T12:06:17Z")
		if err != nil {
			t.Error("error parsing start time: ", err)
		}

		nextCronTriggerTime, err := parseCronSchedule(cronExpression, startTime.UTC())
		if err != nil {
			t.Error("error parsing cron schedule: ", err)
		}

		if nextCronTriggerTime != expectedNextCronTriggerTime {
			t.Errorf("expected %v but got %v", expectedNextCronTriggerTime, nextCronTriggerTime)
		}
	})

	t.Run("malformed cron expression", func(t *testing.T) {
		timeFormat := time.RFC3339
		badCronExpression := "0 10 * * 7" // Sunday is 0 - not 7

		startTime, err := time.Parse(timeFormat, "2020-08-06T12:06:17Z")
		if err != nil {
			t.Error("error parsing start time: ", err)
		}

		_, err = parseCronSchedule(badCronExpression, startTime)
		if err == nil {
			t.Errorf("wanted error but got %v", err)
		}
	})
}
