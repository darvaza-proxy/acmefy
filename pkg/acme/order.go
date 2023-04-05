package acme

// Orders is a list of orders belonging to the account.
// The server SHOULD include pending orders and SHOULD NOT
// include orders that are invalid in the array of URLs.
// The server MAY return an incomplete list, along with a Link
// header field with a "next" link relation indicating where
// further entries can be acquired.
type Orders struct {
	Orders []string `json:"orders,omitempty" validate:"url"`
}
