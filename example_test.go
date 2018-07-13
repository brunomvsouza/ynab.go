package ynab_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
)

func ExampleNewClient() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	u, _ := c.User().GetUser()
	fmt.Println(u.ID)

	// Output: a9398633-7fb4-4951-a3c3-3fa425606be0
}

func ExampleClient_User() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.User()
	fmt.Println(reflect.TypeOf(s))

	// Output: *user.Service
}

func ExampleClient_Account() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Account()
	fmt.Println(reflect.TypeOf(s))

	// Output: *account.Service
}

func ExampleClient_Budget() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Budget()
	fmt.Println(reflect.TypeOf(s))

	// Output: *budget.Service
}

func ExampleClient_Category() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Category()
	fmt.Println(reflect.TypeOf(s))

	// Output: *category.Service
}

func ExampleClient_Month() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Month()
	fmt.Println(reflect.TypeOf(s))

	// Output: *month.Service
}

func ExampleClient_Payee() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Payee()
	fmt.Println(reflect.TypeOf(s))

	// Output: *payee.Service
}

func ExampleClient_Transaction() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s := c.Transaction()
	fmt.Println(reflect.TypeOf(s))

	// Output: *transaction.Service
}
