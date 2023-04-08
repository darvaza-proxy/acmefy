package acme

// Orders is a list of orders belonging to the account.
// The server SHOULD include pending orders and SHOULD NOT
// include orders that are invalid in the array of URLs.
// The server MAY return an incomplete list, along with a Link
// header field with a "next" link relation indicating where
// further entries can be acquired. ([§7.1.2.1])
//
// [§7.1.2.1]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.3
type Orders struct {
	Orders []string `json:"orders,omitempty" validate:"url"`
}

// An Order represents a client's request for a certificate
// and is used to track the progress of that order through to issuance.
// Thus, the object contains information about the requested
// certificate, the authorizations that the server requires the client
// to complete, and any certificates that have resulted from this order.
// ([§7.1.3])
//
// [§7.1.3]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.3
type Order struct {
	// Status indicates the status of this order
	Status OrderStatus `json:"status"`

	// Expires indicates the timestamp after which the server
	// will consider this order invalid
	Expires Timestamp `json:"expires,omitempty"`

	// Identifiers is an array of identifier objects that the
	// order pertains to
	Identifiers []Identifier `json:"identifiers"`

	// NotBefore is the requested notBefore value in the
	// certificate
	NotBefore Timestamp `json:"notBefore,omitempty"`

	// NotAfter is the requested notAfter value in the
	// certificate
	NotAfter Timestamp `json:"notAfter,omitempty"`

	// Error indicates a problem that occurred while processing
	// the order
	Error Problem `json:"error,omitempty"`

	// Authorizations is the array of authorizations
	// associated to this Order. For pending orders, the
	// authorizations that the client needs to complete before the
	// requested certificate can be issued (see Section 7.5), including
	// unexpired authorizations that the client has completed in the past
	// for identifiers specified in the order.  The authorizations
	// required are dictated by server policy; there may not be a 1:1
	// relationship between the order identifiers and the authorizations
	// required.  For final orders (in the "valid" or "invalid" state),
	// the authorizations that were completed.  Each entry is a URL from
	// which an authorization can be fetched with a POST-as-GET request.
	Authorizations []Authorization `json:"authorizations"`

	// Finalize is a URL that a CSR must be POSTed to once
	// all of the order's authorizations are satisfied to finalize the
	// order.  The result of a successful finalization will be the
	// population of the certificate URL for the order.
	Finalize string `json:"finalize" validate:"url"`

	// Certificate is a URL for the certificate that has
	// been issued in response to this order.
	Certificate string `json:"certificate" validate:"url"`
}

// OrderStatus indicates the status of an order ([§7.1.6])
//
// [§7.1.6]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.6
type OrderStatus string

const (
	// OrderPending is the initial state of an Order
	OrderPending OrderStatus = "pending"
	// OrderReady indicates all authorizations listed in the
	// Order are in AuthorizationValid state
	OrderReady OrderStatus = "ready"
	// OrderProcessing indicates the client has submitted the
	// CSR to the finalize URL and the CA begins the issuance process
	// for the certificate
	OrderProcessing OrderStatus = "processing"
	// OrderValid indicates the certificate has been issued
	OrderValid OrderStatus = "valid"
	// OrderInvalid indicates if it expires or one of its
	// authorizations enters a final state other than "valid"
	// ("expired", "revoked", or "deactivated")
	OrderInvalid OrderStatus = "invalid"
)
