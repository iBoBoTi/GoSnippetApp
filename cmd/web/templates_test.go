package main

import (
	"testing"
	"time"
)

func Test_humanDate(t *testing.T) {
	tm := time.Date(2020, 12, 17, 10, 0, 0, 0, time.UTC)
	got := humanDate(tm)
	if got != tm.Format(time.RFC1123) {
		t.Errorf("want %q, got %q", tm.Format(time.RFC1123), got)
	}
}
