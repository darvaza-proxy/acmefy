package acme

import "strings"

// ErrorType is a type of error defined by [§6.7]
//
// [§6.7]: https://www.rfc-editor.org/rfc/rfc8555#section-6.7
type ErrorType string

const (
	// ErrorAccountDoesNotExist indicates the request specified an
	// account that does not exist
	ErrorAccountDoesNotExist ErrorType = "accountDoesNotExist"

	// ErrorAlreadyRevoked indicates the request specified a
	// certificate to be revoked that has already been revoked
	ErrorAlreadyRevoked ErrorType = "alreadyRevoked"

	// ErrorBadCSR indicates the CSR is unacceptable
	// (e.g., due to a short key)
	ErrorBadCSR ErrorType = "badCSR"

	// ErrorBadNonce indicates the client sent an unacceptable
	// anti-replay nonce
	ErrorBadNonce ErrorType = "badNonce"

	// ErrorBadPublicKey indicates the JWS was signed by a public key
	// the server does not support
	ErrorBadPublicKey ErrorType = "badPublicKey"

	// ErrorBadRevocationReason indicates the revocation reason provided
	// is not allowed by the server
	ErrorBadRevocationReason ErrorType = "badRevocationReason"

	// ErrorBadSignatureAlgorithm indicates the JWS was signed with an
	// algorithm the server does not support
	ErrorBadSignatureAlgorithm ErrorType = "badSignatureAlgorithm"

	// ErrorCAA indicates Certification Authority Authorization (CAA)
	// records forbid the CA from issuing a certificate
	ErrorCAA ErrorType = "caa"

	// ErrorCompound indicates specific error conditions are indicated
	// in the "subproblems" array
	ErrorCompound ErrorType = "compound"

	// ErrorConnection indicates the server could not connect to
	// the validation target
	ErrorConnection ErrorType = "connection"

	// ErrorDNS indicates there was a problem with a DNS query during
	// identifier validation
	ErrorDNS ErrorType = "dns"

	// ErrorExternalAccountRequired indicates the request must include
	// a value for the "externalAccountBinding" field
	ErrorExternalAccountRequired ErrorType = "externalAccountRequired"

	// ErrorIncorrectResponse indicates the response received didn't
	// match the challenge's requirements
	ErrorIncorrectResponse ErrorType = "incorrectResponse"

	// ErrorInvalidContact indicates a contact URL for an account was invalid
	ErrorInvalidContact ErrorType = "invalidContact"

	// ErrorMalformed indicates the request message was malformed
	ErrorMalformed ErrorType = "malformed"

	// ErrorOrderNotReady indicates the request attempted to finalize an order
	// that is not ready to be finalized
	ErrorOrderNotReady ErrorType = "orderNotReady"

	// ErrorRateLimited indicates the request exceeds a rate limit
	ErrorRateLimited ErrorType = "rateLimited"

	// ErrorRejectedIdentifier indicates the server will not issue certificates
	// for the identifier
	ErrorRejectedIdentifier ErrorType = "rejectedIdentifier"

	// ErrorServerInternal indicates the server experienced an internal error
	ErrorServerInternal ErrorType = "serverInternal"

	// ErrorTLS the server received a TLS error during validation
	ErrorTLS ErrorType = "tls"

	// ErrorUnauthorized indicates the client lacks sufficient authorization
	ErrorUnauthorized ErrorType = "unauthorized"

	// ErrorUnsupportedContact indicates a contact URL for an account used
	// an unsupported protocol scheme
	ErrorUnsupportedContact ErrorType = "unsupportedContact"

	// ErrorUnsupportedIdentifier indicates an identifier is of an unsupported
	// type
	ErrorUnsupportedIdentifier ErrorType = "unsupportedIdentifier"

	// ErrorUserActionRequired indicates visit the "instance" URL and take actions
	// specified there
	ErrorUserActionRequired ErrorType = "userActionRequired"
)

// String renders an error type
func (e ErrorType) String() string {
	s := string(e)
	if strings.ContainsRune(s, ':') {
		return s
	}
	return "urn:ietf:params:acme:error:" + s
}
