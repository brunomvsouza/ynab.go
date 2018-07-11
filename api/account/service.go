package account

import (
	"fmt"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new account service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB account API endpoints
type Service struct {
	c api.Getter
}

// GetAccounts fetches the list of accounts from a budget
// https://api.youneedabudget.com/v1#/Accounts/getAccounts
func (s *Service) GetAccounts(budgetID string) ([]*Account, error) {
	url := fmt.Sprintf("/budgets/%s/accounts", budgetID)
	resModel := struct {
		Data struct {
			Accounts []*Account `json:"accounts"`
		} `json:"data"`
	}{}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Accounts, nil
}

// GetAccountByID fetches a specific account from a budget
// https://api.youneedabudget.com/v1#/Accounts/getAccountById
func (s *Service) GetAccountByID(budgetID, accountID string) (*Account, error) {
	url := fmt.Sprintf("/budgets/%s/accounts/%s", budgetID, accountID)
	resModel := struct {
		Data struct {
			Account *Account `json:"account"`
		} `json:"data"`
	}{}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Account, nil
}
