package ca

import (
	"net/http"

	"darvaza.org/acmefy/pkg/acme"
	"darvaza.org/acmefy/pkg/respond"
)

// ServeCertificate handles requests for the CA Certificate
func (ca *CA) ServeCertificate(rw http.ResponseWriter, req *http.Request) {
	res, err := certResponder.WithRequest(req)
	switch {
	case err != nil:
		_ = res.BadRequest(rw).Render(err)
	default:
		_ = res.OK(rw).Render(ca.caCert)
	}
}

var (
	certResponder = respond.Registry().
		Supports(acme.ContentTypePEMCertChain,
			acme.ContentTypePEM,
			acme.ContentTypeDERCA)
)
