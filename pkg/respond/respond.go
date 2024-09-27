// Package respond contains encoding/decoding helpers
package respond

import "darvaza.org/x/web/respond"

var reg = respond.NewRegistry()

// Registry returns a reference to out Renderers registry
func Registry() *respond.Registry {
	return reg
}

func register(ct string, h respond.RenderFunc) {
	reg.Register(ct, respond.NewRenderer(ct, h))
}
