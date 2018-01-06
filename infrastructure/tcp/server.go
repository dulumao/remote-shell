package tcp

import (
	"net"

	"github.com/andygeiss/remote-shell/business/server"
)

// Server ...
type Server struct {
	address string
}

// NewServer ...
func NewServer(addr string) server.Server {
	return &Server{addr}
}

// Listen ...
func (s *Server) Listen() {
	srv, err := net.Listen("tcp", s.address)
	if err != nil {
		return
	}
	for {
		con, err := srv.Accept()
		if err != nil {
			continue
		}
		go func() {
			h := NewHandler(con)
			h.Handle()
		}()
	}
}
