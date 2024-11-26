// Package ca provides a basic Certificate Authority
package ca

import (
	"context"
	"crypto/x509"
	"errors"
	"io"

	"darvaza.org/x/tls/store/basic"
	"darvaza.org/x/tls/x509utils"
)

// CA is a basic Certificate Authority
type CA struct {
	cfg  Config
	pool *basic.Store

	caKey  x509utils.PrivateKey
	caCert []*x509.Certificate
}

// ECDSA tells if certificates should be ECDSA instead of RSA
func (ca *CA) ECDSA() bool {
	return ca.cfg.KeyAlgorithm == KeyAlgorithmECDSA
}

// ED25519 tells if certificates should be ED25519 instead of RSA
func (ca *CA) ED25519() bool {
	return ca.cfg.KeyAlgorithm == KeyAlgorithmED25519
}

// GenerateKey generates a new PrivateKey for a Server
func (ca *CA) GenerateKey() (x509utils.PrivateKey, error) {
	return ca.cfg.GenerateKey(false)
}

// WriteKey writes the CA's Private Key PEM encoded
func (ca *CA) WriteKey(w io.Writer) (int64, error) {
	return x509utils.WriteKey(w, ca.caKey)
}

// WriteCert writes the CA's Certificate PEM encoded
func (ca *CA) WriteCert(w io.Writer) (int64, error) {
	var count int64
	for _, crt := range ca.caCert {
		n, err := x509utils.WriteCert(w, crt)
		if n > 0 {
			count += n
		}
		if err != nil {
			return count, err
		}
	}
	return count, nil
}

func (ca *CA) validate() bool {
	if len(ca.caCert) == 0 || ca.caKey == nil || ca.caCert[0] == nil {
		// cert or key missing
		return false
	}

	// confirm the cert uses the given key
	cert := ca.caCert[0]
	ok := x509utils.ValidCertKeyPair(cert, ca.caKey)

	// TODO: validate ca.caCert chain
	return ok
}

func (ca *CA) prepare() (*CA, error) {
	s := &basic.Store{
		OnMissing: ca.onMissing,
	}

	n := len(ca.caCert)
	if n == 0 {
		return nil, errors.New("certificates not specified")
	}

	ctx := context.Background()

	// trust root
	_ = s.AddCACerts(ctx, ca.caCert[n-1])
	// add the rest as intermediate
	for _, c := range ca.caCert[:n-1] {
		_ = s.AddCert(ctx, c)
	}

	ca.pool = s
	return ca, nil
}
