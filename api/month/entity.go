package month

import (
	"bmvs.io/ynab/api"
	"bmvs.io/ynab/api/category"
)

// Month represents a month
type Month struct {
	Month      api.Date             `json:"month"`
	Categories []*category.Category `json:"categories"`

	Note         *string `json:"note"`
	ToBeBudgeted *int64  `json:"to_be_budgeted"`
	AgeOfMoney   *int64  `json:"age_of_money"`
}
