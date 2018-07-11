package payee

// Payee represents a payee
type Payee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`

	TransferAccountID *string `json:"transfer_account_id"`
}

// Location represents the location of a payee
type Location struct {
	ID      string `json:"id"`
	PayeeID string `json:"payee_id"`
	Deleted bool   `json:"deleted"`

	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
