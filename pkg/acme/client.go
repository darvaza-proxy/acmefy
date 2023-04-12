package acme

// Client wraps knowledge about a particular ACME server
// which we need for composing messages
type Client interface {
	Directory() Directory
}
