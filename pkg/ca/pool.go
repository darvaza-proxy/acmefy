package ca

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"io/fs"

	"darvaza.org/core"
	"darvaza.org/x/tls"
	"darvaza.org/x/tls/x509utils"
)

var (
	_ tls.Store = (*CA)(nil)
)

// GetCAPool generates a CertPool only including this CA
func (ca *CA) GetCAPool() *x509.CertPool {
	return ca.pool.GetCAPool()
}

// GetCertificate looks for the TLS certificate for a given chi.ServerName,
// and creates one if it doesn't
func (ca *CA) GetCertificate(chi *tls.ClientHelloInfo) (*tls.Certificate, error) {
	return ca.pool.GetCertificate(chi)
}

func (ca *CA) onMissing(ctx context.Context, name string) (*tls.Certificate, error) {
	cert, err := ca.newCertificate(ctx, ca.caKey, name)
	if err == nil {
		err = ca.pool.Put(ctx, cert)
	}
	return cert, err
}

func (ca *CA) newCertificate(_ context.Context,
	key x509utils.PrivateKey, name string) (*tls.Certificate, error) {
	//
	var crt *x509.Certificate
	var err error

	if key == nil {
		key, err = ca.GenerateKey()
		if err != nil {
			err = core.Wrap(err, "failed to generate key")
			return nil, err
		}
	}

	tpl := ca.cfg.Template.NewCertificateTemplate(name)
	certPEM, err := ca.CreateCertificate(tpl, key.Public())
	if err != nil {
		err = core.Wrap(err, "failed to generate certificate")
		return nil, err
	}

	_ = x509utils.ReadPEM(certPEM, func(_ fs.FS, _ string, block *pem.Block) bool {
		crt, err = x509utils.BlockToCertificate(block)
		return true
	})
	if err != nil {
		err = core.Wrap(err, "failed to decode signed certificate")
		return nil, err
	}

	return ca.bundle(key, crt)
}

func (ca *CA) bundle(key x509utils.PrivateKey, crt *x509.Certificate) (*tls.Certificate, error) {
	// build chain
	n := len(ca.caCert)
	certs := make([][]byte, 1, n+1)
	certs[0] = crt.Raw
	for _, c := range ca.caCert {
		certs = append(certs, c.Raw)
	}

	// assemble certificate
	cert := &tls.Certificate{
		Certificate: certs,
		PrivateKey:  key,
		Leaf:        crt,
	}
	return cert, nil
}
