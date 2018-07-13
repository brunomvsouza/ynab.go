package user

import (
	"go.bmvs.io/ynab/api"
)

// NewService facilitates the creation of a new user service instance
func NewService(c api.ClientReader) *Service {
	return &Service{c}
}

// Service wraps YNAB user API endpoints
type Service struct {
	c api.ClientReader
}

// GetUser fetches information about the authenticated user
// https://api.youneedabudget.com/v1#/User/getUser
func (s *Service) GetUser() (*User, error) {
	resModel := struct {
		Data struct {
			User *User `json:"user"`
		} `json:"data"`
	}{}

	if err := s.c.GET("/user", &resModel); err != nil {
		return nil, err
	}
	return resModel.Data.User, nil
}
