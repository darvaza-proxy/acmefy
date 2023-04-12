package acme

// A Challenge object represents a server's offer to validate a
// client's possession of an identifier in a specific way.  Unlike the
// other objects listed above, there is not a single standard structure
// for a challenge object.  The contents of a challenge object depend on
// the validation method being used.  The general structure of challenge
// objects and an initial set of validation methods are described in [§8]
//
// [§8]: https://www.rfc-editor.org/rfc/rfc8555#section-8
type Challenge struct {
	// Challenge objects all contain the following basic fields:

	// Type is the type of challenge encoded in the object
	Type ChallengeType `json:"type"`
	// URL is where the response can be posted
	URL string `json:"url" validate:"url"`
	// Status indicates the status of this challenge
	Status ChallengeState `json:"status"`
	// Validated indicates the time at which the server validated
	// this challenge
	Validated Timestamp `json:"validated,omitempty"`
	// Error that occurred while the server was validating the challenge,
	// if any. A challenge object with an error MUST have status
	// equal to "invalid"
	Error *Problem `json:"error,omitempty"`

	// KeyAuthorization is a string that concatenates the token for the challenge
	// with a key fingerprint, separated by a '.' character
	//
	//	keyAuthorization = token || '.' || base64url(Thumbprint(accountKey))
	//
	// The "Thumbprint" step indicates the computation specified in
	// [RFC7638], using the SHA-256 digest [FIPS180-4].  As noted in
	// [RFC7518] any prepended zero octets in the fields of a JWK object
	// MUST be stripped before doing the computation.
	//
	// KeyAuthorization is used by http-01 and dns-01
	//
	// [RFC7638]: https://www.rfc-editor.org/rfc/rfc7638
	// [RFC7518]: https://www.rfc-editor.org/rfc/rfc7518
	// [FIPS180-4]: https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.180-4.pdf
	KeyAuthorization string `json:"keyAuthorization,omitempty"`

	// Token is a random value that uniquely identifies the challenge. This value
	// MUST have at least 128 bits of entropy.
	// Token is used by http-01.
	Token string `json:"token,omitempty" validate:"base64url"`
}

// ChallengeState indicates the status of a Challenge ([§7.1.6])
//
// [§7.1.6]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.6
type ChallengeState string

const (
	// ChallengePending is the initial state of a Challenge
	// once it's created
	ChallengePending ChallengeState = "pending"
	// ChallengeProcessing is the state after ChallengePending
	// when the client responds to the challenge ([§7.1.5])
	//
	// [§7.1.5]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.5
	ChallengeProcessing ChallengeState = "processing"
	// ChallengeValid indicates the validation was successful
	ChallengeValid ChallengeState = "valid"
	// ChallengeInvalid indicates there was an error on the validation
	ChallengeInvalid ChallengeState = "invalid"
)

// ChallengeType indicates the type of challenge encoded in the object
type ChallengeType string

const (
	// ChallengeTypeHTTP01 indicates an HTTP Challenge as decribed on [RFC8555 §8.3]
	//
	// [RFC8555 §8.3]: https://www.rfc-editor.org/rfc/rfc8555#section-8.3
	ChallengeTypeHTTP01 ChallengeType = "http-01"

	// ChallengeTypeDNS01 indicates a DNS Challenge as described on [RFC8555 §8.4]
	//
	// [RFC8555 §8.4]: https://www.rfc-editor.org/rfc/rfc8555#section-8.4
	ChallengeTypeDNS01 ChallengeType = "dns-01"

	// ChallengeTypeTLSALPN01 indicates a TLS ALPN Challenge as described on
	// [RFC8737 §3]
	//
	// [RFC8737 §3]: https://www.rfc-editor.org/rfc/rfc8737.html#section-3
	ChallengeTypeTLSALPN01 ChallengeType = "tls-alpn-01"
)
