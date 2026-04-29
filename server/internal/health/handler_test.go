package health

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockChecker struct {
	err error
}

func (m *mockChecker) Check(ctx context.Context) error {
	return m.err
}

func TestHandler_AllServicesHealthy(t *testing.T) {
	h := NewHandler(map[string]Checker{
		"neo4j": &mockChecker{err: nil},
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var resp Response
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("expected status 'ok', got '%s'", resp.Status)
	}

	if resp.Services["neo4j"] != "connected" {
		t.Errorf("expected neo4j 'connected', got '%s'", resp.Services["neo4j"])
	}
}

func TestHandler_ServiceDown(t *testing.T) {
	h := NewHandler(map[string]Checker{
		"neo4j": &mockChecker{err: errors.New("connection refused")},
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var resp Response
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "degraded" {
		t.Errorf("expected status 'degraded', got '%s'", resp.Status)
	}

	if resp.Services["neo4j"] != "disconnected" {
		t.Errorf("expected neo4j 'disconnected', got '%s'", resp.Services["neo4j"])
	}
}

func TestHandler_MultipleServices_OneDown(t *testing.T) {
	h := NewHandler(map[string]Checker{
		"neo4j":  &mockChecker{err: nil},
		"redis":  &mockChecker{err: errors.New("timeout")},
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("expected status 503, got %d", rec.Code)
	}

	var resp Response
	if err := json.NewDecoder(rec.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Status != "degraded" {
		t.Errorf("expected status 'degraded', got '%s'", resp.Status)
	}

	if resp.Services["neo4j"] != "connected" {
		t.Errorf("expected neo4j 'connected', got '%s'", resp.Services["neo4j"])
	}

	if resp.Services["redis"] != "disconnected" {
		t.Errorf("expected redis 'disconnected', got '%s'", resp.Services["redis"])
	}
}

func TestHandler_ContentType(t *testing.T) {
	h := NewHandler(map[string]Checker{
		"neo4j": &mockChecker{err: nil},
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	ct := rec.Header().Get("Content-Type")
	if ct != "application/json" {
		t.Errorf("expected Content-Type 'application/json', got '%s'", ct)
	}
}
