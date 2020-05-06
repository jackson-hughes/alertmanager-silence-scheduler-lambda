package main

import (
	"time"

	"github.com/robfig/cron"
)

// testing cron string: 30 * * * *
func parseCronSchedule() (time.Time, error) {
	testString := "30 * * * *"
	s, err := cron.ParseStandard(testString)
	if err != nil {
		return time.Time{}, err
	}
	return s.Next(time.Now().UTC()), nil // this provides when silence start
}
