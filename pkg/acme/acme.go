// Package acme contains code shared between client and server of [RFC8555]
//
// [RFC8555]: https://www.rfc-editor.org/rfc/rfc8555
package acme

import "time"

// Timestamp renders in the format defined in [RFC3339]
//
// [RFC3399]: https://www.rfc-editor.org/rfc/rfc3339
type Timestamp time.Time

func (ts Timestamp) String() string {
	return time.Time(ts).UTC().Format(time.RFC3339)
}
