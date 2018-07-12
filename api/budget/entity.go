package budget

import (
	"time"

	"bmvs.io/ynab/api"
	"bmvs.io/ynab/api/account"
	"bmvs.io/ynab/api/category"
	"bmvs.io/ynab/api/month"
	"bmvs.io/ynab/api/payee"
	"bmvs.io/ynab/api/transaction"
)

// ResumedBudget represents a budget
type ResumedBudget struct {
	Settings

	ID   string `json:"id"`
	Name string `json:"name"`

	LastModifiedOn *time.Time `json:"last_modified_on"`
	// FirstMonth undocumented field
	FirstMonth *api.Date `json:"first_month"`
	// LastMonth undocumented field
	LastMonth *api.Date `json:"last_month"`
}

// Budget represents the extended version of a budget
type Budget struct {
	ResumedBudget

	Accounts                 []*account.Account
	Payees                   []*payee.Payee
	PayeeLocations           []*payee.Location
	Categories               []*category.Category
	CategoryGroups           []*category.Group
	Months                   []*month.Month
	Transactions             []*transaction.Transaction
	SubTransactions          []*transaction.Sub
	ScheduledTransactions    []*transaction.Scheduled
	ScheduledSubTransactions []*transaction.ScheduledSub
}

// BudgetSummary represents a snapshot of an entire budget
type BudgetSummary struct {
	Budget          *Budget
	ServerKnowledge int64
}

// Settings represents budget settings
type Settings struct {
	DateFormat     DateFormat     `json:"date_format"`
	CurrencyFormat CurrencyFormat `json:"currency_format"`
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
