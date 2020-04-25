package main

import (
	"testing"
	"time"
)

func TestUptime(t *testing.T) {
	startTime = time.Now()
	roundDuration, err := time.ParseDuration("1s")
	if err != nil {
		panic(err)
	}
	seconds := uptime().Round(roundDuration).Seconds()
	if seconds != 0 {
		t.Errorf("Uptime returned invalid time, got: %v, want: %v", seconds, 0)
	}
}
