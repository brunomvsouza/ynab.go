package budget_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api/budget"
)

func ExampleService_GetBudget() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	b, _ := c.Budget().GetBudget("<valid_budget_id>", nil)
	fmt.Println(reflect.TypeOf(b))

	// Output: *budget.Snapshot
}

func ExampleService_GetBudget_filtered() {
	c := ynab.NewClient("<valid_ynab_access_token>")

	f := budget.Filter{LastKnowledgeOfServer: 10}
	b, _ := c.Budget().GetBudget("<valid_budget_id>", &f)
	fmt.Println(reflect.TypeOf(b))

	// Output: *budget.Snapshot
}

func ExampleService_GetBudgets() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	budgets, _ := c.Budget().GetBudgets()
	fmt.Println(reflect.TypeOf(budgets))

	// Output: []*budget.Summary
}

func ExampleService_GetBudgetSettings() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	s, _ := c.Budget().GetBudgetSettings("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(s))

	// Output: *budget.Settings
}
