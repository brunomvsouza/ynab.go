// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package category

import (
	"fmt"

	"go.bmvs.io/ynab/api"
)

const currentMonthID = "current"

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

// GetCategoryForMonth fetches a specific category from a budget month
// https://api.youneedabudget.com/v1#/Categories/getMonthCategoryById
func (s *Service) GetCategoryForMonth(budgetID, categoryID string,
	month api.Date) (*Category, error) {

	return s.getCategoryForMonth(budgetID, categoryID, api.DateFormat(month))
}

// GetCategoryForCurrentMonth fetches a specific category from the current budget month
// https://api.youneedabudget.com/v1#/Categories/getMonthCategoryById
func (s *Service) GetCategoryForCurrentMonth(budgetID, categoryID string) (*Category, error) {
	return s.getCategoryForMonth(budgetID, categoryID, currentMonthID)
}

func (s *Service) getCategoryForMonth(budgetID, categoryID, month string) (*Category, error) {
	resModel := struct {
		Data struct {
			Category *Category `json:"category"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/months/%s/categories/%s", budgetID, month, categoryID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Category, nil
}
