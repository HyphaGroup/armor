package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/HyphaGroup/armor/server/internal/api"
	"github.com/HyphaGroup/armor/server/internal/db"
	"github.com/HyphaGroup/armor/server/internal/validator"
)

func main() {
	port := flag.String("port", "8080", "Server port")
	dbPath := flag.String("db", "./armor.db", "Database path")
	schemasDir := flag.String("schemas", "../schemas", "Path to JSON schemas directory")
	flag.Parse()

	if envPort := os.Getenv("ARMOR_PORT"); envPort != "" {
		*port = envPort
	}
	if envDB := os.Getenv("ARMOR_DB_PATH"); envDB != "" {
		*dbPath = envDB
	}
	if envSchemas := os.Getenv("ARMOR_SCHEMAS_DIR"); envSchemas != "" {
		*schemasDir = envSchemas
	}

	absDBPath, err := filepath.Abs(*dbPath)
	if err != nil {
		log.Fatalf("Failed to resolve database path: %v", err)
	}

	absSchemasDir, err := filepath.Abs(*schemasDir)
	if err != nil {
		log.Fatalf("Failed to resolve schemas directory: %v", err)
	}

	log.Printf("Opening database at %s", absDBPath)
	database, err := db.Open(absDBPath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer database.Close()

	log.Printf("Loading schemas from %s", absSchemasDir)
	val, err := validator.New(absSchemasDir)
	if err != nil {
		log.Fatalf("Failed to load schemas: %v", err)
	}

	server := api.NewServer(database, val)

	addr := ":" + *port
	log.Printf("Starting server on http://localhost%s", addr)
	if err := http.ListenAndServe(addr, server); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
