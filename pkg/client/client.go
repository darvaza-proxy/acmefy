// Package client provides generic interfaces and helpers to implement ACME clients
package client

import "darvaza.org/acmefy/pkg/acme"

var (
	_ acme.Client = (*Client)(nil)
)

// Client represents the connection to a particular ACME server
type Client struct {
	directory acme.Directory
}

// Directory tells the different endpoints provided by a
// particular ACME Server
func (c Client) Directory() acme.Directory {
	return c.directory
}
