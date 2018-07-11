package api

import (
	"strings"
	"time"
)

// Date represents a budget date
type Date struct {
	time.Time
}

// UnmarshalJSON parses the expected format for a Date
func (t *Date) UnmarshalJSON(b []byte) error {
	// b value comes in surrounded by quotes
	s := strings.Trim(string(b), "\"")

	var err error
	t.Time, err = time.Parse("2006-01-02", s)

	return err
}
