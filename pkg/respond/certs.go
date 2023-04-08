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
	register(acme.ContentTypePEMCertChain, RenderPEMCert)
	register(acme.ContentTypePEM, RenderPEMCert)
	register(acme.ContentTypeDERCA, RenderDERCert)
}

// RenderPEMCert renders *x509.Certificate and []*x509.Certificate
// PEM encoded
func RenderPEMCert(w io.Writer, v any) error {
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

// RenderDERCert renders the raw DER in a *x509.Certificate
// or the first entry on a []*x509.Certificate
func RenderDERCert(w io.Writer, v any) error {
	switch d := v.(type) {
	case []*x509.Certificate:
		return writerDERCert(w, d[0].Raw)
	case *x509.Certificate:
		return writerDERCert(w, d.Raw)
	default:
		return fmt.Errorf("invalid data type %T", v)
	}
}

func writerDERCert(w io.Writer, der []byte) error {
	buf := bytes.NewBuffer(der)
	_, err := buf.WriteTo(w)
	return err
}
