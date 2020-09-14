package worker

import (
	"fmt"
	"time"
)

const (
	layout	= "03PM"
)

func parseTime(t string) (*time.Time, error) {
	parsedTime, err := time.Parse(layout, t)
	if err != nil {
		return nil, fmt.Errorf("parsedTime: invalid time (%s) passed: %w", t, err)
	}
	return &parsedTime, nil
}