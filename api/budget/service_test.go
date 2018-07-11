package budget_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
)

func TestService_GetBudgets(t *testing.T) {
	t.Run("success with filled optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "budgets": [
      {
        "id": "aa248caa-eed7-4575-a990-717386438d2c",
        "name": "TestBudget",
        "last_modified_on": "2018-03-05T17:05:23+00:00",
        "first_month": "2018-03-01",
        "last_month": "2018-04-01",
        "date_format": {
          "format": "DD.MM.YYYY"
        },
        "currency_format": {
          "iso_code": "EUR",
          "example_format": "123,456.78",
          "decimal_digits": 2,
          "decimal_separator": ".",
          "symbol_first": false,
          "group_separator": ",",
          "currency_symbol": "€",
          "display_symbol": true
        }
      }
    ]
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		budgets, err := client.Budget().GetBudgets()
		assert.NoError(t, err)

		budget := budgets[0]
		assert.NotNil(t, budget)

		assert.Equal(t, "aa248caa-eed7-4575-a990-717386438d2c", budget.ID)
		assert.Equal(t, "TestBudget", budget.Name)
		assert.Equal(t, "2018-03-05 17:05:23 +0000 +0000", budget.LastModifiedOn.String())
		assert.Equal(t, "2018-03-01 00:00:00 +0000 UTC", budget.FirstMonth.String())
		assert.Equal(t, "2018-04-01 00:00:00 +0000 UTC", budget.LastMonth.String())
		assert.Equal(t, "DD.MM.YYYY", budget.DateFormat.Format)
		assert.Equal(t, "EUR", budget.CurrencyFormat.ISOCode)
		assert.Equal(t, "123,456.78", budget.CurrencyFormat.ExampleFormat)
		assert.Equal(t, uint64(2), budget.CurrencyFormat.DecimalDigits)
		assert.Equal(t, ".", budget.CurrencyFormat.DecimalSeparator)
		assert.Equal(t, false, budget.CurrencyFormat.SymbolFirst)
		assert.Equal(t, ",", budget.CurrencyFormat.GroupSeparator)
		assert.Equal(t, "€", budget.CurrencyFormat.CurrencySymbol)
		assert.Equal(t, true, budget.CurrencyFormat.DisplaySymbol)
	})

	t.Run("success with empty optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "budgets": [
      {
        "id": "aa248caa-eed7-4575-a990-717386438d2c",
        "name": "TestBudget",
        "date_format": {
          "format": "DD.MM.YYYY"
        },
        "currency_format": {
          "iso_code": "EUR",
          "example_format": "123,456.78",
          "decimal_digits": 2,
          "decimal_separator": ".",
          "symbol_first": false,
          "group_separator": ",",
          "currency_symbol": "€",
          "display_symbol": true
        }
      }
    ]
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		budgets, err := client.Budget().GetBudgets()
		assert.NoError(t, err)

		budget := budgets[0]
		assert.NotNil(t, budget)

		assert.Equal(t, "aa248caa-eed7-4575-a990-717386438d2c", budget.ID)
		assert.Equal(t, "TestBudget", budget.Name)
		assert.Nil(t, budget.LastModifiedOn)
		assert.Nil(t, budget.FirstMonth)
		assert.Nil(t, budget.LastMonth)
		assert.Equal(t, "DD.MM.YYYY", budget.DateFormat.Format)
		assert.Equal(t, "EUR", budget.CurrencyFormat.ISOCode)
		assert.Equal(t, "123,456.78", budget.CurrencyFormat.ExampleFormat)
		assert.Equal(t, uint64(2), budget.CurrencyFormat.DecimalDigits)
		assert.Equal(t, ".", budget.CurrencyFormat.DecimalSeparator)
		assert.Equal(t, false, budget.CurrencyFormat.SymbolFirst)
		assert.Equal(t, ",", budget.CurrencyFormat.GroupSeparator)
		assert.Equal(t, "€", budget.CurrencyFormat.CurrencySymbol)
		assert.Equal(t, true, budget.CurrencyFormat.DisplaySymbol)
	})

}
