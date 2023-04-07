module darvaza.org/acmefy/pkg/server

go 1.19

replace (
	darvaza.org/core => ../../../core
	darvaza.org/darvaza/shared => ../../../darvaza/shared
	darvaza.org/darvaza/shared/web => ../../../darvaza/shared/web
	darvaza.org/resolve => ../../../resolve
)

replace (
	darvaza.org/acmefy => ../../
	darvaza.org/acmefy/pkg/respond => ../../pkg/respond
)
