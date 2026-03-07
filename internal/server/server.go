package server

import (
	"context"
	"net/http"

	"github.com/pacdouglas/stream-pac/internal/event"
	"github.com/pacdouglas/stream-pac/internal/hub"
	"github.com/pacdouglas/stream-pac/internal/overlay"
)

// Server is the HTTP + SSE server that exposes the hub, overlays, and events
// to TUI clients, OBS Browser Sources, and debug tooling (curl, browser).
//
// Since the TUI is a pure HTTP client, every TUI action maps to an API call.
// This also means the entire app can be operated without the TUI (e.g. curl).
type Server struct {
	addr    string
	hub     *hub.Hub
	bus     event.Bus
	manager overlay.Manager
	mux     *http.ServeMux
	srv     *http.Server
}

// New creates a Server. Call Start to begin serving.
func New(addr string, h *hub.Hub, bus event.Bus, mgr overlay.Manager) *Server {
	panic("not implemented")
}

// Start registers all routes and begins serving HTTP.
// Blocks until ctx is cancelled or a fatal error occurs.
func (s *Server) Start(ctx context.Context) error {
	panic("not implemented")
}

// Stop gracefully shuts down the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	panic("not implemented")
}

// routes registers all API, SSE, and static overlay routes on s.mux.
func (s *Server) routes() {
	panic("not implemented")
}
