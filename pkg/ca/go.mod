module darvaza.org/acmefy/pkg/ca

go 1.19

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)

require (
	darvaza.org/acmefy v0.4.3
	darvaza.org/acmefy/pkg/respond v0.1.1
	darvaza.org/core v0.9.7
	darvaza.org/darvaza/shared v0.5.8
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9
)

require (
	darvaza.org/darvaza/shared/web v0.3.9 // indirect
	darvaza.org/slog v0.5.3 // indirect
	darvaza.org/slog/handlers/discard v0.4.5 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
