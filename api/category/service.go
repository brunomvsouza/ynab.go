package category

import (
	"fmt"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new category service instance
func NewService(c api.ClientReader) *Service {
	return &Service{c}
}

// Service wraps YNAB category API endpoints
type Service struct {
	c api.ClientReader
}

// GetCategories fetches the list of category groups for a budget
// https://api.youneedabudget.com/v1#/Categories/getCategories
func (s *Service) GetCategories(budgetID string) ([]*GroupWithCategories, error) {
	resModel := struct {
		Data struct {
			CategoryGroups []*GroupWithCategories `json:"category_groups"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.CategoryGroups, nil
}

// GetCategory fetches a specific category from a budget
// https://api.youneedabudget.com/v1#/Categories/getCategoryById
func (s *Service) GetCategory(budgetID, categoryID string) (*Category, error) {
	resModel := struct {
		Data struct {
			Category *Category `json:"category"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories/%s", budgetID, categoryID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Category, nil
}
