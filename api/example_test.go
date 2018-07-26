package api_test

import (
	"fmt"

	"go.bmvs.io/ynab/api"
)

func ExampleDateFromString() {
	date, _ := api.DateFromString("2020-01-20")
	fmt.Println(date)

	// Output: 2020-01-20 00:00:00 +0000 UTC
}

func ExampleParseRateLimit() {
	r, _ := api.ParseRateLimit("1/200")
	fmt.Println(r.Used())

	// Output: 1
}
