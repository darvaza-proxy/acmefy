package ca

import (
	"crypto"
	"crypto/rand"
	"crypto/x509"

	"darvaza.org/core"
	"darvaza.org/darvaza/shared/x509utils"
)

// revive:disable:flag-parameter

// NewKeyCertPair creates a new key+cert based on a given template,
// returning the resulting PEM encoded
func (ca *CA) NewKeyCertPair(clientAuth bool,
	tpl *x509.Certificate) (keyPEM, certPEM []byte, err error) {
	// revive:enable:flag-parameter
	if clientAuth {
		tpl.ExtKeyUsage = append(tpl.ExtKeyUsage, x509.ExtKeyUsageClientAuth)
	}

	key, err := ca.GenerateKey()
	if err != nil {
		err = core.Wrap(err, "failed to generate certificate key")
		return nil, nil, err
	}

	pub := key.Public()
	keyPEM = x509utils.EncodePKCS8PrivateKey(key)

	certPEM, err = ca.CreateCertificate(tpl, pub)
	if err != nil {
		return nil, nil, err
	}

	return keyPEM, certPEM, nil
}

// CreateCertificate signs a [x509.Certificate] returning the
// result PEM encoded
func (ca *CA) CreateCertificate(tpl *x509.Certificate,
	pub crypto.PublicKey) (certPEM []byte, err error) {
	//
	caCert := ca.caCert[0]
	certDER, err := x509.CreateCertificate(rand.Reader, tpl, caCert, pub, ca.caKey)
	if err != nil {
		err = core.Wrap(err, "failed to generate certificate")
		return nil, err
	}

	certPEM = x509utils.EncodeCertificate(certDER)

	return certPEM, nil
}
