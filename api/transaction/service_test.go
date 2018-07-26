package transaction_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/jarcoal/httpmock.v1"

	"go.bmvs.io/ynab"
	"go.bmvs.io/ynab/api"
	"go.bmvs.io/ynab/api/transaction"
)

func TestService_GetTransactions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "id": "e6ad88f5-6f16-4480-9515-5377012750dd",
        "date": "2018-03-10",
        "amount": -43950,
        "memo": "nice memo",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "09eaca5e-6f16-4480-9515-828fb90638f2",
        "account_name": "Bank Name",
        "payee_id": "6216ab4b-6f16-4480-9515-be2dee26ab0d",
        "payee_name": "Supermarket",
        "category_id": "e9517027-6f16-4480-9515-5981bed2e9e1",
        "category_name": "Split (Multiple Categories)...",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false,
        "subtransactions": [
          {
            "id": "9453526b-2f58-4c02-9683-a30c2a1192d7",
            "transaction_id": "e6ad88f5-6f16-4480-9515-5377012750dd",
            "amount": -33970,
            "memo": "Debit Card Payment",
            "payee_id": "6216ab4b-bb05-4574-b4b5-be2dee26ab0d",
            "category_id": "080985e4-4175-43e4-96bb-d207a9d2c8ce",
            "transfer_account_id": null,
            "deleted": false
          }
				]
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
	transactions, err := client.Transaction().GetTransactions("aa248caa-eed7-4575-a990-717386438d2c", nil)
	assert.NoError(t, err)

	expectedDate, err := api.DateFromString("2018-03-10")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedPayeeID := "6216ab4b-6f16-4480-9515-be2dee26ab0d"
	expectedPayeeName := "Supermarket"
	expectedCategoryID := "e9517027-6f16-4480-9515-5981bed2e9e1"
	expectedCategoryName := "Split (Multiple Categories)..."

	expectedSubTransactionMemo := "Debit Card Payment"
	expectedSubTransactionPayeeID := "6216ab4b-bb05-4574-b4b5-be2dee26ab0d"
	expectedSubTransactionCategoryID := "080985e4-4175-43e4-96bb-d207a9d2c8ce"

	expected := []*transaction.Transaction{
		{
			ID:           "e6ad88f5-6f16-4480-9515-5377012750dd",
			Date:         expectedDate,
			Amount:       int64(-43950),
			Memo:         &expectedMemo,
			Cleared:      transaction.ClearingStatusReconciled,
			Approved:     true,
			AccountID:    "09eaca5e-6f16-4480-9515-828fb90638f2",
			AccountName:  "Bank Name",
			PayeeID:      &expectedPayeeID,
			PayeeName:    &expectedPayeeName,
			CategoryID:   &expectedCategoryID,
			CategoryName: &expectedCategoryName,
			Deleted:      false,
			SubTransactions: []*transaction.SubTransaction{
				{
					ID:            "9453526b-2f58-4c02-9683-a30c2a1192d7",
					TransactionID: "e6ad88f5-6f16-4480-9515-5377012750dd",
					Amount:        int64(-33970),
					Memo:          &expectedSubTransactionMemo,
					PayeeID:       &expectedSubTransactionPayeeID,
					CategoryID:    &expectedSubTransactionCategoryID,
					Deleted:       false,
				},
			},
		},
	}
	assert.Equal(t, expected, transactions)
}

func TestService_GetTransaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions/e6ad88f5-6f16-4480-9515-5377012750dd"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transaction": {
			"id": "e6ad88f5-6f16-4480-9515-5377012750dd",
			"date": "2018-03-10",
			"amount": -43950,
			"memo": "nice memo",
			"cleared": "reconciled",
			"approved": true,
			"flag_color": null,
			"account_id": "09eaca5e-6f16-4480-9515-828fb90638f2",
			"account_name": "Bank Name",
			"payee_id": "6216ab4b-6f16-4480-9515-be2dee26ab0d",
			"payee_name": "Supermarket",
			"category_id": "e9517027-6f16-4480-9515-5981bed2e9e1",
			"category_name": "Split (Multiple Categories)...",
			"transfer_account_id": null,
			"import_id": null,
			"deleted": false,
			"subtransactions": [
			  {
			    "id": "9453526b-2f58-4c02-9683-a30c2a1192d7",
			    "transaction_id": "e6ad88f5-6f16-4480-9515-5377012750dd",
			    "amount": -33970,
			    "memo": "Debit Card Payment",
			    "payee_id": "6216ab4b-bb05-4574-b4b5-be2dee26ab0d",
			    "category_id": "080985e4-4175-43e4-96bb-d207a9d2c8ce",
			    "transfer_account_id": null,
			    "deleted": false
			  }
			]
		}
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	tx, err := client.Transaction().GetTransaction(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"e6ad88f5-6f16-4480-9515-5377012750dd",
	)
	assert.NoError(t, err)

	expectedDate, err := api.DateFromString("2018-03-10")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedPayeeID := "6216ab4b-6f16-4480-9515-be2dee26ab0d"
	expectedPayeeName := "Supermarket"
	expectedCategoryID := "e9517027-6f16-4480-9515-5981bed2e9e1"
	expectedCategoryName := "Split (Multiple Categories)..."

	expectedSubTransactionMemo := "Debit Card Payment"
	expectedSubTransactionPayeeID := "6216ab4b-bb05-4574-b4b5-be2dee26ab0d"
	expectedSubTransactionCategoryID := "080985e4-4175-43e4-96bb-d207a9d2c8ce"

	expected := &transaction.Transaction{
		ID:           "e6ad88f5-6f16-4480-9515-5377012750dd",
		Date:         expectedDate,
		Amount:       int64(-43950),
		Memo:         &expectedMemo,
		Cleared:      transaction.ClearingStatusReconciled,
		Approved:     true,
		AccountID:    "09eaca5e-6f16-4480-9515-828fb90638f2",
		AccountName:  "Bank Name",
		PayeeID:      &expectedPayeeID,
		PayeeName:    &expectedPayeeName,
		CategoryID:   &expectedCategoryID,
		CategoryName: &expectedCategoryName,
		Deleted:      false,
		SubTransactions: []*transaction.SubTransaction{
			{
				ID:            "9453526b-2f58-4c02-9683-a30c2a1192d7",
				TransactionID: "e6ad88f5-6f16-4480-9515-5377012750dd",
				Amount:        int64(-33970),
				Memo:          &expectedSubTransactionMemo,
				PayeeID:       &expectedSubTransactionPayeeID,
				CategoryID:    &expectedSubTransactionCategoryID,
				Deleted:       false,
			},
		},
	}
	assert.Equal(t, expected, tx)
}

func TestService_GetTransactionsByAccount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/accounts/09eaca5e-6f16-4480-9515-828fb90638f2/transactions"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "id": "e6ad88f5-6f16-4480-9515-5377012750dd",
        "date": "2018-03-10",
        "amount": -43950,
        "memo": "nice memo",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "09eaca5e-6f16-4480-9515-828fb90638f2",
        "account_name": "Bank Name",
        "payee_id": "6216ab4b-6f16-4480-9515-be2dee26ab0d",
        "payee_name": "Supermarket",
        "category_id": "e9517027-6f16-4480-9515-5981bed2e9e1",
        "category_name": "Split (Multiple Categories)...",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false,
        "subtransactions": [
          {
            "id": "9453526b-2f58-4c02-9683-a30c2a1192d7",
            "transaction_id": "e6ad88f5-6f16-4480-9515-5377012750dd",
            "amount": -33970,
            "memo": "Debit Card Payment",
            "payee_id": "6216ab4b-bb05-4574-b4b5-be2dee26ab0d",
            "category_id": "080985e4-4175-43e4-96bb-d207a9d2c8ce",
            "transfer_account_id": null,
            "deleted": false
          }
				]
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
	transactions, err := client.Transaction().GetTransactionsByAccount(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"09eaca5e-6f16-4480-9515-828fb90638f2",
		nil,
	)
	assert.NoError(t, err)

	expectedDate, err := api.DateFromString("2018-03-10")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedPayeeID := "6216ab4b-6f16-4480-9515-be2dee26ab0d"
	expectedPayeeName := "Supermarket"
	expectedCategoryID := "e9517027-6f16-4480-9515-5981bed2e9e1"
	expectedCategoryName := "Split (Multiple Categories)..."

	expectedSubTransactionMemo := "Debit Card Payment"
	expectedSubTransactionPayeeID := "6216ab4b-bb05-4574-b4b5-be2dee26ab0d"
	expectedSubTransactionCategoryID := "080985e4-4175-43e4-96bb-d207a9d2c8ce"

	expected := []*transaction.Transaction{
		{
			ID:           "e6ad88f5-6f16-4480-9515-5377012750dd",
			Date:         expectedDate,
			Amount:       int64(-43950),
			Memo:         &expectedMemo,
			Cleared:      transaction.ClearingStatusReconciled,
			Approved:     true,
			AccountID:    "09eaca5e-6f16-4480-9515-828fb90638f2",
			AccountName:  "Bank Name",
			PayeeID:      &expectedPayeeID,
			PayeeName:    &expectedPayeeName,
			CategoryID:   &expectedCategoryID,
			CategoryName: &expectedCategoryName,
			Deleted:      false,
			SubTransactions: []*transaction.SubTransaction{
				{
					ID:            "9453526b-2f58-4c02-9683-a30c2a1192d7",
					TransactionID: "e6ad88f5-6f16-4480-9515-5377012750dd",
					Amount:        int64(-33970),
					Memo:          &expectedSubTransactionMemo,
					PayeeID:       &expectedSubTransactionPayeeID,
					CategoryID:    &expectedSubTransactionCategoryID,
					Deleted:       false,
				},
			},
		},
	}
	assert.Equal(t, expected, transactions)
}

func TestService_GetTransactionsByCategory(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/categories/a33c906e-444c-469c-be27-04c8e0c9959f/transactions"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "type": "transaction",
        "id": "c132c55c-1200-4606-a321-99f4ec24b4df",
        "parent_transaction_id": null,
        "date": "2018-01-10",
        "amount": -42000,
        "memo": "",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "134d159-444c-469c-be27-44094e388fa0",
        "account_name": "Cash",
        "payee_id": "b391144e-444c-469c-be27-fed6aa352a7a",
        "payee_name": "Landlord",
        "category_id": "a33c906e-444c-469c-be27-04c8e0c9959f",
        "category_name": "Rent",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false
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
	transactions, err := client.Transaction().GetTransactionsByCategory(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"a33c906e-444c-469c-be27-04c8e0c9959f",
		nil,
	)
	assert.NoError(t, err)

	expectedDate, err := api.DateFromString("2018-01-10")
	assert.NoError(t, err)

	expectedMemo := ""
	expectedPayeeID := "b391144e-444c-469c-be27-fed6aa352a7a"
	expectedPayeeName := "Landlord"
	expectedCategoryID := "a33c906e-444c-469c-be27-04c8e0c9959f"
	expectedCategoryName := "Rent"

	expected := []*transaction.Hybrid{
		{
			Type:         transaction.TypeTransaction,
			ID:           "c132c55c-1200-4606-a321-99f4ec24b4df",
			Date:         expectedDate,
			Amount:       int64(-42000),
			Memo:         &expectedMemo,
			Cleared:      transaction.ClearingStatusReconciled,
			Approved:     true,
			AccountID:    "134d159-444c-469c-be27-44094e388fa0",
			AccountName:  "Cash",
			PayeeID:      &expectedPayeeID,
			PayeeName:    &expectedPayeeName,
			CategoryID:   &expectedCategoryID,
			CategoryName: &expectedCategoryName,
			Deleted:      false,
		},
	}
	assert.Equal(t, expected, transactions)
}

func TestService_GetTransactionsByPayee(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/payees/b391144e-444c-469c-be27-fed6aa352a7a/transactions"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transactions": [
      {
        "type": "transaction",
        "id": "c132c55c-1200-4606-a321-99f4ec24b4df",
        "parent_transaction_id": null,
        "date": "2018-01-10",
        "amount": -42000,
        "memo": "",
        "cleared": "reconciled",
        "approved": true,
        "flag_color": null,
        "account_id": "134d159-444c-469c-be27-44094e388fa0",
        "account_name": "Cash",
        "payee_id": "b391144e-444c-469c-be27-fed6aa352a7a",
        "payee_name": "Landlord",
        "category_id": "a33c906e-444c-469c-be27-04c8e0c9959f",
        "category_name": "Rent",
        "transfer_account_id": null,
        "import_id": null,
        "deleted": false
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
	transactions, err := client.Transaction().GetTransactionsByPayee(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"b391144e-444c-469c-be27-fed6aa352a7a",
		nil,
	)
	assert.NoError(t, err)

	expectedDate, err := api.DateFromString("2018-01-10")
	assert.NoError(t, err)

	expectedMemo := ""
	expectedPayeeID := "b391144e-444c-469c-be27-fed6aa352a7a"
	expectedPayeeName := "Landlord"
	expectedCategoryID := "a33c906e-444c-469c-be27-04c8e0c9959f"
	expectedCategoryName := "Rent"

	expected := []*transaction.Hybrid{
		{
			Type:         transaction.TypeTransaction,
			ID:           "c132c55c-1200-4606-a321-99f4ec24b4df",
			Date:         expectedDate,
			Amount:       int64(-42000),
			Memo:         &expectedMemo,
			Cleared:      transaction.ClearingStatusReconciled,
			Approved:     true,
			AccountID:    "134d159-444c-469c-be27-44094e388fa0",
			AccountName:  "Cash",
			PayeeID:      &expectedPayeeID,
			PayeeName:    &expectedPayeeName,
			CategoryID:   &expectedCategoryID,
			CategoryName: &expectedCategoryName,
			Deleted:      false,
		},
	}
	assert.Equal(t, expected, transactions)
}

func TestService_GetScheduledTransactions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/scheduled_transactions"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "scheduled_transactions": [
      {
        "id": "56f4fc86-2ed7-4b3b-9116-7a214261b3cd",
        "date_first": "2018-11-13",
        "date_next": "2018-11-13",
        "frequency": "never",
        "amount": -9000,
        "memo": "nice memo",
        "flag_color": "yellow",
        "account_id": "09eaca5e-312a-4bcd-89c4-828fb90638f2",
        "account_name": "Bank Name",
        "payee_id": "0d0e928d-312a-4bcd-89c4-e02f40d1fe46",
        "payee_name": "bla bla bla",
        "category_id": "f3cc4f55-312a-4bcd-89c4-db34379cb1dc",
        "category_name": "Yearly subscription",
        "transfer_account_id": null,
        "deleted": false,
        "subtransactions": []
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
	transactions, err := client.Transaction().GetScheduledTransactions(
		"aa248caa-eed7-4575-a990-717386438d2c")
	assert.NoError(t, err)

	expectedFirstAndLastDate, err := api.DateFromString("2018-11-13")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedFlagColor := transaction.FlagColorYellow
	expectedPayeeID := "0d0e928d-312a-4bcd-89c4-e02f40d1fe46"
	expectedPayeeName := "bla bla bla"
	expectedCategoryID := "f3cc4f55-312a-4bcd-89c4-db34379cb1dc"
	expectedCategoryName := "Yearly subscription"

	expected := []*transaction.Scheduled{
		{
			ID:              "56f4fc86-2ed7-4b3b-9116-7a214261b3cd",
			DateFirst:       expectedFirstAndLastDate,
			DateNext:        expectedFirstAndLastDate,
			Frequency:       transaction.FrequencyNever,
			Amount:          int64(-9000),
			Memo:            &expectedMemo,
			FlagColor:       &expectedFlagColor,
			AccountID:       "09eaca5e-312a-4bcd-89c4-828fb90638f2",
			AccountName:     "Bank Name",
			PayeeID:         &expectedPayeeID,
			PayeeName:       &expectedPayeeName,
			CategoryID:      &expectedCategoryID,
			CategoryName:    &expectedCategoryName,
			Deleted:         false,
			SubTransactions: []*transaction.ScheduledSubTransaction{},
		},
	}
	assert.Equal(t, expected, transactions)
}

func TestService_GetScheduledTransaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/scheduled_transactions/56f4fc86-2ed7-4b3b-9116-7a214261b3cd"
	httpmock.RegisterResponder(http.MethodGet, url,
		func(req *http.Request) (*http.Response, error) {
			res := httpmock.NewStringResponse(200, `{
  "data": {
    "scheduled_transaction": {
			"id": "56f4fc86-2ed7-4b3b-9116-7a214261b3cd",
			"date_first": "2018-11-13",
			"date_next": "2018-11-13",
			"frequency": "never",
			"amount": -9000,
			"memo": "nice memo",
			"flag_color": "yellow",
			"account_id": "09eaca5e-312a-4bcd-89c4-828fb90638f2",
			"account_name": "Bank Name",
			"payee_id": "0d0e928d-312a-4bcd-89c4-e02f40d1fe46",
			"payee_name": "bla bla bla",
			"category_id": "f3cc4f55-312a-4bcd-89c4-db34379cb1dc",
			"category_name": "Yearly subscription",
			"transfer_account_id": null,
			"deleted": false,
			"subtransactions": []
    }
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	stx, err := client.Transaction().GetScheduledTransaction(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"56f4fc86-2ed7-4b3b-9116-7a214261b3cd",
	)
	assert.NoError(t, err)

	expectedFirstAndLastDate, err := api.DateFromString("2018-11-13")
	assert.NoError(t, err)

	expectedMemo := "nice memo"
	expectedFlagColor := transaction.FlagColorYellow
	expectedPayeeID := "0d0e928d-312a-4bcd-89c4-e02f40d1fe46"
	expectedPayeeName := "bla bla bla"
	expectedCategoryID := "f3cc4f55-312a-4bcd-89c4-db34379cb1dc"
	expectedCategoryName := "Yearly subscription"

	expected := &transaction.Scheduled{
		ID:              "56f4fc86-2ed7-4b3b-9116-7a214261b3cd",
		DateFirst:       expectedFirstAndLastDate,
		DateNext:        expectedFirstAndLastDate,
		Frequency:       transaction.FrequencyNever,
		Amount:          int64(-9000),
		Memo:            &expectedMemo,
		FlagColor:       &expectedFlagColor,
		AccountID:       "09eaca5e-312a-4bcd-89c4-828fb90638f2",
		AccountName:     "Bank Name",
		PayeeID:         &expectedPayeeID,
		PayeeName:       &expectedPayeeName,
		CategoryID:      &expectedCategoryID,
		CategoryName:    &expectedCategoryName,
		Deleted:         false,
		SubTransactions: []*transaction.ScheduledSubTransaction{},
	}
	assert.Equal(t, expected, stx)
}

func TestService_CreateTransaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	payloadDate, err := api.DateFromString("2018-11-13")
	assert.NoError(t, err)

	payloadPayeeID := "0d0e928d-312a-4bcd-89c4-e02f40d1fe46"
	payloadPayeeName := "bla bla bla"
	payloadCategoryID := "f3cc4f55-312a-4bcd-89c4-db34379cb1dc"
	payloadMemo := "nice memo"
	payloadFlagColor := transaction.FlagColorBlue

	payload := transaction.PayloadTransaction{
		AccountID:  "09eaca5e-312a-4bcd-89c4-828fb90638f2",
		Date:       payloadDate,
		Amount:     int64(-9000),
		Cleared:    transaction.ClearingStatusCleared,
		Approved:   true,
		PayeeID:    &payloadPayeeID,
		PayeeName:  &payloadPayeeName,
		CategoryID: &payloadCategoryID,
		Memo:       &payloadMemo,
		FlagColor:  &payloadFlagColor,
	}

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions"
	httpmock.RegisterResponder(http.MethodPost, url,
		func(req *http.Request) (*http.Response, error) {
			resModel := struct {
				Transaction *transaction.PayloadTransaction `json:"transaction"`
			}{}
			err := json.NewDecoder(req.Body).Decode(&resModel)
			assert.NoError(t, err)
			assert.Equal(t, &payload, resModel.Transaction)

			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transaction": {
      "id": "0f5b3f73-ded2-4dd7-8b01-c23022622cd6",
      "date": "2018-11-13",
      "amount": -9000,
      "memo": "nice memo",
      "cleared": "cleared",
      "approved": true,
      "flag_color": "blue",
      "account_id": "09eaca5e-312a-4bcd-89c4-828fb90638f2",
      "account_name": "Bank Name",
      "payee_id": "0d0e928d-312a-4bcd-89c4-e02f40d1fe46",
      "payee_name": "bla bla bla",
      "category_id": "f3cc4f55-312a-4bcd-89c4-db34379cb1dc",
      "category_name": "Groceries",
      "transfer_account_id": null,
      "import_id": null,
      "deleted": false,
      "subtransactions": []
    }
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	tx, err := client.Transaction().CreateTransaction("aa248caa-eed7-4575-a990-717386438d2c", payload)
	assert.NoError(t, err)

	expectedCategoryName := "Groceries"
	expectedTransaction := &transaction.Transaction{
		ID:              "0f5b3f73-ded2-4dd7-8b01-c23022622cd6",
		Date:            payload.Date,
		Amount:          payload.Amount,
		Memo:            payload.Memo,
		Cleared:         payload.Cleared,
		Approved:        payload.Approved,
		FlagColor:       payload.FlagColor,
		AccountID:       payload.AccountID,
		AccountName:     "Bank Name",
		PayeeID:         payload.PayeeID,
		PayeeName:       payload.PayeeName,
		CategoryID:      payload.CategoryID,
		CategoryName:    &expectedCategoryName,
		Deleted:         false,
		SubTransactions: []*transaction.SubTransaction{},
	}

	assert.Equal(t, expectedTransaction, tx)
}

func TestService_BulkCreateTransactions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	payloadDate, err := api.DateFromString("2018-11-13")
	assert.NoError(t, err)

	payloadPayeeID := "0d0e928d-312a-4bcd-89c4-e02f40d1fe46"
	payloadPayeeName := "bla bla bla"
	payloadCategoryID := "f3cc4f55-312a-4bcd-89c4-db34379cb1dc"
	payloadMemo := "nice memo"
	payloadFlagColor := transaction.FlagColorBlue
	payloadImportID := "asdfg"

	payload := []transaction.PayloadTransaction{
		{
			AccountID:  "09eaca5e-312a-4bcd-89c4-828fb90638f2",
			Date:       payloadDate,
			Amount:     int64(-9000),
			Cleared:    transaction.ClearingStatusCleared,
			Approved:   true,
			PayeeID:    &payloadPayeeID,
			PayeeName:  &payloadPayeeName,
			CategoryID: &payloadCategoryID,
			Memo:       &payloadMemo,
			FlagColor:  &payloadFlagColor,
			ImportID:   &payloadImportID,
		},
		{
			AccountID:  "09eaca5e-312a-4bcd-89c4-828fb90638f2",
			Date:       payloadDate,
			Amount:     int64(-9000),
			Cleared:    transaction.ClearingStatusCleared,
			Approved:   true,
			PayeeID:    &payloadPayeeID,
			PayeeName:  &payloadPayeeName,
			CategoryID: &payloadCategoryID,
			Memo:       &payloadMemo,
			FlagColor:  &payloadFlagColor,
			ImportID:   &payloadImportID,
		},
	}

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions/bulk"
	httpmock.RegisterResponder(http.MethodPost, url,
		func(req *http.Request) (*http.Response, error) {
			resModel := struct {
				Transactions []transaction.PayloadTransaction `json:"transactions"`
			}{}
			err := json.NewDecoder(req.Body).Decode(&resModel)
			assert.NoError(t, err)
			assert.Equal(t, payload, resModel.Transactions)

			res := httpmock.NewStringResponse(200, `{
  "data": {
    "bulk": {
      "transaction_ids": ["aaaaa321-eed7-4575-a990-717386438d2c"],
      "duplicate_import_ids": ["asdfg"]
    }
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	bulk, err := client.Transaction().BulkCreateTransactions(
		"aa248caa-eed7-4575-a990-717386438d2c",
		payload,
	)
	assert.NoError(t, err)

	expectedBunk := &transaction.Bulk{
		TransactionIDs:     []string{"aaaaa321-eed7-4575-a990-717386438d2c"},
		DuplicateImportIDs: []string{"asdfg"},
	}

	assert.Equal(t, expectedBunk, bulk)
}

func TestService_UpdateTransaction(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	payloadDate, err := api.DateFromString("2018-11-13")
	assert.NoError(t, err)

	payloadPayeeID := "0d0e928d-312a-4bcd-89c4-e02f40d1fe46"
	payloadPayeeName := "bla bla bla"
	payloadCategoryID := "f3cc4f55-312a-4bcd-89c4-db34379cb1dc"

	payload := transaction.PayloadTransaction{
		AccountID:  "09eaca5e-312a-4bcd-89c4-828fb90638f2",
		Date:       payloadDate,
		Amount:     int64(-100000),
		Cleared:    transaction.ClearingStatusCleared,
		Approved:   true,
		PayeeID:    &payloadPayeeID,
		PayeeName:  &payloadPayeeName,
		CategoryID: &payloadCategoryID,
		Memo:       nil,
		FlagColor:  nil,
	}

	url := "https://api.youneedabudget.com/v1/budgets/aa248caa-eed7-4575-a990-717386438d2c/transactions/0f5b3f73-ded2-4dd7-8b01-c23022622cd6"
	httpmock.RegisterResponder(http.MethodPut, url,
		func(req *http.Request) (*http.Response, error) {
			resModel := struct {
				Transaction *transaction.PayloadTransaction `json:"transaction"`
			}{}
			err := json.NewDecoder(req.Body).Decode(&resModel)
			assert.NoError(t, err)
			assert.Equal(t, &payload, resModel.Transaction)

			res := httpmock.NewStringResponse(200, `{
  "data": {
    "transaction": {
      "id": "0f5b3f73-ded2-4dd7-8b01-c23022622cd6",
      "date": "2018-11-13",
      "amount": -100000,
      "memo": null,
      "cleared": "cleared",
      "approved": true,
      "flag_color": null,
      "account_id": "09eaca5e-312a-4bcd-89c4-828fb90638f2",
      "account_name": "Bank Name",
      "payee_id": "0d0e928d-312a-4bcd-89c4-e02f40d1fe46",
      "payee_name": "bla bla bla",
      "category_id": "f3cc4f55-312a-4bcd-89c4-db34379cb1dc",
      "category_name": "Groceries",
      "transfer_account_id": null,
      "import_id": null,
      "deleted": false,
      "subtransactions": []
    }
	}
}
		`)
			res.Header.Add("X-Rate-Limit", "36/200")
			return res, nil
		},
	)

	client := ynab.NewClient("")
	tx, err := client.Transaction().UpdateTransaction(
		"aa248caa-eed7-4575-a990-717386438d2c",
		"0f5b3f73-ded2-4dd7-8b01-c23022622cd6",
		payload,
	)
	assert.NoError(t, err)

	expectedCategoryName := "Groceries"
	expectedTransaction := &transaction.Transaction{
		ID:              "0f5b3f73-ded2-4dd7-8b01-c23022622cd6",
		Date:            payload.Date,
		Amount:          payload.Amount,
		Cleared:         payload.Cleared,
		Approved:        payload.Approved,
		AccountID:       payload.AccountID,
		AccountName:     "Bank Name",
		PayeeID:         payload.PayeeID,
		PayeeName:       payload.PayeeName,
		CategoryID:      payload.CategoryID,
		CategoryName:    &expectedCategoryName,
		Deleted:         false,
		SubTransactions: []*transaction.SubTransaction{},
	}

	assert.Equal(t, expectedTransaction, tx)
}

func TestFilter_ToQuery(t *testing.T) {
	sinceDate, err := api.DateFromString("2020-02-02")
	assert.NoError(t, err)

	var zeroDate api.Date

	uncategorizedTransaction := transaction.StatusUncategorized
	unapprovedTransaction := transaction.StatusUnapproved

	table := []struct {
		Input  transaction.Filter
		Output string
	}{
		{
			Input:  transaction.Filter{Since: &sinceDate, Type: &unapprovedTransaction},
			Output: "since_date=2020-02-02&type=unapproved",
		},
		{
			Input:  transaction.Filter{Since: &sinceDate, Type: &uncategorizedTransaction},
			Output: "since_date=2020-02-02&type=uncategorized",
		},
		{
			Input:  transaction.Filter{Since: &sinceDate},
			Output: "since_date=2020-02-02",
		},
		{
			Input:  transaction.Filter{Type: &uncategorizedTransaction},
			Output: "type=uncategorized",
		},
		{
			Input:  transaction.Filter{Since: &zeroDate, Type: &uncategorizedTransaction},
			Output: "type=uncategorized",
		},
		{
			Input:  transaction.Filter{},
			Output: "",
		},
	}

	for _, test := range table {
		assert.Equal(t, test.Output, test.Input.ToQuery())
	}
}
