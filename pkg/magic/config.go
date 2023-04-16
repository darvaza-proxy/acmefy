package magic

import (
	"errors"

	"darvaza.org/slog"
)

type Config struct {
	// URL references the directory of the chosen ACME server
	URL string `validate:"url"`
	// ForwardALPN is the optional remote to dial to finish a
	// TLS-ALPN-01 challenge
	ForwardALPN string `validate:"tcp"`
	// Logger is an optional [slog.Logger] to use on the Magic
	// ACME client
	Logger slog.Logger

	Key   string
	Cert  string
	Roots string
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

	// CA Certs
	if cfg.Roots != "" {
		err = m.SetRoots(cfg.Roots)
		if err != nil {
			return nil, err
		}
	}

	// Own Key/Cert pair
	switch {
	case cfg.Key != "" && cfg.Cert != "":
		err = m.SetKeyCert(cfg.Key, cfg.Cert)
	case cfg.Key != "":
		err = m.SetKey(cfg.Key)
	case cfg.Cert != "":
		err = errors.New("cert without key")
	}

	if err == nil {
		err = m.Prepare()
	}

	if err != nil {
		return nil, err
	}
	return m, nil
}
