package payee

// Payee represents a payee for a budget
type Payee struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Deleted bool   `json:"deleted"`

	TransferAccountID *string `json:"transfer_account_id"`
}

// Location represents a location for a payee
// When a transaction is saved with an specified payee on the YNAB mobile apps,
// the GPS coordinates for that location are stored, so that the next time
// the user is in the same place (like the Grocery store) YNAB can pre-populate
// nearby payees for you. Locations will not be available for all payees.
type Location struct {
	ID      string `json:"id"`
	PayeeID string `json:"payee_id"`
	Deleted bool   `json:"deleted"`

	Latitude  *float64 `json:"latitude,string"`
	Longitude *float64 `json:"longitude,string"`
}
