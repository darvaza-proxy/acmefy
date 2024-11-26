module darvaza.org/acmefy/pkg/ca

go 1.22

require (
	darvaza.org/acmefy v0.4.6
	darvaza.org/acmefy/pkg/respond v0.2.2
	darvaza.org/core v0.16.1
	darvaza.org/x/container v0.2.1 // indirect
	darvaza.org/x/fs v0.4.1 // indirect
	darvaza.org/x/tls v0.5.1
	darvaza.org/x/web v0.10.1 // indirect
)

require (
	github.com/go-jose/go-jose/v4 v4.0.5 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/zeebo/blake3 v0.2.4 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

replace (
	darvaza.org/acmefy => ../..
	darvaza.org/acmefy/pkg/respond => ../respond
)
