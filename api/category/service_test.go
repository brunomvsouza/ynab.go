package category_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"bmvs.io/ynab"
	"bmvs.io/ynab/api/category"
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

	expected := &category.Group{
		ResumedGroup: category.ResumedGroup{
			ID:      "13419c12-78d3-4818-a5dc-601b2b8a6064",
			Name:    "Credit Card Payments",
			Hidden:  false,
			Deleted: false,
		},
		Categories: []*category.Category{
			{
				ID:              "13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
				CategoryGroupID: "13419c12-78d3-4818-a5dc-601b2b8a6064",
				Name:            "MasterCard",
				Hidden:          false,
				Budgeted:        int64(0),
				Activity:        int64(12190),
				Balance:         int64(18740),
				Deleted:         false,
			},
		},
	}

	assert.Equal(t, expected, groups[0])
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

	expected := &category.Category{
		ID:              "13419c12-78d3-4a26-82ca-1cde7aa1d6f8",
		CategoryGroupID: "13419c12-78d3-4818-a5dc-601b2b8a6064",
		Name:            "MasterCard",
		Hidden:          false,
		Budgeted:        int64(0),
		Activity:        int64(12190),
		Balance:         int64(18740),
		Deleted:         false,
	}
	assert.Equal(t, expected, c)
}
