package ynab_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
)

func ExampleNewClient() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	c.User().GetUser()
}

func ExampleClientServicer_User() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.User()
	fmt.Println(reflect.TypeOf(s))

	// Output: *user.Service
}

func ExampleClientServicer_Account() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Account()
	fmt.Println(reflect.TypeOf(s))

	// Output: *account.Service
}

func ExampleClientServicer_Budget() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Budget()
	fmt.Println(reflect.TypeOf(s))

	// Output: *budget.Service
}

func ExampleClientServicer_Category() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Category()
	fmt.Println(reflect.TypeOf(s))

	// Output: *category.Service
}

func ExampleClientServicer_Month() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Month()
	fmt.Println(reflect.TypeOf(s))

	// Output: *month.Service
}

func ExampleClientServicer_Payee() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Payee()
	fmt.Println(reflect.TypeOf(s))

	// Output: *payee.Service
}

func ExampleClientServicer_Transaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Transaction()
	fmt.Println(reflect.TypeOf(s))

	// Output: *transaction.Service
}

func ExampleClientServicer_RateLimit() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	r := c.RateLimit()
	fmt.Println(reflect.TypeOf(r))

	// Output: *api.RateLimit
}
