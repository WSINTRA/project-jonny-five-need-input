package config

import (
	"os"
)

type Config struct {
	ServerAddr string
	ServerPort string
	Neo4jURI   string
	Neo4jUser  string
	Neo4jPass  string
}

func Load() Config {
	return Config{
		ServerAddr: getEnv("SERVER_ADDR", "localhost:9090"),
		ServerPort: getEnv("SERVER_PORT", ":9090"),
		Neo4jURI:   getEnv("NEO4J_URI", "bolt://localhost:7687"),
		Neo4jUser:  getEnv("NEO4J_USER", "neo4j"),
		Neo4jPass:  getEnv("NEO4J_PASSWORD", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
