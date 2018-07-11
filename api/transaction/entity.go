package transaction

import "bmvs.io/ynab/api"

// Transaction represents a transaction
type Transaction struct {
	ID        string   `json:"id"`
	Date      api.Date `json:"date"`
	Amount    int64    `json:"amount"`
	Cleared   Status   `json:"cleared"`
	Approved  bool     `json:"approved"`
	AccountID string   `json:"account_id"`
	Deleted   bool     `json:"deleted"`

	Memo              *string    `json:"memo"`
	FlagColor         *FlagColor `json:"flag_color"`
	PayeeID           *string    `json:"payee_id"`
	CategoryID        *string    `json:"category_id"`
	TransferAccountID *string    `json:"transfer_account_id"`
	ImportID          *string    `json:"import_id"`
}

// Sub represents a sub transaction
type Sub struct {
	ID            string `json:"id"`
	TransactionID string `json:"transaction_id"`
	Amount        int64  `json:"amount"`
	Deleted       bool   `json:"deleted"`

	Memo              *string `json:"memo"`
	PayeeID           *string `json:"payee_id"`
	CategoryID        *string `json:"category_id"`
	TransferAccountID *string `json:"transfer_account_id"`
}

// Scheduled represents a scheduled transaction
type Scheduled struct {
	ID        string    `json:"id"`
	DateFirst api.Date  `json:"date_first"`
	DateNext  api.Date  `json:"date_next"`
	Frequency Frequency `json:"frequency"`
	Amount    int64     `json:"amount"`
	AccountID string    `json:"account_id"`
	Deleted   bool      `json:"deleted"`

	Memo              *string    `json:"memo"`
	FlagColor         *FlagColor `json:"flag_color"`
	PayeeID           *string    `json:"payee_id"`
	CategoryID        *string    `json:"category_id"`
	TransferAccountID *string    `json:"transfer_account_id"`
}

// ScheduledSub represents a scheduled sub transaction
type ScheduledSub struct {
	ID                     string `json:"id"`
	ScheduledTransactionID string `json:"scheduled_transaction_id"`
	Amount                 int64  `json:"amount"`
	Deleted                bool   `json:"deleted"`

	Memo              *string `json:"memo"`
	PayeeID           *string `json:"payee_id"`
	CategoryID        *string `json:"category_id"`
	TransferAccountID *string `json:"transfer_account_id"`
}
