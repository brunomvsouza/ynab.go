package budget_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"bmvs.io/ynab/api/budget"
)

func TestDate_UnmarshalJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *budget.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB", "Date": "2009-01-29"}`), &wrapper)
		assert.NoError(t, err)
		assert.Equal(t, "2009-01-29 00:00:00 +0000 UTC", wrapper.Date.String())
	})

	t.Run("json string without Date field", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *budget.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB"}`), &wrapper)
		assert.NoError(t, err)
		assert.Nil(t, wrapper.Date)
	})

	t.Run("json string with null Date field", func(t *testing.T) {
		wrapper := struct {
			Name string
			Date *budget.Date
		}{}

		err := json.Unmarshal([]byte(`{"Name": "YNAB", "Date": null}`), &wrapper)
		assert.NoError(t, err)
		assert.Nil(t, wrapper.Date)
	})
}
