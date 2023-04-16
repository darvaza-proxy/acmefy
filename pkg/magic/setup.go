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

// SetRoots adds accepted Root CAs. It can't be called
// after Prepare()
func (m *Magic) SetRoots(certs ...string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.pb == nil {
		return errSetupFinished
	}
	return m.pb.AddCACerts(certs...)
}

// SetKey sets the Private Key, and optionally more certificates.
// It can't be called after Prepare()
func (m *Magic) SetKey(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.pb == nil {
		return errSetupFinished
	}

	return m.pb.Add(key)
}

// SetKeyCert sets the Private Key and a list of certificates
func (m *Magic) SetKeyCert(key string, certs ...string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.pb == nil {
		return errSetupFinished
	}

	if err := m.pb.Add(key); err != nil {
		return err
	}

	return m.pb.AddCerts(certs...)
}

// Prepare composes the initial store using the provided
// keys and certs, or initialises it as self-signed
// if the relevant data wasn't provided
func (*Magic) Prepare() error {
	return nil
}
