package ca

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"time"

	"github.com/darvaza-proxy/core"
	"github.com/darvaza-proxy/darvaza/shared/x509utils"
	"golang.org/x/exp/slices"
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

const (
	Day = 24 * time.Hour // Day is 24 hours

	// DefaultCertificateDuration is 90 days
	DefaultCertificateDuration = 90 * Day
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

// LoadCA creates a new CA using the given key and certs chain
func (cfg Config) LoadCA(key x509utils.PrivateKey, certs []*x509.Certificate) (*CA, error) {
	ca := &CA{
		cfg:    cfg,
		caKey:  key,
		caCert: slices.Clone(certs),
	}

	if !ca.validate() {
		err := errors.New("incompatible pair provided")
		return nil, err
	}

	return ca, nil
}
