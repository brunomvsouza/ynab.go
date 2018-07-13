package api

import (
	"fmt"
	"strings"
	"time"
)

// DateLayout expected layout format for the Date type
const DateLayout = "2006-01-02"

// Date represents a budget date
type Date struct {
	time.Time
}

// UnmarshalJSON parses the expected format for a Date
func (d *Date) UnmarshalJSON(b []byte) error {
	// b value comes in surrounded by quotes
	s := strings.Trim(string(b), "\"")

	date, err := DateFromString(s)
	if err != nil {
		return err
	}

	*d = date
	return err
}

// MarshalJSON parses the expected format for a Date
func (d *Date) MarshalJSON() ([]byte, error) {
	val := d.Format(DateLayout)
	return []byte(fmt.Sprintf(`"%s"`, val)), nil
}

// DateFromString creates a new Date from a given string date
// formatted as DateLayout
func DateFromString(s string) (Date, error) {
	t, err := time.Parse(DateLayout, s)
	if err != nil {
		return Date{}, err
	}
	d := Date{
		Time: t,
	}
	return d, nil
}
