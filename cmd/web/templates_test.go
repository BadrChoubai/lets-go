package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tm := time.Date(1999, 3, 2, 12, 0, 0, 0, time.UTC)
	hd := humanDate(tm)

	if hd != "02 Mar 1999 UTC 12:00" {
		t.Errorf("got %q; want %q", hd, "02 Mar 1999 UTC 12:00")
	}
}
