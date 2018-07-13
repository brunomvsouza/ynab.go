package account

import (
	"fmt"

	"go.bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new account service instance
func NewService(c api.ClientReader) *Service {
	return &Service{c}
}

// Service wraps YNAB account API endpoints
type Service struct {
	c api.ClientReader
}

// GetAccounts fetches the list of accounts from a budget
// https://api.youneedabudget.com/v1#/Accounts/getAccounts
func (s *Service) GetAccounts(budgetID string) ([]*Account, error) {
	resModel := struct {
		Data struct {
			Accounts []*Account `json:"accounts"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/accounts", budgetID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Accounts, nil
}

// GetAccount fetches a specific account from a budget
// https://api.youneedabudget.com/v1#/Accounts/getAccountById
func (s *Service) GetAccount(budgetID, accountID string) (*Account, error) {
	resModel := struct {
		Data struct {
			Account *Account `json:"account"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/accounts/%s", budgetID, accountID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Account, nil
}
