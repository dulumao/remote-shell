package tcp

import (
	"bufio"
	"net"
	"net/textproto"
	"os/exec"

	"github.com/andygeiss/remote-shell/business/handler"
)

// Handler ...
type Handler struct {
	con net.Conn
}

// NewHandler ...
func NewHandler(con net.Conn) handler.Handler {
	return &Handler{con}
}

// Handle ...
func (h *Handler) Handle() {
	in := bufio.NewReader(h.con)
	txt := textproto.NewReader(in)
	for {
		line, err := txt.ReadLine()
		if err != nil {
			break
		}
		cmd := exec.Command("/usr/bin/env", "sh", "-c", line)
		out, err := cmd.CombinedOutput()
		if err != nil {
			continue
		}
		h.con.Write(out)
	}
	h.con.Close()
}
