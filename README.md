# YNAB API Go Library

[![GoDoc Reference](https://godoc.org/go.bmvs.io/ynab?status.svg)](http://godoc.org/go.bmvs.io/ynab) [![Build Status](https://travis-ci.com/brunomvsouza/ynab.go.svg?branch=master)](https://travis-ci.com/brunomvsouza/ynab.go) [![Build status](https://ci.appveyor.com/api/projects/status/bik3wi4c9i0ut2u5/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/ynab-go/branch/master) [![Coverage Status](https://coveralls.io/repos/github/brunomvsouza/ynab.go/badge.svg?branch=master)](https://coveralls.io/github/brunomvsouza/ynab.go)

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

See the [godoc](https://godoc.org/go.bmvs.io/ynab) to see all the available methods with example usage.

## Development

- Install dependencies with [`dep`](https://github.com/golang/dep)
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
