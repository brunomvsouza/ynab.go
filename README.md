# YNAB API Go Library

[![Build Status](https://travis-ci.com/brunomvsouza/ynab.go.svg?branch=master)](https://travis-ci.com/brunomvsouza/ynab.go) [![Build Status](https://ci.appveyor.com/api/projects/status/0m2n0q97usi79q27/branch/master?svg=true)](https://ci.appveyor.com/project/brunomvsouza/ynab-go-y5fjk/branch/master) [![Go Report Card](https://goreportcard.com/badge/github.com/brunomvsouza/ynab.go)](https://goreportcard.com/report/github.com/brunomvsouza/ynab.go) [![GoDoc Reference](https://godoc.org/github.com/brunomvsouza/ynab.go?status.svg)](https://godoc.org/github.com/brunomvsouza/ynab.go)

This is an UNOFFICIAL Go client for the YNAB API. It covers 100% of the resources made available by the [YNAB API](https://api.youneedabudget.com).

## Installation

```
go get github.com/brunomvsouza/ynab.go
```

## Usage

To use this client you must [obtain an access token](https://api.youneedabudget.com/#authentication-overview) from your [My Account](https://app.youneedabudget.com/settings) page of the YNAB web app.

```go
package main

import (
	"fmt"

	"github.com/brunomvsouza/ynab.go"
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

See the [godoc](https://godoc.org/github.com/brunomvsouza/ynab.go) to see all the available methods with example usage.

## Development

- Make sure you have Go 1.11 or later installed
- Make sure you have exported `GO111MODULE=on` in your environment to be able do handle dependencies
- Run tests with `go test -race ./...`

## License

BSD-2-Clause
