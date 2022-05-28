package main

import (
	"database/sql"
	"github/jabutech/simplebank/api"
	db "github/jabutech/simplebank/db/sqlc"
	"log"

	// Use driver sql for postgres
	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	// Open connection to database
	conn, err := sql.Open(dbDriver, dbSource)
	// If error
	if err != nil {
		// Response error
		log.Fatal("cannot connect to db:", err)
	}

	// Set connection db is argument store
	store := db.NewStore(conn)
	// Ser store is argument api server
	server := api.NewServer(store)

	// Run server with server address set manual
	err = server.Start(serverAddress)
	// If error
	if err != nil {
		// Response error
		log.Fatal("cannot start server:", err)
	}
}
