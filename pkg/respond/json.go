package respond

import (
	"io"

	"github.com/go-jose/go-jose/v4/json"
)

// RenderJSON renders JSON using github.com/go-jose/go-jose/v4/json
// for Marshalling
func RenderJSON(w io.Writer, v any) error {
	enc := json.NewEncoder(w)
	return enc.Encode(v)
}
