package account_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
	"bmvs.io/ynab/api/account"
)

func TestService_GetAccounts(t *testing.T) {
	t.Run("success with filled optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aadccdb0-9007-42aa-a6fe-02a3e94476be/accounts",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "accounts": [
      {
        "id": "aa248caa-eed7-4575-a990-717386438d2c",
        "name": "Test Account",
        "type": "checking",
        "on_budget": true,
        "closed": true,
        "balance": 0,
        "cleared_balance": 0,
        "uncleared_balance": 0,
        "deleted": false
      }
    ]
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		accounts, err := client.Account().GetAccounts("aadccdb0-9007-42aa-a6fe-02a3e94476be")
		assert.NoError(t, err)

		assert.Equal(t, "aa248caa-eed7-4575-a990-717386438d2c", accounts[0].ID)
		assert.Equal(t, "Test Account", accounts[0].Name)
		assert.Equal(t, account.TypeChecking, accounts[0].Type)
		assert.Nil(t, accounts[0].Note)
		assert.Equal(t, true, accounts[0].OnBudget)
		assert.Equal(t, true, accounts[0].Closed)
		assert.Equal(t, int64(0), accounts[0].Balance)
		assert.Equal(t, int64(0), accounts[0].ClearedBalance)
		assert.Equal(t, int64(0), accounts[0].UnclearedBalance)
		assert.Equal(t, false, accounts[0].Deleted)

	})

	t.Run("success with empty optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/bbdccdb0-9007-42aa-a6fe-02a3e94476be/accounts",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "accounts": [
			{
				"id": "bb248caa-eed7-4575-a990-717386438d2c",
				"name": "Test Account 2",
				"type": "savings",
				"on_budget": false,
				"closed": true,
				"note": "omg omg omg",
				"balance": -123930,
				"cleared_balance": -123930,
				"uncleared_balance": 0,
				"deleted": false
			}
    ]
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		accounts, err := client.Account().GetAccounts("bbdccdb0-9007-42aa-a6fe-02a3e94476be")
		assert.NoError(t, err)

		assert.Equal(t, "bb248caa-eed7-4575-a990-717386438d2c", accounts[0].ID)
		assert.Equal(t, "Test Account 2", accounts[0].Name)
		assert.Equal(t, account.TypeSavings, accounts[0].Type)
		assert.Equal(t, false, accounts[0].OnBudget)
		assert.Equal(t, true, accounts[0].Closed)
		assert.EqualValues(t, "omg omg omg", *accounts[0].Note)
		assert.Equal(t, int64(-123930), accounts[0].Balance)
		assert.Equal(t, int64(-123930), accounts[0].ClearedBalance)
		assert.Equal(t, int64(0), accounts[0].UnclearedBalance)
		assert.Equal(t, false, accounts[0].Deleted)
	})
}

func TestService_GetAccountByID(t *testing.T) {
	t.Run("success with filled optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aadccdb0-9007-42aa-a6fe-02a3e94476be/accounts/aa248caa-eed7-4575-a990-717386438d2c",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "account": {
      "id": "aa248caa-eed7-4575-a990-717386438d2c",
      "name": "Test Account",
      "type": "checking",
      "on_budget": true,
      "closed": true,
      "balance": 0,
      "cleared_balance": 0,
      "uncleared_balance": 0,
      "deleted": false
    }
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		a, err := client.Account().GetAccountByID(
			"aadccdb0-9007-42aa-a6fe-02a3e94476be",
			"aa248caa-eed7-4575-a990-717386438d2c",
		)
		assert.NoError(t, err)

		assert.Equal(t, "aa248caa-eed7-4575-a990-717386438d2c", a.ID)
		assert.Equal(t, "Test Account", a.Name)
		assert.Equal(t, account.TypeChecking, a.Type)
		assert.Nil(t, a.Note)
		assert.Equal(t, true, a.OnBudget)
		assert.Equal(t, true, a.Closed)
		assert.Equal(t, int64(0), a.Balance)
		assert.Equal(t, int64(0), a.ClearedBalance)
		assert.Equal(t, int64(0), a.UnclearedBalance)
		assert.Equal(t, false, a.Deleted)
	})

	t.Run("success with empty optional fields", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/bbdccdb0-9007-42aa-a6fe-02a3e94476be/accounts/bb248caa-eed7-4575-a990-717386438d2c",
			func(req *http.Request) (*http.Response, error) {
				return httpmock.NewStringResponse(200, `{
  "data": {
    "account": {
			"id": "bb248caa-eed7-4575-a990-717386438d2c",
			"name": "Test Account 2",
			"type": "savings",
			"on_budget": false,
			"closed": true,
			"note": "omg omg omg",
			"balance": -123930,
			"cleared_balance": -123930,
			"uncleared_balance": 0,
			"deleted": false
		}
  }
}
		`), nil
			},
		)

		client := ynab.NewClient("")
		a, err := client.Account().GetAccountByID(
			"bbdccdb0-9007-42aa-a6fe-02a3e94476be",
			"bb248caa-eed7-4575-a990-717386438d2c",
		)
		assert.NoError(t, err)

		assert.Equal(t, "bb248caa-eed7-4575-a990-717386438d2c", a.ID)
		assert.Equal(t, "Test Account 2", a.Name)
		assert.Equal(t, account.TypeSavings, a.Type)
		assert.Equal(t, false, a.OnBudget)
		assert.Equal(t, true, a.Closed)
		assert.EqualValues(t, "omg omg omg", *a.Note)
		assert.Equal(t, int64(-123930), a.Balance)
		assert.Equal(t, int64(-123930), a.ClearedBalance)
		assert.Equal(t, int64(0), a.UnclearedBalance)
		assert.Equal(t, false, a.Deleted)
	})
}
