package month // import "go.bmvs.io/ynab/api/month"

import (
	"go.bmvs.io/ynab/api"
	"go.bmvs.io/ynab/api/category"
)

// Month represents a month for a budget
// Each budget contains one or more months, which is where To be Budgeted,
// Age of Money and Category (budgeted / activity / balances)
// amounts are available.
type Month struct {
	Month      api.Date             `json:"month"`
	Categories []*category.Category `json:"categories"`

	Note         *string `json:"note"`
	ToBeBudgeted *int64  `json:"to_be_budgeted"`
	AgeOfMoney   *int64  `json:"age_of_money"`

	// Income the total amount in transactions categorized to "Inflow: To be Budgeted"
	// in the month (milliunits format)
	Income *int64 `json:"income"`
	// Budgeted the total amount budgeted in the month (milliunits format)
	Budgeted *int64 `json:"budgeted"`
	// Activity the total amount in transactions in the month, excluding those
	// categorized to "Inflow: To be Budgeted" (milliunits format)
	Activity *int64 `json:"activity"`
}

// Summary represents the summary of a month for a budget
// Each budget contains one or more months, which is where To be Budgeted,
// Age of Money and Category (budgeted / activity / balances)
// amounts are available.
type Summary struct {
	Month api.Date `json:"month"`

	Note         *string `json:"note"`
	ToBeBudgeted *int64  `json:"to_be_budgeted"`
	AgeOfMoney   *int64  `json:"age_of_money"`

	// Income the total amount in transactions categorized to "Inflow: To be Budgeted"
	// in the month (milliunits format)
	Income *int64 `json:"income"`
	// Budgeted the total amount budgeted in the month (milliunits format)
	Budgeted *int64 `json:"budgeted"`
	// Activity the total amount in transactions in the month, excluding those
	// categorized to "Inflow: To be Budgeted" (milliunits format)
	Activity *int64 `json:"activity"`
}
