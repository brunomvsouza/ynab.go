package user_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
)

func TestService_GetUser(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/user",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "user": {
      "id": "aa248caa-eed7-4575-a990-717386438d2c"
    }
  }
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	u, err := client.User().GetUser()
	assert.NoError(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, "aa248caa-eed7-4575-a990-717386438d2c", u.ID)

}
