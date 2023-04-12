package magic

import (
	"net/url"

	"darvaza.org/core"
	"darvaza.org/resolver"
	"darvaza.org/slog"
	"darvaza.org/slog/handlers/discard"
)

// New creates a new Magic client with the provided
// entrypoint for the ACME server
func New(entrypoint string) (*Magic, error) {
	u, err := url.Parse(entrypoint)
	if err != nil {
		err = core.Wrap(err, "acmefy/magic: invalid ACME URL")
		return nil, err
	}

	m := &Magic{
		entrypoint: u.String(),
		logger:     defaultLogger(),
		resolver:   defaultResolver(),
	}
	return m, nil
}

// SetLogger attaches an [slog.Logger] to the client.
// if nil a [discard.Logger] will be used.
func (m *Magic) SetLogger(l slog.Logger) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if l == nil {
		l = defaultLogger()
	}
	m.logger = l
}

func defaultLogger() slog.Logger {
	return discard.New()
}

// SetResolver specifies to resolver to be used when dialing
// the ACME server and the optional TLS-ALPN-01 remote.
// The Go standard net.Resolver with PreferGo set to true will
// be used unless one provided.
func (m *Magic) SetResolver(dns resolver.Resolver) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if dns == nil {
		dns = defaultResolver()
	}

	m.resolver = dns
	return nil
}

func defaultResolver() resolver.Resolver {
	return resolver.SystemResolver(true)
}

// SetTLSALPN01Remote sets the optional address intended
// to be used when forwarding TLS-ALPN-01 requests to
// a dedicated ACME server to handle the authorization
func (m *Magic) SetTLSALPN01Remote(remote string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	_, _, err := core.SplitHostPort(remote)
	if err != nil {
		return core.Wrap(err, "acmefy/magic: SetTLSALPN01Remote")
	}

	m.forwardTLSALPN01 = remote
	return nil
}
