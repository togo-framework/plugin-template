// Package example holds the plugin's backend logic. Structure it like a mini app
// (services, models, queries) — the plugin is a self-contained unit.
package example

import (
	"net/http"

	"github.com/togo-framework/togo"
)

// Service is the plugin's backend service, built from the kernel.
type Service struct{ k *togo.Kernel }

// New constructs the service.
func New(k *togo.Kernel) *Service { return &Service{k: k} }

// Ping is a sample endpoint (GET /api/example/ping).
func (s *Service) Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"plugin":"example","status":"ok"}`))
}
