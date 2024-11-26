module darvaza.org/acmefy/pkg/ca

go 1.21

require (
	darvaza.org/acmefy v0.4.6
	darvaza.org/acmefy/pkg/respond v0.2.2
	darvaza.org/core v0.15.2
	darvaza.org/x/fs v0.3.3 // indirect
	darvaza.org/x/tls v0.3.1
	darvaza.org/x/web v0.9.0 // indirect
)

require (
	darvaza.org/x/container v0.1.1 // indirect
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/zeebo/blake3 v0.2.4 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)
