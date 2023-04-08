package respond

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io"

	"darvaza.org/darvaza/shared/x509utils"

	"darvaza.org/acmefy/pkg/acme"
)

func init() {
	register(acme.ContentTypePEMCertChain, renderPEM)
	register(acme.ContentTypePEM, renderPEM)
	register(acme.ContentTypeDERCA, renderDERCA)
}

func renderPEM(w io.Writer, v any) error {
	switch d := v.(type) {
	case []*x509.Certificate:
		return writePEMCert(w, d...)
	case *x509.Certificate:
		return writePEMCert(w, d)
	default:
		return fmt.Errorf("invalid data type %T", v)
	}
}

func writePEMCert(w io.Writer, certs ...*x509.Certificate) error {
	for _, crt := range certs {
		_, err := x509utils.WriteCert(w, crt)
		if err != nil {
			return err
		}
	}
	return nil
}

func renderDERCA(w io.Writer, v any) error {
	switch d := v.(type) {
	case []*x509.Certificate:
		return writerDERCACert(w, d[0].Raw)
	case *x509.Certificate:
		return writerDERCACert(w, d.Raw)
	default:
		return fmt.Errorf("invalid data type %T", v)
	}
}

func writerDERCACert(w io.Writer, der []byte) error {
	buf := bytes.NewBuffer(der)
	_, err := buf.WriteTo(w)
	return err
}
