package month_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
)

func ExampleService_GetMonth() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	d, _ := api.DateFromString("2010-01-01")
	m, _ := c.Month().GetMonth("<valid_budget_id>", d)
	fmt.Println(reflect.TypeOf(m))

	// Output: *month.Month
}

func ExampleService_GetMonths() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	months, _ := c.Month().GetMonths("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(months))

	// Output: []*month.Summary
}
