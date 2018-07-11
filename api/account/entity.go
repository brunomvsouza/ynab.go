package account

// Account represents an account
type Account struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Type             Type   `json:"type"`
	OnBudget         bool   `json:"on_budget"`
	Balance          int64  `json:"balance"`
	ClearedBalance   int64  `json:"cleared_balance"`
	UnclearedBalance int64  `json:"uncleared_balance"`
	Closed           bool   `json:"closed"`
	Deleted          bool   `json:"deleted"`

	Note *string `json:"note"`
}
