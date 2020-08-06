package main

import (
	"time"

	"github.com/robfig/cron"
)

func parseCronSchedule(c string, startTime time.Time) (time.Time, error) {
	s, err := cron.ParseStandard(c)
	if err != nil {
		return time.Time{}, err
	}
	return s.Next(startTime), nil // this provides when silence start
}
