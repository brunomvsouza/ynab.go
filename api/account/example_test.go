package account_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
)

func ExampleService_GetAccount() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	account, _ := c.Account().GetAccount("<valid_budget_id>", "<valid_account_id>")
	fmt.Println(reflect.TypeOf(account))

	// Output: *account.Account
}

func ExampleService_GetAccounts() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	accounts, _ := c.Account().GetAccounts("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(accounts))

	// Output: []*account.Account
}
