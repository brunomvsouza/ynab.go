package transaction

import (
	"fmt"

	"strings"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new transaction service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB transaction API endpoints
type Service struct {
	c api.Getter
}

// GetTransactions fetches the list of transactions from
// a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactions
func (s *Service) GetTransactions(budgetID string, filter *Filter) ([]*Transaction, error) {
	resModel := struct {
		Data struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions", budgetID)
	if filter != nil {
		url = fmt.Sprintf("%s?%s", url, filter.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Transactions, nil
}

// Filter represents the optional filter while fetching transactions
type Filter struct {
	SinceDate *api.Date
	Type      *Type
}

// ToQuery returns the filters as a HTTP query string
func (f *Filter) ToQuery() string {
	pairs := make([]string, 0, 2)
	if f.SinceDate != nil && !f.SinceDate.IsZero() {
		pairs = append(pairs, fmt.Sprintf("since_date=%s",
			f.SinceDate.Format(api.DateLayout)))
	}
	if f.Type != nil {
		pairs = append(pairs, fmt.Sprintf("type=%s", string(*f.Type)))
	}
	return strings.Join(pairs, "&")
}
