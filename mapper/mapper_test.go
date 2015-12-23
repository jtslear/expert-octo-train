package mapper

import (
	"fmt"
	"testing"
)

func TestGetDuration(t *testing.T) {
	for i, c := range []struct {
		test, origin, dest, duration string
	}{
		{"works", "New York City, NY", "Raleigh, NC", "7 hours 50 mins"},
	} {
		got, err := GetDuration(c.origin, c.dest)
		if got != c.duration {
			t.Error("Error:", i, c.test, got)
		}
		if err != nil {
			fmt.Print("got")
		}
	}
	for i, c := range []struct {
		test, origin, dest, duration string
	}{
		{"should fail, missing origin", "", "Raleigh, NC", "7 hours 50 mins"},
		{"should fail, missing dest", "New York City, NY", "", "7 hours 50 mins"},
	} {
		got, err := GetDuration(c.origin, c.dest)
		if got != "" {
			t.Error("Error:", i, c.test, got)
		}
		if err != nil {
			fmt.Print(i, got)
		}
	}
}
