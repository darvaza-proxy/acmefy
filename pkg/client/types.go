package client

import (
	"crypto/tls"
	"crypto/x509"
)

// Store is the interface we implement to use on tls.Config
type Store interface {
	GetCAPool() *x509.CertPool
	GetCertificate(hello *tls.ClientHelloInfo) (*tls.Certificate, error)
}
