package budget

import (
	"fmt"

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
func (s *Service) GetBudgets() ([]*ResumedBudget, error) {
	resModel := struct {
		Data struct {
			Budgets []*ResumedBudget `json:"budgets"`
		} `json:"data"`
	}{}

	if err := s.c.GET("/budgets", &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Budgets, nil
}

// GetBudgetByID fetches a single budget with all related entities,
// effectively a full budget export.
// https://api.youneedabudget.com/v1#/Budgets/getBudgetById
func (s *Service) GetBudgetByID(ID string) (*BudgetSummary, error) {
	resModel := struct {
		Data struct {
			Budget          *Budget `json:"budget"`
			ServerKnowledge int64   `json:"server_knowledge"`
		} `json:"data"`
	}{}

	if err := s.c.GET(fmt.Sprintf("/budgets/%s", ID), &resModel); err != nil {
		return nil, err
	}

	return &BudgetSummary{
		Budget:          resModel.Data.Budget,
		ServerKnowledge: resModel.Data.ServerKnowledge,
	}, nil
}

// GetBudgetDeltaByID fetches the delta of a single budget since the last
// server knowledge number
// https://api.youneedabudget.com/v1#/Budgets/getBudgetById
func (s *Service) GetBudgetDeltaByID(ID string, serverKnowledge int64) (*BudgetSummary, error) {
	resModel := struct {
		Data struct {
			Budget          *Budget `json:"budget"`
			ServerKnowledge int64   `json:"server_knowledge"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s?last_knowledge_of_server=%d", ID, serverKnowledge)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return &BudgetSummary{
		Budget:          resModel.Data.Budget,
		ServerKnowledge: resModel.Data.ServerKnowledge,
	}, nil
}

// GetBudgetSettingsByID fetches a budget settings
// https://api.youneedabudget.com/v1#/Budgets/getBudgetSettingsById
func (s *Service) GetBudgetSettingsByID(ID string) (*Settings, error) {
	resModel := struct {
		Data struct {
			Settings *Settings `json:"settings"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/settings", ID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Settings, nil
}
