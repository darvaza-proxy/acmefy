package ca

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"

	"github.com/darvaza-proxy/core"
	"github.com/darvaza-proxy/darvaza/shared/x509utils"
)

// KeyAlgorithm specifies the algorithm to use when generating a Private Key
type KeyAlgorithm int

const (
	// KeyAlgorithmUnspecified is treated as if it was KeyAlgorithmRSA
	KeyAlgorithmUnspecified KeyAlgorithm = iota
	// KeyAlgorithmRSA uses RSA3072 for the CA, and 2048 for servers
	KeyAlgorithmRSA
	// KeyAlgorithmECDSA uses ECDSA 256 for either CA or server
	KeyAlgorithmECDSA
)

// Config describes how the [CA] will operate
type Config struct {
	// KeyAlgorithm specifies the algorithm to use when
	// generating a PrivateKey. Defaults to RSA.
	KeyAlgorithm KeyAlgorithm
}

// GenerateKey generates a new PrivateKey
func (cfg *Config) GenerateKey(rootCA bool) (x509utils.PrivateKey, error) {
	switch cfg.KeyAlgorithm {
	case KeyAlgorithmECDSA:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	default:
		bits := core.IIf(rootCA, 3072, 2048)
		return rsa.GenerateKey(rand.Reader, bits)
	}
}
