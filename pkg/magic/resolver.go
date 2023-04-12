package magic

import "darvaza.org/resolver"

func defaultResolver() resolver.Resolver {
	return resolver.SystemResolver(true)
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
