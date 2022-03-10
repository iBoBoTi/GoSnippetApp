package main

import (
	"testing"
	"time"
)

func Test_humanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{"UTC", time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC), "Thu, 17 Dec 2020 10:00:00 UTC"},
		{"Empty", time.Time{}, ""},
		{"CET", time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)), "Thu, 17 Dec 2020 10:00:00 CET"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := humanDate(tt.tm)
			if got != tt.want {
				t.Errorf("want %q, got %q", tt.want, got)
			}
		})
	}
}
