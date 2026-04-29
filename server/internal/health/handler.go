package health

import (
	"context"
	"encoding/json"
	"net/http"
)

type Checker interface {
	Check(ctx context.Context) error
}

type Response struct {
	Status   string            `json:"status"`
	Services map[string]string `json:"services"`
}

type Handler struct {
	checkers map[string]Checker
}

func NewHandler(checkers map[string]Checker) *Handler {
	return &Handler{checkers: checkers}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	services := make(map[string]string, len(h.checkers))
	overall := "ok"

	for name, checker := range h.checkers {
		if err := checker.Check(ctx); err != nil {
			services[name] = "disconnected"
			overall = "degraded"
		} else {
			services[name] = "connected"
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if overall != "ok" {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	json.NewEncoder(w).Encode(Response{
		Status:   overall,
		Services: services,
	})
}
