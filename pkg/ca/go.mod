module darvaza.org/acmefy/pkg/ca

go 1.21

require (
	darvaza.org/acmefy v0.4.4
	darvaza.org/acmefy/pkg/respond v0.1.2
	darvaza.org/core v0.9.9
	darvaza.org/darvaza/shared v0.5.9
	darvaza.org/darvaza/shared/web v0.3.10 // indirect
	darvaza.org/slog v0.5.4 // indirect
	darvaza.org/slog/handlers/discard v0.4.6 // indirect
)

require golang.org/x/exp v0.0.0-20230905200255-921286631fa9

require (
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/zeebo/blake3 v0.2.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)
