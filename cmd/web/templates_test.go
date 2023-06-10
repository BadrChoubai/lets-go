package main

import (
	"snippetbox.badrchoubai.dev/internal/assert"
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	testTable := []struct {
		name     string
		tm       time.Time
		expected string
	}{
		{
			name:     "UTC",
			tm:       time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			expected: "17 Mar 2022 UTC 10:15",
		},
		{
			name:     "Empty",
			tm:       time.Time{},
			expected: "",
		},
		{
			name:     "CET",
			tm:       time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			expected: "17 Mar 2022 UTC 09:15",
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			actual := humanDate(tt.tm)

			assert.Equal(t, actual, tt.expected)

		})
	}

}
