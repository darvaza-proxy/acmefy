package acme

// Directory is the index of ACME resources on a server ([§7.1.1]).
//
// There is no constraint on the URL of the directory except that it
// should be different from the other ACME server resources' URLs, and
// that it should not clash with other services.  For instance:
//
//   - a host that functions as both an ACME and a Web server may want to
//     keep the root path "/" for an HTML "front page" and place the ACME
//     directory under the path "/acme".
//
//   - a host that only functions as an ACME server could place the
//     directory under the path "/".
//
// If the ACME server does not implement pre-authorization ([§7.4.1])
// it MUST omit the "newAuthz" field of the directory.
//
// [§7.1.1]: https://www.rfc-editor.org/rfc/rfc8555#section-7.1.1)
// [§7.4.1]: https://www.rfc-editor.org/rfc/rfc8555#section-7.4.1)
type Directory struct {
	NewNonce   string         `json:"newNonce"`
	NewAccount string         `json:"newAccount"`
	NewOrder   string         `json:"newOrder"`
	NewAuthz   string         `json:"newAuthz,omitempty"`
	RevokeCert string         `json:"revokeCert"`
	KeyChange  string         `json:"keyChange"`
	Meta       *DirectoryMeta `json:"meta,omitempty"`
}

// DirectoryMeta is an optional field of [Directory] that
// provides additional information about the ACME server
type DirectoryMeta struct {
	// TermsOfService is a URL identifying the current terms of service.
	TermsOfService string `json:"termsOfService,omitempty"`

	// Website is an HTTP or HTTPS URL locating a website
	// providing more information about the ACME server.
	Website string `json:"website,omitempty"`

	// CAAIdentities indicates hostnames that the
	// ACME server recognizes as referring to itself for the purposes of
	// CAA record validation as defined in [RFC6844].  Each string MUST
	// represent the same sequence of ASCII code points that the server
	// will expect to see as the "Issuer Domain Name" in a CAA issue or
	// issuewild property tag.  This allows clients to determine the
	// correct issuer domain name to use when configuring CAA records.
	CAAIdentities []string `json:"caaIdentities,omitempty"`

	// ExternalAccountRequired indicates if the CA requires that all
	// newAccount requests include an "externalAccountBinding" field
	// associating the new account with an external account.
	ExternalAccountRequired bool `json:"externalAccountRequired,omitempty"`
}
