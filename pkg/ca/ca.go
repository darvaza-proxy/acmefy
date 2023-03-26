// Package ca provides a basic Certificate Authority
package ca

import (
	"crypto/x509"
	"io"

	"github.com/darvaza-proxy/core"
	"github.com/darvaza-proxy/darvaza/shared/storage/certpool"
	"github.com/darvaza-proxy/darvaza/shared/storage/simple"
	"github.com/darvaza-proxy/darvaza/shared/x509utils"
)

// CA is a basic Certificate Authority
type CA struct {
	cfg  Config
	pool *simple.Store

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
	pub, ok := cert.PublicKey.(x509utils.PublicKey)
	if ok {
		ok = pub.Equal(ca.caKey.Public())
	}

	// TODO: validate ca.caCert chain
	return ok
}

func (ca *CA) prepare() (*CA, error) {
	var pb certpool.PoolBuffer

	pb.AddKey("", ca.caKey)
	for _, c := range ca.caCert {
		pb.AddCert("", c)
	}

	pool, err := simple.NewFromBuffer(&pb, nil)
	if err != nil {
		err = core.Wrap(err, "failed to create certificate store")
		return nil, err
	}

	// done
	ca.pool = pool
	return ca, nil
}
