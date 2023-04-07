module darvaza.org/acmefy/pkg/magic

go 1.19

replace (
	darvaza.org/core => ../../../core
	darvaza.org/darvaza/shared => ../../../darvaza/shared
	darvaza.org/darvaza/shared/web => ../../../darvaza/shared/web
	darvaza.org/resolver => ../../../resolver
)

replace (
	darvaza.org/acmefy => ../../
	darvaza.org/acmefy/pkg/acme => ../acme
	darvaza.org/acmefy/pkg/ca => ../ca
	darvaza.org/acmefy/pkg/client => ../client
	darvaza.org/acmefy/pkg/respond => ../respond
)

require (
	darvaza.org/core v0.9.2
	darvaza.org/resolver v0.5.0
	darvaza.org/slog v0.5.1
	darvaza.org/slog/handlers/discard v0.4.0
)

require (
	github.com/miekg/dns v1.1.54 // indirect
	golang.org/x/mod v0.10.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	golang.org/x/tools v0.9.3 // indirect
)
