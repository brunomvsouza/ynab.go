package budget

import (
	"fmt"

	"go.bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new budget service instance
func NewService(c api.ClientReader) *Service {
	return &Service{c}
}

// Service wraps YNAB budget API endpoints
type Service struct {
	c api.ClientReader
}

// GetBudgets fetches the list of budgets of the logger in user
// https://api.youneedabudget.com/v1#/Budgets/getBudgets
func (s *Service) GetBudgets() ([]*Summary, error) {
	resModel := struct {
		Data struct {
			Budgets []*Summary `json:"budgets"`
		} `json:"data"`
	}{}

	if err := s.c.GET("/budgets", &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Budgets, nil
}

// GetBudget fetches a single budget with all related entities,
// effectively a full budget export with filtering capabilities
// https://api.youneedabudget.com/v1#/Budgets/getBudgetById
func (s *Service) GetBudget(budgetID string, f *Filter) (*BudgetDetail, error) {
	resModel := struct {
		Data struct {
			Budget          *Budget `json:"budget"`
			ServerKnowledge int64   `json:"server_knowledge"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s", budgetID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return &BudgetDetail{
		Budget:          resModel.Data.Budget,
		ServerKnowledge: resModel.Data.ServerKnowledge,
	}, nil
}

// GetBudgetSettings fetches a budget settings
// https://api.youneedabudget.com/v1#/Budgets/getBudgetSettingsById
func (s *Service) GetBudgetSettings(budgetID string) (*Settings, error) {
	resModel := struct {
		Data struct {
			Settings *Settings `json:"settings"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/settings", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Settings, nil
}

// Filter represents the optional version filter while fetching a budget
type Filter struct {
	// LastKnowledgeOfServer The starting server knowledge. If provided,
	// only entities that have changed since last_knowledge_of_server
	// will be included
	LastKnowledgeOfServer uint64
}

// ToQuery returns the filters as a HTTP query string
func (f *Filter) ToQuery() string {
	return fmt.Sprintf("last_knowledge_of_server=%d", f.LastKnowledgeOfServer)
}
