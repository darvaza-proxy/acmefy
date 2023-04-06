// Package respond contains encoding/decoding helpers
package respond

import "darvaza.org/darvaza/shared/web/respond"

var reg = respond.NewRegistry()

// Registry returns a reference to out Renderers registry
func Registry() *respond.Registry {
	return reg
}

const (
	// ContentTypePEM is the Content-Type for PEM encoded files
	ContentTypePEM = "application/x-pem-file"
	// ContentTypeDERCA is the Content-Type for DER encoded CA Certificates
	ContentTypeDERCA = "application/x-x509-ca-cert"
)
