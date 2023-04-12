// Package magic provides a ACME client implementation
package magic

import (
	"sync"

	"darvaza.org/resolver"
	"darvaza.org/slog"
)

// Magic is an ACME client that attempts to
// simplify usage for dns and ip authorizations
type Magic struct {
	mu       sync.Mutex
	logger   slog.Logger
	resolver resolver.Resolver

	entrypoint       string
	forwardTLSALPN01 string
}
