package account // import "go.bmvs.io/ynab/api/account"

// Account represents an account for a budget
type Account struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     Type   `json:"type"`
	OnBudget bool   `json:"on_budget"`
	// Balance The current balance of the account in milliunits format
	Balance int64 `json:"balance"`
	// ClearedBalance The current cleared balance of the account in milliunits format
	ClearedBalance int64 `json:"cleared_balance"`
	// ClearedBalance The current uncleared balance of the account in milliunits format
	UnclearedBalance int64 `json:"uncleared_balance"`
	Closed           bool  `json:"closed"`
	// Deleted Deleted accounts will only be included in delta requests
	Deleted bool `json:"deleted"`

	Note *string `json:"note"`
}
