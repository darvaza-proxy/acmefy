module darvaza.org/acmefy/pkg/ca

go 1.21

require (
	darvaza.org/acmefy v0.4.4
	darvaza.org/acmefy/pkg/respond v0.1.2
	darvaza.org/core v0.14.8
	darvaza.org/darvaza/shared v0.6.0
	darvaza.org/slog v0.5.10 // indirect
	darvaza.org/slog/handlers/discard v0.4.13 // indirect
	darvaza.org/x/fs v0.3.2 // indirect
	darvaza.org/x/web v0.8.6 // indirect
)

require golang.org/x/exp v0.0.0-20230905200255-921286631fa9

require (
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)
