package transaction

import (
	"fmt"

	"strings"

	"bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new transaction service instance
func NewService(c api.Getter) *Service {
	return &Service{c}
}

// Service wraps YNAB transaction API endpoints
type Service struct {
	c api.Getter
}

// GetTransactions fetches the list of transactions from
// a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactions
func (s *Service) GetTransactions(budgetID string, f *Filter) ([]*Transaction, error) {
	resModel := struct {
		Data struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions", budgetID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetTransaction fetches a specific transaction from a budget
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsById
func (s *Service) GetTransaction(budgetID, transactionID string) (*Transaction, error) {
	resModel := struct {
		Data struct {
			Transaction *Transaction `json:"transaction"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/transactions/%s", budgetID, transactionID)
	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.Transaction, nil
}

// GetTransactionsByAccount fetches the list of transactions of a specific account
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByAccount
func (s *Service) GetTransactionsByAccount(budgetID, accountID string,
	f *Filter) ([]*Transaction, error) {

	resModel := struct {
		Data struct {
			Transactions []*Transaction `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/accounts/%s/transactions", budgetID, accountID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetTransactionsByCategory fetches the list of transactions of a specific category
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByCategory
func (s *Service) GetTransactionsByCategory(budgetID, categoryID string,
	f *Filter) ([]*Hybrid, error) {

	resModel := struct {
		Data struct {
			Transactions []*Hybrid `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/categories/%s/transactions", budgetID, categoryID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// GetTransactionsByPayee fetches the list of transactions of a specific payee
// from a budget with filtering capabilities
// https://api.youneedabudget.com/v1#/Transactions/getTransactionsByPayee
func (s *Service) GetTransactionsByPayee(budgetID, payeeID string,
	f *Filter) ([]*Hybrid, error) {

	resModel := struct {
		Data struct {
			Transactions []*Hybrid `json:"transactions"`
		} `json:"data"`
	}{}

	url := fmt.Sprintf("/budgets/%s/payees/%s/transactions", budgetID, payeeID)
	if f != nil {
		url = fmt.Sprintf("%s?%s", url, f.ToQuery())
	}

	if err := s.c.GET(url, &resModel); err != nil {
		return nil, err
	}

	return resModel.Data.Transactions, nil
}

// Filter represents the optional filter while fetching transactions
type Filter struct {
	SinceDate *api.Date
	Type      *Status
}

// ToQuery returns the filters as a HTTP query string
func (f *Filter) ToQuery() string {
	pairs := make([]string, 0, 2)
	if f.SinceDate != nil && !f.SinceDate.IsZero() {
		pairs = append(pairs, fmt.Sprintf("since_date=%s",
			f.SinceDate.Format(api.DateLayout)))
	}
	if f.Type != nil {
		pairs = append(pairs, fmt.Sprintf("type=%s", string(*f.Type)))
	}
	return strings.Join(pairs, "&")
}
