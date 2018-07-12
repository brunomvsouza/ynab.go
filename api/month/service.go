package month

import (
	"fmt"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new month service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB month API endpoints
type Service struct {
	c api.Getter
}

// GetMonths fetches the list of months from a budget
// https://api.youneedabudget.com/v1#/Months/getBudgetMonths
func (s *Service) GetMonths(budgetID string) ([]*Month, error) {
	resModel := struct {
		Data struct {
			Months []*Month `json:"months"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/months", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Months, nil
}

// GetMonth fetches a specific month from a budget
// https://api.youneedabudget.com/v1#/Months/getBudgetMonth
func (s *Service) GetMonth(budgetID string, month api.Date) (*Month, error) {
	resModel := struct {
		Data struct {
			Month *Month `json:"month"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/months/%s", budgetID,
		month.Format(api.DateLayout))
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Month, nil
}
