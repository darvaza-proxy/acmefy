module darvaza.org/acmefy/pkg/ca

go 1.19

replace (
	darvaza.org/core => ../../../core
	darvaza.org/darvaza/shared => ../../../darvaza/shared
	darvaza.org/darvaza/shared/web => ../../../darvaza/shared/web
	darvaza.org/resolver => ../../../resolver
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)

require (
	darvaza.org/acmefy v0.4.3
	darvaza.org/acmefy/pkg/respond v0.1.1
	darvaza.org/core v0.9.7
	darvaza.org/darvaza/shared v0.5.8
	golang.org/x/exp v0.0.0-20230713183714-613f0c0eb8a1
)

require (
	darvaza.org/darvaza/shared/web v0.3.9 // indirect
	darvaza.org/slog v0.5.3 // indirect
	darvaza.org/slog/handlers/discard v0.4.5 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
)
