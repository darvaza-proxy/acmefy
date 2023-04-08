package acme

// An Authorization object represents a server's authorization for
// an account to represent an identifier. ([§7.1.4])
//
// [§7.1.4]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.4
type Authorization struct {
	// Identifier that the account is authorized to represent
	Identifier Identifier `json:"identifier"`
	// Status of the authorization
	Status AuthorizationStatus `json:"status"`
	// Expires indicates the timestamp after which the server
	// will consider this authorization invalid.
	// This field is REQUIRED for objects with "valid" status
	Expires Timestamp `json:"expires,omitempty"`
	// Challenges is an array of Challenge objects.
	//
	// - For pending authorizations, the challenges that the client
	//   can fulfill in order to prove possession of the identifier.
	// - For valid authorizations, the challenge that was validated.
	// - For invalid authorizations, the challenge that was attempted
	//   and failed.
	//
	// Each array entry is an object with parameters required to
	// validate the challenge.  A client should attempt to fulfill
	// one of these challenges, and a server should consider any one
	// of the challenges sufficient to make the authorization valid.
	Challenges []Challenge `json:"challenges"`

	// Wildcard indicates the authorization was created as a result
	// of a newOrder request containing a DNS identifier with a value
	// that was a wildcard domain name.
	Wildcard bool `json:"wildcard,omitempty"`
}

// AuthorizationStatus represents the status of an Authorization
//
// [§7.1.6]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.6
type AuthorizationStatus string

const (
	// AuthorizationPending is the initial state after creation
	AuthorizationPending AuthorizationStatus = "pending"
	// AuthorizationValid indicates at least one challenge transitioned
	// to ChallengeValid
	AuthorizationValid AuthorizationStatus = "valid"
	// AuthorizationInvalid indicates the client attempted to fulfill a
	// challenge and failed, or there was an error while the authorization
	// was still pending.
	AuthorizationInvalid AuthorizationStatus = "invalid"
	// AuthorizationExpired indicates a valid Authorization has expired
	AuthorizationExpired AuthorizationStatus = "expired"
	// AuthorizationDeactivated indicates a valid Authorization has been
	// deactivated by the client
	AuthorizationDeactivated AuthorizationStatus = "deactivated"
	// AuthorizationRevoked indicates a valid Authorization has been
	// revoked by the server
	AuthorizationRevoked AuthorizationStatus = "revoked"
)
