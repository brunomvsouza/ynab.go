# YNAB API Go Library

This is an UNOFFICIAL Go client for the YNAB API. It covers 100% of the resources made available by the [YNAB API](https://api.youneedabudget.com).

## Installation

```
go get go.bmvs.io/ynab
```

## Usage

To use this client you must [obtain an access token](https://api.youneedabudget.com/#authentication-overview) from your [My Account](https://app.youneedabudget.com/settings) page of the YNAB web app.

```go
package main

import (
	"fmt"

	"go.bmvs.io/ynab"
)

const accessToken = "bf0cbb14b4330-not-real-3de12e66a389eaafe2"

func main() {
	c := ynab.NewClient(accessToken)
	budgets, err := c.Budget().GetBudgets()
	if err != nil {
		panic(err)
	}
	
	for _, budget := range budgets {
		fmt.Println(budget.Name)
		// ...
	}
}
```

See the [godoc](godoc.org/go.bmvs.io/ynab) to see all the available methods with example usage.

## Development

- Install dependencies with [`dep`](https://github.com/golang/dep)
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
