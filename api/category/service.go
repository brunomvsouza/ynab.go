package category

import (
	"fmt"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new category service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB budget API endpoints
type Service struct {
	c api.Getter
}

// GetCategories fetches the list of category groups for a budget
// https://api.youneedabudget.com/v1#/Categories/getCategories
func (s *Service) GetCategories(budgetID string) ([]*Group, error) {
	resModel := struct {
		Data struct {
			CategoryGroups []*Group `json:"category_groups"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.CategoryGroups, nil
}
