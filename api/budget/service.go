package budget

import (
	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new budget service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB budget API endpoints
type Service struct {
	c api.Getter
}

// GetBudgets fetches the list of budgets of the logger in user
// https://api.youneedabudget.com/v1#/Budgets/getBudgets
func (s *Service) GetBudgets() ([]*Budget, error) {
	resModel := struct {
		Data struct {
			Budgets []*Budget `json:"budgets"`
		} `json:"data"`
	}{}

	if err := s.c.GET("/budgets", &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Budgets, nil
}
