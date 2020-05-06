package main

import (
	"time"

	"github.com/robfig/cron"
)

// testing cron string: 30 * * * *
func parseCronSchedule(c string) (time.Time, error) {
	// c = "0 4 * * 0" // test string
	s, err := cron.ParseStandard(c)
	if err != nil {
		return time.Time{}, err
	}
	return s.Next(time.Now().UTC()), nil // this provides when silence start
}
