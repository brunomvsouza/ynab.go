package budget

import (
	"strings"
	"time"
)

// Budget represents a budget
type Budget struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	LastModifiedOn *time.Time     `json:"last_modified_on"`
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`

	// FirstMonth undocumented field
	FirstMonth *Date `json:"first_month"`
	// LastMonth undocumented field
	LastMonth *Date `json:"last_month"`
}

// DateFormat represents a date format
type DateFormat struct {
	Format string `json:"format"`
}

// CurrencyFormat represents a currency format
type CurrencyFormat struct {
	ISOCode          string `json:"iso_code"`
	ExampleFormat    string `json:"example_format"`
	DecimalDigits    uint64 `json:"decimal_digits"`
	DecimalSeparator string `json:"decimal_separator"`
	GroupSeparator   string `json:"group_separator"`
	SymbolFirst      bool   `json:"symbol_first"`
	CurrencySymbol   string `json:"currency_symbol"`
	DisplaySymbol    bool   `json:"display_symbol"`
}

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
