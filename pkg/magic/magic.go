// Package magic provides a ACME client implementation
package magic

import (
	"net/url"
	"sync"

	"darvaza.org/core"
	"darvaza.org/resolver"
	"darvaza.org/slog"
)

// Magic is an ACME client that attempts to
// simplify usage for dns and ip authorizations
type Magic struct {
	mu sync.Mutex

	entrypoint string
	logger     slog.Logger
	resolver   resolver.Resolver

	forwardTLSALPN01 string
}

// New creates a new Magic client with the provided
// entrypoint for the ACME server
func New(entrypoint string) (*Magic, error) {
	if _, err := url.Parse(entrypoint); err != nil {
		err = core.Wrap(err, "acmefy/magic: invalid ACME URL")
		return nil, err
	}

	m := &Magic{
		entrypoint: entrypoint,
		logger:     defaultLogger(),
		resolver:   defaultResolver(),
	}
	return m, nil
}
