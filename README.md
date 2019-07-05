# YNAB API Go Library

![Maintained](https://img.shields.io/maintenance/yes/2019.svg) [![GoDoc Reference](https://godoc.org/go.bmvs.io/ynab?status.svg)](https://godoc.org/go.bmvs.io/ynab) [![Build Status](https://lab.bmvs.io/_/ynab.go/badges/master/build.svg)](https://lab.bmvs.io/_/ynab.go) [![Build status](https://ci.appveyor.com/api/projects/status/bik3wi4c9i0ut2u5/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/ynab-go/branch/master) [![Coverage Status](https://coveralls.io/repos/github/brunomvsouza/ynab.go/badge.svg?branch=master)](https://coveralls.io/github/brunomvsouza/ynab.go) [![Go Report Card](https://goreportcard.com/badge/github.com/brunomvsouza/ynab.go)](https://goreportcard.com/report/github.com/brunomvsouza/ynab.go)

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

- Make sure you have Go 1.11 or later installed
- Make sure you have exported `GO111MODULE=on` in your environment to be able do handle dependencies
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
