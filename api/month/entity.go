package month // import "bmvs.io/ynab/api/month"

import (
	"bmvs.io/ynab/api"
	"bmvs.io/ynab/api/category"
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
}

// Month represents the summary of a month for a budget
// Each budget contains one or more months, which is where To be Budgeted,
// Age of Money and Category (budgeted / activity / balances)
// amounts are available.
type Summary struct {
	Month api.Date `json:"month"`

	Note         *string `json:"note"`
	ToBeBudgeted *int64  `json:"to_be_budgeted"`
	AgeOfMoney   *int64  `json:"age_of_money"`
}
