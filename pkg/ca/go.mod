module darvaza.org/acmefy/pkg/ca

go 1.22

require (
	darvaza.org/acmefy v0.4.6
	darvaza.org/acmefy/pkg/respond v0.2.2
	darvaza.org/core v0.16.0
	darvaza.org/darvaza/shared v0.7.0
	darvaza.org/slog v0.6.0 // indirect
	darvaza.org/slog/handlers/discard v0.5.0 // indirect
	darvaza.org/x/fs v0.4.0 // indirect
	darvaza.org/x/tls v0.5.0 // indirect
	darvaza.org/x/web v0.10.0 // indirect
)

require (
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/zeebo/blake3 v0.2.4 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)
