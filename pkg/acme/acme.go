// Package acme contains code shared between client and server of [RFC8555]
//
// [RFC8555]: https://www.rfc-editor.org/rfc/rfc8555
package acme

import (
	"fmt"
	"time"
)

// Timestamp renders in the format defined in [RFC3339]
//
// [RFC3399]: https://www.rfc-editor.org/rfc/rfc3339
type Timestamp time.Time

func (ts Timestamp) String() string {
	return time.Time(ts).UTC().Format(time.RFC3339)
}

// LinkRelation indicates the relationship a link has to the current document
type LinkRelation string

const (
	// LinkRelationUP is used with challenge resources to indicate
	// the authorization resource to which a challenge belongs.  It is also
	// used, with some media types, from certificate resources to indicate a
	// resource from which the client may fetch a chain of CA certificates
	// that could be used to validate the certificate in the original
	// resource. (See https://www.rfc-editor.org/rfc/rfc8555#section-7.1)
	LinkRelationUP LinkRelation = "up"

	// LinkRelationIndex is present on all resources other than the
	// directory and indicates the URL of the directory.
	// (See https://www.rfc-editor.org/rfc/rfc8555#section-7.1)
	LinkRelationIndex LinkRelation = "index"

	// LinkRelationNext is present to indicate there are more entries on
	// an orders list [§7.1.2.1]
	LinkRelationNext LinkRelation = "next"
)

// A Link is a Link HTTP header
type Link struct {
	URL      string
	Relation LinkRelation
}

func (l Link) String() string {
	return fmt.Sprintf("<%s>;rel=%s", l.URL, string(l.Relation))
}
