package budget // import "bmvs.io/ynab/api/budget"

import (
	"time"

	"bmvs.io/ynab/api"
	"bmvs.io/ynab/api/account"
	"bmvs.io/ynab/api/category"
	"bmvs.io/ynab/api/month"
	"bmvs.io/ynab/api/payee"
	"bmvs.io/ynab/api/transaction"
)

// Budget represents a budget
type Budget struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`

	Accounts                 []*account.Account                     `json:"accounts"`
	Payees                   []*payee.Payee                         `json:"payees"`
	PayeeLocations           []*payee.Location                      `json:"payee_locations"`
	Categories               []*category.Category                   `json:"categories"`
	CategoryGroups           []*category.Group                      `json:"category_groups"`
	Months                   []*month.Month                         `json:"months"`
	Transactions             []*transaction.Summary                 `json:"transactions"`
	SubTransactions          []*transaction.SubTransaction          `json:"subtransactions"`
	ScheduledTransactions    []*transaction.ScheduledSummary        `json:"scheduled_transactions"`
	ScheduledSubTransactions []*transaction.ScheduledSubTransaction `json:"scheduled_sub_transactions"`

	LastModifiedOn *time.Time `json:"last_modified_on"`
	// FirstMonth undocumented field
	FirstMonth *api.Date `json:"first_month"`
	// LastMonth undocumented field
	LastMonth *api.Date `json:"last_month"`
}

// Summary represents the summary of a budget
type Summary struct {
	ID             string         `json:"id"`
	Name           string         `json:"name"`
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`

	LastModifiedOn *time.Time `json:"last_modified_on"`
	// FirstMonth undocumented field
	FirstMonth *api.Date `json:"first_month"`
	// LastMonth undocumented field
	LastMonth *api.Date `json:"last_month"`
}

// BudgetDetail represents a versioned snapshot for a budget
type BudgetDetail struct {
	Budget          *Budget
	ServerKnowledge int64
}

// Settings represents the settings for a budget
type Settings struct {
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
}

// DateFormat represents date format for a budget settings
type DateFormat struct {
	Format string `json:"format"`
}

// CurrencyFormat represents a currency format for a budget settings
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
