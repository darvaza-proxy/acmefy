// Package ca provides a basic Certificate Authority
package ca

import (
	"crypto/x509"

	"github.com/darvaza-proxy/darvaza/shared/x509utils"
)

// CA is a basic Certificate Authority
type CA struct {
	cfg Config

	caKey  x509utils.PrivateKey
	caCert []*x509.Certificate
}

// ECDSA tells if certificates should be ECDSA instead of RSA
func (ca *CA) ECDSA() bool {
	return ca.cfg.KeyAlgorithm == KeyAlgorithmECDSA
}

// GenerateKey generates a new PrivateKey for a Server
func (ca *CA) GenerateKey() (x509utils.PrivateKey, error) {
	return ca.cfg.GenerateKey(false)
}

func (ca *CA) validate() bool {
	if len(ca.caCert) == 0 || ca.caKey == nil || ca.caCert[0] == nil {
		// cert or key missing
		return false
	}

	// confirm the cert uses the given key
	cert := ca.caCert[0]
	pub, ok := cert.PublicKey.(x509utils.PublicKey)
	if ok {
		ok = pub.Equal(ca.caKey.Public())
	}

	// TODO: validate ca.caCert chain
	return ok
}
