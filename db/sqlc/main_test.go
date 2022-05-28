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

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	// Open connection
	testDB, err = sql.Open(dbDriver, dbSource)
	// Check error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create new connection to db
	testQueries = New(testDB)

	// Run test
	os.Exit(m.Run())

}
