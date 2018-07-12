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

func TestService_GetBudgetByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "budget": {
      "id": "aa248caa-eed7-4575-a990-717386438d2c",
      "name": "Test Budget",
      "last_modified_on": "2018-03-05T17:24:36+00:00",
      "date_format": {
        "format": "DD/MM/YYYY"
      },
      "currency_format": {
        "iso_code": "BRL",
        "example_format": "123.456,78",
        "decimal_digits": 2,
        "decimal_separator": ",",
        "symbol_first": true,
        "group_separator": ".",
        "currency_symbol": "R$",
        "display_symbol": true
      },
      "first_month": "2017-12-01",
      "last_month": "2018-02-01",
      "accounts": [
        {
          "id": "312bf0ae-9d1a-42d7-84c1-8f1d5e4e7bb0",
          "name": "Cash",
          "type": "cash",
          "on_budget": true,
          "closed": false,
          "note": null,
          "balance": 0,
          "cleared_balance": 0,
          "uncleared_balance": 0,
          "deleted": false
        }
			],
			"payees": [
        {
          "id": "793846ad-f8f5-454e-9ae4-8d938d0d89ca",
          "name": "Starting Balance",
          "transfer_account_id": null,
          "deleted": false
        }
			],
			"payee_locations": [
        {
          "id": "47471638-da3e-4cdd-9288-e373b50fafa7",
          "payee_id": "793846ad-f8f5-454e-9ae4-8d938d0d89ca",
          "latitude": "20.8988754",
          "longitude": "-33.9167891",
          "deleted": false
        }
			],
			"category_groups": [
        {
          "id": "840512c5-3b1d-426f-b033-f7c64a16a076",
          "name": "Category group",
          "hidden": false,
          "deleted": false
        }
			],
			"categories": [
        {
          "id": "138c8bcd-6ca3-4c09-82ca-1cde7aa1d6f8",
          "category_group_id": "840512c5-3b1d-426f-b033-f7c64a16a076",
          "name": "Category",
          "hidden": false,
          "original_category_group_id": null,
          "note": null,
          "budgeted": 0,
          "activity": 12190,
          "balance": 18740,
          "deleted": false
        }
			],
			"months": [
        {
          "month": "2018-03-01",
          "note": null,
          "to_be_budgeted": 0,
          "age_of_money": null,
          "categories": [
            {
              "id": "138c8bcd-6ca3-4c09-82ca-1cde7aa1d6f8",
          		"category_group_id": "840512c5-3b1d-426f-b033-f7c64a16a076",
              "name": "Category",
              "hidden": true,
              "note": null,
              "budgeted": 0,
              "activity": 12190,
              "balance": 18740,
              "deleted": false
            }
					]
				}
			],
			"transactions": [
        {
          "id": "e31928db-b236-4c88-9a99-7aa46ff7a6f7",
          "date": "2018-01-09",
          "amount": -85440,
          "memo": null,
          "cleared": "cleared",
          "approved": true,
          "flag_color": null,
          "account_id": "312bf0ae-9d1a-42d7-84c1-8f1d5e4e7bb0",
          "payee_id": "fa8d442e-0bfc-4386-8e5b-480c4f70733a",
          "category_id": "0d3552a4-49da-4191-bac6-e22f80eb2056",
          "transfer_account_id": null,
          "import_id": null,
          "deleted": false
        }
			],
			"subtransactions": [
        {
          "id": "254049fe-cadc-4657-b36e-99baac0bd9ca",
          "transaction_id": "891a41b8-bc0f-4c0b-b3a3-97d5d6d61276",
          "amount": 0,
          "memo": null,
          "payee_id": "33fc3c91-8489-4da7-aef5-57ccd19d60dd",
          "category_id": "2d9e60f6-0c7e-472f-8064-0465aa1c58d4",
          "transfer_account_id": null,
          "deleted": false
        }
			],
      "scheduled_transactions": [
        {
          "id": "0971ec91-0961-42be-8598-c6d79c800b28",
          "date_first": "2018-11-20",
          "date_next": "2018-11-20",
          "frequency": "never",
          "amount": -17000,
          "memo": "Domain bmvs.me",
          "flag_color": "yellow",
          "account_id": "09eaca5e-2a34-4baa-89c4-828fb90638f2",
          "payee_id": "793846ad-f8f5-454e-9ae4-8d938d0d89ca",
          "category_id": "138c8bcd-6ca3-4c09-82ca-1cde7aa1d6f8",
          "transfer_account_id": null,
          "deleted": false
        }
			],
      "scheduled_subtransactions": []
		},
    "server_knowledge": 473
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		_, err := client.Budget().GetBudgetByID("aa248caa-eed7-4575-a990-717386438d2c")
		assert.NoError(t, err)
	})
}

func TestService_GetBudgetDeltaByID(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c?last_knowledge_of_server=470",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "budget": {
      "id": "aa248caa-eed7-4575-a990-717386438d2c",
      "name": "Test Budget",
      "last_modified_on": "2018-03-05T17:24:36+00:00",
      "date_format": {
        "format": "DD/MM/YYYY"
      },
      "currency_format": {
        "iso_code": "BRL",
        "example_format": "123.456,78",
        "decimal_digits": 2,
        "decimal_separator": ",",
        "symbol_first": true,
        "group_separator": ".",
        "currency_symbol": "R$",
        "display_symbol": true
      },
      "first_month": "2017-12-01",
      "last_month": "2018-02-01",
      "accounts": [
        {
          "id": "312bf0ae-9d1a-42d7-84c1-8f1d5e4e7bb0",
          "name": "Cash",
          "type": "cash",
          "on_budget": true,
          "closed": false,
          "note": null,
          "balance": 0,
          "cleared_balance": 0,
          "uncleared_balance": 0,
          "deleted": false
        }
			],
			"payees": [
        {
          "id": "793846ad-f8f5-454e-9ae4-8d938d0d89ca",
          "name": "Starting Balance",
          "transfer_account_id": null,
          "deleted": false
        }
			],
			"payee_locations": [
        {
          "id": "47471638-da3e-4cdd-9288-e373b50fafa7",
          "payee_id": "793846ad-f8f5-454e-9ae4-8d938d0d89ca",
          "latitude": "20.8988754",
          "longitude": "-33.9167891",
          "deleted": false
        }
			],
			"category_groups": [
        {
          "id": "840512c5-3b1d-426f-b033-f7c64a16a076",
          "name": "Category group",
          "hidden": false,
          "deleted": false
        }
			],
			"categories": [
        {
          "id": "138c8bcd-6ca3-4c09-82ca-1cde7aa1d6f8",
          "category_group_id": "840512c5-3b1d-426f-b033-f7c64a16a076",
          "name": "Category",
          "hidden": false,
          "original_category_group_id": null,
          "note": null,
          "budgeted": 0,
          "activity": 12190,
          "balance": 18740,
          "deleted": false
        }
			],
			"months": [
        {
          "month": "2018-03-01",
          "note": null,
          "to_be_budgeted": 0,
          "age_of_money": null,
          "categories": [
            {
              "id": "138c8bcd-6ca3-4c09-82ca-1cde7aa1d6f8",
          		"category_group_id": "840512c5-3b1d-426f-b033-f7c64a16a076",
              "name": "Category",
              "hidden": true,
              "note": null,
              "budgeted": 0,
              "activity": 12190,
              "balance": 18740,
              "deleted": false
            }
					]
				}
			],
			"transactions": [
        {
          "id": "e31928db-b236-4c88-9a99-7aa46ff7a6f7",
          "date": "2018-01-09",
          "amount": -85440,
          "memo": null,
          "cleared": "cleared",
          "approved": true,
          "flag_color": null,
          "account_id": "312bf0ae-9d1a-42d7-84c1-8f1d5e4e7bb0",
          "payee_id": "fa8d442e-0bfc-4386-8e5b-480c4f70733a",
          "category_id": "0d3552a4-49da-4191-bac6-e22f80eb2056",
          "transfer_account_id": null,
          "import_id": null,
          "deleted": false
        }
			],
			"subtransactions": [],
      "scheduled_transactions": [],
      "scheduled_subtransactions": []
		},
    "server_knowledge": 473
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		_, err := client.Budget().GetBudgetDeltaByID("aa248caa-eed7-4575-a990-717386438d2c", 470)
		assert.NoError(t, err)
	})
}

func TestService_GetBudgetSettingsByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/settings",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "settings": {
      "date_format": {
        "format": "DD/MM/YYYY"
      },
      "currency_format": {
        "iso_code": "BRL",
        "example_format": "123.456,78",
        "decimal_digits": 2,
        "decimal_separator": ",",
        "symbol_first": true,
        "group_separator": ".",
        "currency_symbol": "R$",
        "display_symbol": true
      }
    }
  }
}`), nil
		},
	)

	client := ynab.NewClient("")
	settings, err := client.Budget().GetBudgetSettingsByID("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	assert.Equal(t, "DD/MM/YYYY", settings.DateFormat.Format)
	assert.Equal(t, "BRL", settings.CurrencyFormat.ISOCode)
	assert.Equal(t, "123.456,78", settings.CurrencyFormat.ExampleFormat)
	assert.Equal(t, uint64(2), settings.CurrencyFormat.DecimalDigits)
	assert.Equal(t, ",", settings.CurrencyFormat.DecimalSeparator)
	assert.True(t, settings.CurrencyFormat.SymbolFirst)
	assert.Equal(t, ".", settings.CurrencyFormat.GroupSeparator)
	assert.Equal(t, "R$", settings.CurrencyFormat.CurrencySymbol)
	assert.True(t, settings.CurrencyFormat.DisplaySymbol)
}
