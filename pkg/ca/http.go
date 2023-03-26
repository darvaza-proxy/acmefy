package ca

import (
	"bytes"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"

	"darvaza.org/darvaza/shared/web/respond"
	"darvaza.org/darvaza/shared/x509utils"
)

const (
	// ContentTypePEM is the Content-Type for PEM encoded files
	ContentTypePEM = "application/x-pem-file"
	// ContentTypeDERCA is the Content-Type for DER encoded CA Certificates
	ContentTypeDERCA = "application/x-x509-ca-cert"
)

// ServeCertificate handles requests for the CA Certificate
func (ca *CA) ServeCertificate(rw http.ResponseWriter, req *http.Request) {
	res, err := certResponder.WithRequest(req)
	switch {
	case err != nil:
		res.BadRequest(rw).Render(err)
	default:
		res.OK(rw).Render(ca.caCert)
	}
}

var (
	certResponder *respond.Responder
)

func init() {
	reg := respond.NewRegistry()
	reg.Register("", respond.NewRenderer(ContentTypePEM, renderPEM))
	reg.Register("", respond.NewRenderer(ContentTypeDERCA, renderDERCA))

	certResponder = reg.Supports(ContentTypePEM, ContentTypeDERCA)
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
