package payee

import (
	"fmt"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new payee service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB payee API endpoints
type Service struct {
	c api.Getter
}

// GetPayees fetches the list of payees from a budget
// https://api.youneedabudget.com/v1#/Payees/getPayees
func (s *Service) GetPayees(budgetID string) ([]*Payee, error) {
	resModel := struct {
		Data struct {
			Payees []*Payee `json:"payees"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Payees, nil
}
