package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/researchbot/server/internal/config"
	"github.com/researchbot/server/internal/health"
)

type mockChecker struct {
	err error
}

func (m *mockChecker) Check(ctx context.Context) error {
	return m.err
}

func newTestConfig() *config.Config {
	return &config.Config{
		ServerPort: ":9090",
	}
}

func TestNew_ReturnsServer(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{}

	srv := New(cfg, checkers)

	if srv == nil {
		t.Fatal("expected non-nil server")
	}
	if srv.Server == nil {
		t.Fatal("expected non-nil http.Server")
	}
}

func TestNew_ServerAddr(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{}

	srv := New(cfg, checkers)

	if srv.Server.Addr != ":9090" {
		t.Errorf("expected addr ':9090', got '%s'", srv.Server.Addr)
	}
}

func TestNew_ServerTimeouts(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{}

	srv := New(cfg, checkers)

	if srv.Server.ReadTimeout != 10*time.Second {
		t.Errorf("expected read timeout 10s, got %v", srv.Server.ReadTimeout)
	}
	if srv.Server.WriteTimeout != 10*time.Second {
		t.Errorf("expected write timeout 10s, got %v", srv.Server.WriteTimeout)
	}
	if srv.Server.IdleTimeout != 60*time.Second {
		t.Errorf("expected idle timeout 60s, got %v", srv.Server.IdleTimeout)
	}
}

func TestNew_HealthEndpoint_ReturnsOK(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{
		"test": &mockChecker{err: nil},
	}

	srv := New(cfg, checkers)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Server.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp map[string]interface{}
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp["status"] != "ok" {
		t.Errorf("expected status 'ok', got '%v'", resp["status"])
	}
}

func TestNew_HealthEndpoint_ServiceDown(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{
		"test": &mockChecker{err: context.DeadlineExceeded},
	}

	srv := New(cfg, checkers)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Server.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}
}

func TestNew_UnregisteredRoute_Returns404(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{}

	srv := New(cfg, checkers)

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	rec := httptest.NewRecorder()

	srv.Server.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rec.Code)
	}
}

func TestNew_ContentTypeIsJSON(t *testing.T) {
	cfg := newTestConfig()
	checkers := map[string]health.Checker{
		"test": &mockChecker{err: nil},
	}

	srv := New(cfg, checkers)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	srv.Server.Handler.ServeHTTP(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", ct)
	}
}
