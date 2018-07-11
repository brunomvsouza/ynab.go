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
	t.Run("success scenarios", func(t *testing.T) {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/5adccdb0-9007-42aa-a6fe-02a3e94476be/accounts",
			func(req *http.Request) (*http.Response, error) {
				assert.Equal(t, "application/json", req.Header.Get("Accept"))
				assert.Equal(t, "Bearer [XtokenX]", req.Header.Get("Authorization"))
				return httpmock.NewStringResponse(200, `{
  "data": {
    "accounts": [ 
      {
        "id": "a6248caa-eed7-4575-a990-717386438d2c",
        "name": "Test Account",
        "type": "checking",
        "on_budget": true,
        "closed": true,
        "balance": 0,
        "cleared_balance": 0,
        "uncleared_balance": 0,
        "deleted": false
      },
			{
        "id": "b6248caa-eed7-4575-a990-717386438d2c",
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

		client := ynab.NewClient("[XtokenX]")
		accounts, err := client.Account().GetAccounts("5adccdb0-9007-42aa-a6fe-02a3e94476be")
		assert.NoError(t, err)

		t.Run("with filled optional fields", func(t *testing.T) {
			assert.Equal(t, "b6248caa-eed7-4575-a990-717386438d2c", accounts[1].ID)
			assert.Equal(t, "Test Account 2", accounts[1].Name)
			assert.Equal(t, account.TypeSavings, accounts[1].Type)
			assert.Equal(t, false, accounts[1].OnBudget)
			assert.Equal(t, true, accounts[1].Closed)
			assert.EqualValues(t, "omg omg omg", *accounts[1].Note)
			assert.Equal(t, int64(-123930), accounts[1].Balance)
			assert.Equal(t, int64(-123930), accounts[1].ClearedBalance)
			assert.Equal(t, int64(0), accounts[1].UnclearedBalance)
			assert.Equal(t, false, accounts[1].Deleted)
		})

		t.Run("with empty optional fields", func(t *testing.T) {
			assert.Equal(t, "a6248caa-eed7-4575-a990-717386438d2c", accounts[0].ID)
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
	})

}
