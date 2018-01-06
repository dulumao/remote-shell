package remote

import "github.com/andygeiss/remote-shell/business/server"

// Shell ...
type Shell struct {
	isInstallCalled bool
	localServer     server.Server
}

// NewShell initializes the data structure and returns its address.
func NewShell(srv server.Server) *Shell {
	return &Shell{false, srv}
}

// Install the shell into the current process.
func (s *Shell) Install() {
	// Install only once
	if !s.isInstallCalled {
		s.localServer.Listen()
		s.isInstallCalled = true
	}
}
