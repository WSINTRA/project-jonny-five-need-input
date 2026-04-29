package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/researchbot/server/internal/config"
	"github.com/researchbot/server/internal/health"
	"github.com/researchbot/server/internal/neo4j"
	"github.com/researchbot/server/internal/server"
)

func main() {
	cfg := config.Load()

	client, err := neo4j.New(cfg.Neo4jURI, cfg.Neo4jUser, cfg.Neo4jPass)
	if err != nil {
		log.Fatalf("neo4j init failed: %v", err)
	}
	defer client.Close(context.Background())

	checkers := map[string]health.Checker{
		"neo4j": client,
	}

	srv := server.New(&cfg, checkers)

	go func() {
		log.Printf("server listening on %s", cfg.ServerAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	<-ctx.Done()
	log.Println("shutting down server...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(shutdownCtx)
}
