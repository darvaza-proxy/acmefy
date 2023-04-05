package acme

import "github.com/go-jose/go-jose/v3/json"

// AccountStatus indicates the status of an account.
// See [§7.1.6]
//
// [§7.1.6]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.6
type AccountStatus string

const (
	// AccountValid indicates the account is valid
	AccountValid AccountStatus = "valid"
	// AccountDeactivated indicates client-initiated deactivation
	AccountDeactivated AccountStatus = "deactivated"
	// AccountRevoked indicates server-initiated deactivation
	AccountRevoked AccountStatus = "revoked"
)

// Account represents a set of metadata associated with an ACME
// account on the server ([§7.1.2])
//
// [§7.1.2]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.2)
type Account[EAB any] struct {
	// Status indicates the status of this account
	Status AccountStatus `json:"status"`

	// Contact is an optional array of URLs that the server can use to
	// contact the client for issues related to this account.  For
	// example, the server may wish to notify the client about
	// server-initiated revocation or certificate expiration.
	// For information on supported URL schemes, see
	// https://www.rfc-editor.org/rfc/rfc8555#section-7.3.
	Contact []string `json:"contact,omitempty" validate:"url"`

	// TermsOfServiceAgreed optionally indicates in a newAccount request
	// the client's agreement with the terms of service.
	// This field cannot be updated by the client.
	TermsOfServiceAgreed bool `json:"termsOfServiceAgreed,omitempty"`

	// ExternalAccountBinding is an optional field in a newAccount request
	// indicating approval by the holder of an existing
	// non-ACME account to bind that account to this ACME account.
	// This field is not updateable by the client, see
	// https://www.rfc-editor.org/rfc/rfc8555#section-7.3.4
	ExternalAccountBinding *json.RawMessage `json:"externalAccountBinding,omitempty"`

	// Orders is a URL from which a list of orders
	// submitted by this account can be fetched via a POST-as-GET
	// request, as described in
	// https://www.rfc-editor.org/rfc/rfc8555#section-7.1.2.1
	Orders string `json:"orders" validate:"url"`
}

// SetExternalAccountBinding stores ExternalAccountBinding
func (*Account[EAB]) SetExternalAccountBinding(_ EAB) {}

// GetExternalAccountBinding retrieves ExternalAccountBinding
func (*Account[EAB]) GetExternalAccountBinding() (EAB, bool) {
	var zero EAB
	return zero, false
}
