// Copyright (c) 2018, Bruno M V Souza <github@b.bmvs.io>. All rights reserved.
// Use of this source code is governed by a BSD-2-Clause license that can be
// found in the LICENSE file.

package category_test

import (
	"fmt"
	"reflect"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
)

func ExampleService_GetCategory() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	category, _ := c.Category().GetCategory("<valid_budget_id>", "<valid_category_id>")
	fmt.Println(reflect.TypeOf(category))

	// Output: *category.Category
}

func ExampleService_GetCategories() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	categories, _ := c.Category().GetCategories("<valid_budget_id>")
	fmt.Println(reflect.TypeOf(categories))

	// Output: []*category.GroupWithCategories
}

func ExampleService_GetCategoryForMonth() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	category, _ := c.Category().GetCategoryForMonth("<valid_budget_id>",
		"<valid_category_id>", api.Date{})
	fmt.Println(reflect.TypeOf(category))

	// Output: *category.Category
}

func ExampleService_GetCategoryForCurrentMonth() {
	c := ynab.NewClient("<valid_ynab_access_token>")
	category, _ := c.Category().GetCategoryForCurrentMonth("<valid_budget_id>",
		"<valid_category_id>")
	fmt.Println(reflect.TypeOf(category))

	// Output: *category.Category
}
