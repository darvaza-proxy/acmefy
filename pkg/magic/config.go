package magic

import (
	"time"

	"darvaza.org/slog"
)

type Config struct {
	// Logger is an optional [slog.Logger] to use on the Magic
	// ACME client
	Logger slog.Logger

	// URL references the directory of the chosen ACME server
	URL string `validate:"url"`

	// ForwardALPN is the optional remote to dial to finish a
	// TLS-ALPN-01 challenge
	ForwardALPN string `validate:"tcp"`

	// Timeout indicates how long to wait for Getter before
	// issuing its own. Negative means no timeout, and zero
	// milliseconds falls back to a 1s default
	Timeout time.Duration

	Keys  []string
	Certs []string
	Roots []string
}

func (cfg Config) New() (*Magic, error) {
	// Client
	m, err := New(cfg.URL)
	if err != nil {
		return nil, err
	}

	// cfg.Logger != nil by SetDefaults()
	m.SetLogger(cfg.Logger)

	// TLS-ALPN-01 Forward
	if cfg.ForwardALPN != "" {
		err = m.SetTLSALPN01Remote(cfg.ForwardALPN)
		if err != nil {
			return nil, err
		}
	}

	if err := cfg.applyFiles(m); err != nil {
		return nil, err
	}

	if err := m.Prepare(); err != nil {
		return nil, err
	}

	return m, nil
}

func (cfg Config) applyFiles(m *Magic) error {
	if err := applyStrings(m.AddKey, cfg.Keys...); err != nil {
		return err
	}

	if err := applyStrings(m.AddCACert, cfg.Roots...); err != nil {
		return err
	}

	return applyStrings(m.AddCert, cfg.Certs...)
}

func applyStrings(cb func(string) error, data ...string) error {
	for _, s := range data {
		if err := cb(s); err != nil {
			return err
		}
	}
	return nil
}
