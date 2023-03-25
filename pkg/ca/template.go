package ca

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"net"
	"net/mail"
	"net/url"
	"time"

	"github.com/darvaza-proxy/core"
)

// TemplateConfig describes the details to compose a new
// Certificate and CertificateRequest template
type TemplateConfig struct {
	O  string // O is the Subject.Organizaton
	OU string // OU is the Subject.OrganizationalUnit
	CN string // CN in the Subject.CommonName

	// Duration is how long the certificate will last
	Duration time.Duration
}

// revive:disable:cognitive-complexity

// NewCertificateTemplate creates a Certificate Template for a list of names.
// These names can be IP addresses, e-mail addresses, URIs or DNS names.
func (tc *TemplateConfig) NewCertificateTemplate(names ...string) *x509.Certificate {
	// revive:enable:cognitive-complexity
	duration := core.IIf(tc.Duration > 0, tc.Duration, DefaultCertificateDuration)
	from := time.Now()
	until := from.Add(duration)

	tpl := &x509.Certificate{
		SerialNumber: RandomSerialNumber(),
		Subject: pkix.Name{
			Organization:       []string{tc.O},
			OrganizationalUnit: []string{tc.OU},
			CommonName:         tc.CN,
		},

		NotBefore: from,
		NotAfter:  until,

		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
	}

	for _, h := range names {
		if ip, ok := tc.asIPAddress(h); ok {
			tpl.IPAddresses = append(tpl.IPAddresses, ip)
		} else if s, ok := tc.asEmailAddress(h); ok {
			tpl.EmailAddresses = append(tpl.EmailAddresses, s)
		} else if u, ok := tc.asURI(h); ok {
			tpl.URIs = append(tpl.URIs, u)
		} else {
			tpl.DNSNames = append(tpl.DNSNames, h)
		}
	}

	if len(tpl.IPAddresses) > 0 || len(tpl.DNSNames) > 0 || len(tpl.URIs) > 0 {
		tpl.ExtKeyUsage = append(tpl.ExtKeyUsage, x509.ExtKeyUsageServerAuth)
	}
	if len(tpl.EmailAddresses) > 0 {
		tpl.ExtKeyUsage = append(tpl.ExtKeyUsage, x509.ExtKeyUsageEmailProtection)
	}

	if len(names) > 0 {
		tpl.Subject.CommonName = names[0]
	}

	return tpl
}

func (*TemplateConfig) asIPAddress(h string) (net.IP, bool) {
	if ip, err := core.ParseNetIP(h); err == nil {
		return ip, true
	}
	return nil, false
}

func (*TemplateConfig) asEmailAddress(h string) (string, bool) {
	email, err := mail.ParseAddress(h)
	if err == nil && email.Address == h {
		return h, true
	}
	return "", false
}

func (*TemplateConfig) asURI(h string) (*url.URL, bool) {
	u, err := url.Parse(h)
	if err == nil && u.Scheme != "" && u.Host != "" {
		return u, true
	}
	return nil, false
}

// RandomSerialNumber generates a random serial number for a new
// Certificate
func RandomSerialNumber() *big.Int {
	sn, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		core.PanicWrap(err, "failed to generate serial number")
	}
	return sn
}

var serialNumberLimit = new(big.Int).Lsh(big.NewInt(1), 128)
