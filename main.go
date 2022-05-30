package main

import (
	"database/sql"
	"github/jabutech/simplebank/api"
	db "github/jabutech/simplebank/db/sqlc"
	"github/jabutech/simplebank/util"
	"log"

	// Use driver sql for postgres
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	config, err := util.LoadConfig(".") // "." as location file app.env / in root folder
	// if error
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open connection to database
	conn, err := sql.Open(config.DBDriver, config.DBSource)
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
	err = server.Start(config.ServerAddress)
	// If error
	if err != nil {
		// Response error
		log.Fatal("cannot start server:", err)
	}
}
