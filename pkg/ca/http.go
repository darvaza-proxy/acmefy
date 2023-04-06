package ca

import (
	"net/http"

	"darvaza.org/acmefy/pkg/respond"
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
	certResponder = respond.Registry().
		Supports(respond.ContentTypePEM,
			respond.ContentTypeDERCA)
)
