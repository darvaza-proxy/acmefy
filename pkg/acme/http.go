package acme

var (
	// ContentTypeDERCA is the Content-Type used when rendering
	// binary DER encoded CA Certificates
	ContentTypeDERCA = "application/x-x509-ca-cert"

	// ContentTypeJOSE is the ContentType used when rendering
	// JWS as flattened JSON
	ContentTypeJOSE = "application/jose+json"

	// ContentTypePEMCertChain is the ContentType used when rendering
	// PEM Certificate chains
	ContentTypePEMCertChain = "application/pem-certificate-chain"

	// ContentTypeProblem is the ContentType used when rendering a Problem
	ContentTypeProblem = "application/problem+json"
)
