package month_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
)

func TestService_GetMonths(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/months"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "months": [
      {
        "month": "2017-10-01",
        "note": null,
        "to_be_budgeted": 0,
        "age_of_money": 14
      }
		]
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	months, err := client.Month().GetMonths("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	m := months[0]

	var (
		expectedAgeOfMoney   int64 = 14
		expectedToBeBudgeted int64
	)
	assert.Equal(t, "2017-10-01 00:00:00 +0000 UTC", m.Month.String())
	assert.Equal(t, &expectedToBeBudgeted, m.ToBeBudgeted)
	assert.Equal(t, &expectedAgeOfMoney, m.AgeOfMoney)
	assert.Nil(t, m.Note)
}

func TestService_GetMonth(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/months/2017-10-01"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "month": {
			"month": "2017-10-01",
			"note": null,
			"to_be_budgeted": 0,
			"age_of_money": 14
		}
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	date, err := api.DateFromString("2017-10-01")
	assert.NoError(t, err)

	client := ynab.NewClient("")
	m, err := client.Month().GetMonth("aa248caa-eed7-4575-a990-717386438d2c", date)
	assert.NoError(t, err)

	var (
		expectedAgeOfMoney   int64 = 14
		expectedToBeBudgeted int64
	)
	assert.Equal(t, "2017-10-01 00:00:00 +0000 UTC", m.Month.String())
	assert.Equal(t, &expectedToBeBudgeted, m.ToBeBudgeted)
	assert.Equal(t, &expectedAgeOfMoney, m.AgeOfMoney)
	assert.Nil(t, m.Note)
}
