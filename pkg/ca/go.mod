module darvaza.org/acmefy/pkg/ca

go 1.19

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)

require (
	darvaza.org/acmefy v0.4.1
	darvaza.org/acmefy/pkg/respond v0.0.1
	darvaza.org/core v0.9.2
	darvaza.org/darvaza/shared v0.5.1
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
)

require (
	darvaza.org/darvaza/shared/web v0.3.6 // indirect
	darvaza.org/slog v0.5.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.4 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
)
