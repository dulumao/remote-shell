package server

// Server can listen to a specific address or interface.
type Server interface {
	Listen()
}
