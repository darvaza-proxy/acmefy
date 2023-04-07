module darvaza.org/acmefy/pkg/magic

go 1.19

replace (
	darvaza.org/core => ../../../core
	darvaza.org/darvaza/shared => ../../../darvaza/shared
	darvaza.org/darvaza/shared/web => ../../../darvaza/shared/web
	darvaza.org/resolver => ../../../resolver
)

require (
	darvaza.org/core v0.9.7
	darvaza.org/resolver v0.5.4
	darvaza.org/slog v0.5.3
	darvaza.org/slog/handlers/discard v0.4.5
)

require (
	darvaza.org/cache v0.2.2 // indirect
	github.com/miekg/dns v1.1.55 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	golang.org/x/tools v0.12.0 // indirect
)
