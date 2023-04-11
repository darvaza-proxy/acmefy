package acme

// Identifier is an entity the account is
// authorized to represent
type Identifier struct {
	// Type indicates the type of identifier
	Type IdentifierType `json:"type"`
	// Value indicates the identifier itself
	Value string `json:"value"`
}

// IdentifierType indicates the type of identifier
// that is present in ACME authorization objects
// ([§9.7.7]).
//
// [§9.7.7]: https://www.rfc-editor.org/rfc/rfc8555#section-9.7.7
type IdentifierType string

const (
	// IdentifierDNS defined by [RFC8555] for fully
	// qualified domain names.
	//
	// [RFC8555]: https://www.rfc-editor.org/rfc/rfc8555
	IdentifierDNS IdentifierType = "dns"

	// IdentifierIP defined by [RFC8738] for authorizing
	// IPv4 and IPv6 addresses
	//
	// [RFC8738]: https://www.rfc-editor.org/rfc/rfc8738
	IdentifierIP IdentifierType = "ip"
)
