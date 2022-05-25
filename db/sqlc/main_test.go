package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// Use driver sql for postgres
	_ "github.com/lib/pq"
)

// Variable
const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

// Variable for tes query
var testQueries *Queries

func TestMain(m *testing.M) {
	// Open connection
	conn, err := sql.Open(dbDriver, dbSource)
	// Check error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create new connection to db
	testQueries = New(conn)

	// Run test
	os.Exit(m.Run())

}
