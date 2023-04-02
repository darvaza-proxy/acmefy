package ca

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"errors"
	"time"

	"darvaza.org/core"
	"darvaza.org/darvaza/shared/x509utils"
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
	// KeyAlgorithmED25519 uses ED25519 for either CA or server
	KeyAlgorithmED25519
)

const (
	Day  = 24 * time.Hour // Day is 24 hours
	Year = 365 * Day      // Year is 365 days

	// DefaultCertificateDuration is 90 days
	DefaultCertificateDuration = 90 * Day
	// DefaultCADuration is 10 years
	DefaultCADuration = 10 * Year
)

// Config describes how the [CA] will operate
type Config struct {
	// KeyAlgorithm specifies the algorithm to use when
	// generating a PrivateKey. Defaults to RSA.
	KeyAlgorithm KeyAlgorithm

	// Template is used to create new certificates
	Template TemplateConfig
}

// GenerateKey generates a new PrivateKey
func (cfg *Config) GenerateKey(rootCA bool) (x509utils.PrivateKey, error) {
	switch cfg.KeyAlgorithm {
	case KeyAlgorithmECDSA:
		return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case KeyAlgorithmED25519:
		_, key, err := ed25519.GenerateKey(rand.Reader)
		return key, err
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

	return ca.prepare()
}

// NewCA generates a new self-signed CA using the provided TemplateConfig
func (cfg Config) NewCA(tc *TemplateConfig) (*CA, error) {
	// New CA Private Key
	key, err := cfg.GenerateKey(true)
	if err != nil {
		err = core.Wrap(err, "failed to generate certificate key")
		return nil, err
	}

	// SubjectKeyID
	pub := key.Public()
	skid, err := x509utils.SubjectPublicKeySHA1(pub)
	if err != nil {
		return nil, err
	}

	// New CA Certificate
	tpl := tc.NewCATemplate(skid[:])

	certDER, err := x509.CreateCertificate(rand.Reader, tpl, tpl, pub, key)
	if err != nil {
		err = core.Wrap(err, "failed to generate certificate")
		return nil, err
	}

	cert, err := x509.ParseCertificate(certDER)
	if err != nil {
		err = core.Wrap(err, "failed to parse generated certificate")
		return nil, err
	}

	ca := &CA{
		cfg:    cfg,
		caKey:  key,
		caCert: []*x509.Certificate{cert},
	}

	if !ca.validate() {
		err = errors.New("incompatible pair created")
		return nil, err
	}

	return ca.prepare()
}
