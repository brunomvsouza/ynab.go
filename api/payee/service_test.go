package payee_test

import (
	"net/http"
	"testing"

	"bmvs.io/ynab"
	"bmvs.io/ynab/api/payee"
	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"
)

func TestService_GetPayees(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "payees": [
      {
        "id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
        "name": "Supermarket",
        "transfer_account_id": null,
        "deleted": false
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	payees, err := client.Payee().GetPayees("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	expected := []*payee.Payee{
		{
			ID:      "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			Name:    "Supermarket",
			Deleted: false,
		},
	}

	assert.Equal(t, expected, payees)
}

func TestService_GetPayeeByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees/34e88373-ef48-4386-9ab3-7f86c2a8988f"
	httpmock.RegisterResponder("GET", url,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
		"payee": {
			"id": "34e88373-ef48-4386-9ab3-7f86c2a8988f",
			"name": "Supermarket",
			"transfer_account_id": null,
			"deleted": false
		}
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	p, err := client.Payee().GetPayeeByID(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"34e88373-ef48-4386-9ab3-7f86c2a8988f",
	)
	assert.NoError(t, err)

	expected := &payee.Payee{
		ID:      "34e88373-ef48-4386-9ab3-7f86c2a8988f",
		Name:    "Supermarket",
		Deleted: false,
	}

	assert.Equal(t, expected, p)
}
