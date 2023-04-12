package magic

import "darvaza.org/core"

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
