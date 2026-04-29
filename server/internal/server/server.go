package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/researchbot/server/internal/config"
	"github.com/researchbot/server/internal/health"
)

type Server struct {
	*http.Server
}

func New(cfg *config.Config, checkers map[string]health.Checker) *Server {
	h := health.NewHandler(checkers)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	registerRoutes(r, h)

	return &Server{&http.Server{
		Addr:         cfg.ServerPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}}
}

func registerRoutes(r chi.Router, h *health.Handler) {
	r.Get("/health", h.ServeHTTP)
}
