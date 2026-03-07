package server

import "net/http"

// API routes:
//
//	GET    /api/platforms                   → list all platforms + status
//	POST   /api/platforms/{name}/connect    → connect platform
//	POST   /api/platforms/{name}/disconnect → disconnect platform
//	GET    /api/platforms/{name}/status     → single platform status
//
//	GET    /api/overlays                    → list overlays
//	POST   /api/overlays                    → create overlay
//	GET    /api/overlays/{name}             → get overlay metadata
//	PUT    /api/overlays/{name}             → update overlay
//	DELETE /api/overlays/{name}             → delete overlay
//	POST   /api/overlays/{name}/reload      → force hot-reload event to OBS
//	POST   /api/overlays/generate           → AI-generate overlay from prompt
//
//	GET    /api/config                      → get current config
//	PUT    /api/config                      → update config
//	GET    /api/stats                       → message counts, uptime, viewer totals
//
//	GET    /sse/chat                        → SSE stream of all platform messages
//	GET    /sse/logs                        → SSE stream of application logs
//	GET    /sse/events                      → SSE stream of internal bus events
//
//	GET    /overlays/{name}/                → serve static overlay files (for OBS)

func (s *Server) handleListPlatforms(w http.ResponseWriter, r *http.Request)      { panic("not implemented") }
func (s *Server) handleConnectPlatform(w http.ResponseWriter, r *http.Request)    { panic("not implemented") }
func (s *Server) handleDisconnectPlatform(w http.ResponseWriter, r *http.Request) { panic("not implemented") }
func (s *Server) handlePlatformStatus(w http.ResponseWriter, r *http.Request)     { panic("not implemented") }

func (s *Server) handleListOverlays(w http.ResponseWriter, r *http.Request)   { panic("not implemented") }
func (s *Server) handleCreateOverlay(w http.ResponseWriter, r *http.Request)  { panic("not implemented") }
func (s *Server) handleGetOverlay(w http.ResponseWriter, r *http.Request)     { panic("not implemented") }
func (s *Server) handleUpdateOverlay(w http.ResponseWriter, r *http.Request)  { panic("not implemented") }
func (s *Server) handleDeleteOverlay(w http.ResponseWriter, r *http.Request)  { panic("not implemented") }
func (s *Server) handleReloadOverlay(w http.ResponseWriter, r *http.Request)  { panic("not implemented") }
func (s *Server) handleGenerateOverlay(w http.ResponseWriter, r *http.Request) { panic("not implemented") }

func (s *Server) handleGetConfig(w http.ResponseWriter, r *http.Request)    { panic("not implemented") }
func (s *Server) handleUpdateConfig(w http.ResponseWriter, r *http.Request) { panic("not implemented") }
func (s *Server) handleStats(w http.ResponseWriter, r *http.Request)        { panic("not implemented") }

func (s *Server) handleSSEChat(w http.ResponseWriter, r *http.Request)   { panic("not implemented") }
func (s *Server) handleSSELogs(w http.ResponseWriter, r *http.Request)   { panic("not implemented") }
func (s *Server) handleSSEEvents(w http.ResponseWriter, r *http.Request) { panic("not implemented") }

func (s *Server) handleServeOverlay(w http.ResponseWriter, r *http.Request) { panic("not implemented") }
