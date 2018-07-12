package category_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
)

func TestService_GetCategories(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/categories",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "category_groups": [
      {
        "id": "13419c12-78d3-4818-a5dc-601b2b8a6064",
        "name": "Credit Card Payments",
        "hidden": false,
        "deleted": false,
        "categories": [
          {
            "id": "13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
            "category_group_id": "13419c12-78d3-4818-a5dc-601b2b8a6064",
            "name": "MasterCard",
            "hidden": false,
            "original_category_group_id": null,
            "note": null,
            "budgeted": 0,
            "activity": 12190,
            "balance": 18740,
            "deleted": false
          }
        ]
      }
		]
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	groups, err := client.Category().GetCategories("aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	g := groups[0]
	assert.Equal(t, "13419c12-78d3-4818-a5dc-601b2b8a6064", g.ID)
	assert.Equal(t, "Credit Card Payments", g.Name)
	assert.False(t, g.Hidden)
	assert.False(t, g.Deleted)

	c := g.Categories[0]
	assert.Equal(t, "13419c12-78d3-4a26-82ca-1cde7aa1d6f8", c.ID)
	assert.Equal(t, "13419c12-78d3-4818-a5dc-601b2b8a6064", c.CategoryGroupID)
	assert.Equal(t, "MasterCard", c.Name)
	assert.False(t, c.Hidden)
	assert.Nil(t, c.OriginalCategoryGroupID)
	assert.Nil(t, c.Note)
	assert.Equal(t, int64(0), c.Budgeted)
	assert.Equal(t, int64(12190), c.Activity)
	assert.Equal(t, int64(18740), c.Balance)
	assert.False(t, c.Deleted)
}

func TestService_GetCategoryByID(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/categories/13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, `{
  "data": {
    "category": {
			"id": "13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
			"category_group_id": "13419c12-78d3-4818-a5dc-601b2b8a6064",
			"name": "MasterCard",
			"hidden": false,
			"original_category_group_id": null,
			"note": null,
			"budgeted": 0,
			"activity": 12190,
			"balance": 18740,
			"deleted": false
    }
	}
}
		`), nil
		},
	)

	client := ynab.NewClient("")
	c, err := client.Category().GetCategoryByID(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
	)
	assert.NoError(t, err)

	assert.Equal(t, "13419c12-78d3-4a26-82ca-1cde7aa1d6f8", c.ID)
	assert.Equal(t, "13419c12-78d3-4818-a5dc-601b2b8a6064", c.CategoryGroupID)
	assert.Equal(t, "MasterCard", c.Name)
	assert.False(t, c.Hidden)
	assert.Nil(t, c.OriginalCategoryGroupID)
	assert.Nil(t, c.Note)
	assert.Equal(t, int64(0), c.Budgeted)
	assert.Equal(t, int64(12190), c.Activity)
	assert.Equal(t, int64(18740), c.Balance)
	assert.False(t, c.Deleted)
}
