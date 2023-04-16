package magic

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"darvaza.org/acmefy/pkg/client"
	"darvaza.org/darvaza/shared/storage"
	"darvaza.org/darvaza/shared/x509utils"
)

var (
	_ client.Store  = (*Magic)(nil)
	_ storage.Store = (*Magic)(nil)
)

func (m *Magic) GetCAPool() *x509.CertPool {
	return m.pool.GetCAPool()
}

func (m *Magic) GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return m.pool.GetCertificateWithCallback(hello, m.fetchCertificate)
}

func (*Magic) fetchCertificate(_ context.Context,
	_ x509utils.PrivateKey, _ string) (*tls.Certificate, error) {
	return nil, nil
}
