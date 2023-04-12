package acme

import "strings"

// HTTP01ResourcePath returns the path at which the
func (c Challenge) HTTP01ResourcePath() string {
	if c.Token != "" {
		var s = []string{
			HTTP01ResourcePrefix,
			c.Token,
		}

		return strings.Join(s, "/")
	}
	return ""
}

// HTTP01ResourcePrefix is the fixed prefix where the server
// should find the KeyAuthorization for a challenge
// based on the Token
const HTTP01ResourcePrefix = "/.well-known/acme-challenge"
