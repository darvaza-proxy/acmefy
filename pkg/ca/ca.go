// Package ca provides a basic Certificate Authority
package ca

import (
	"github.com/darvaza-proxy/darvaza/shared/x509utils"
)

// CA is a basic Certificate Authority
type CA struct {
	cfg Config
}

// ECDSA tells if certificates should be ECDSA instead of RSA
func (ca *CA) ECDSA() bool {
	return ca.cfg.KeyAlgorithm == KeyAlgorithmECDSA
}

// GenerateKey generates a new PrivateKey for a Server
func (ca *CA) GenerateKey() (x509utils.PrivateKey, error) {
	return ca.cfg.GenerateKey(false)
}
