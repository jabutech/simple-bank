package db

import (
	"database/sql"
	"github/jabutech/simplebank/util"
	"log"
	"os"
	"testing"

	// Use driver sql for postgres
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load config
	config, err := util.LoadConfig("../..") // "../.." as location file app.env / in root folder
	// if error
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	// Open connection
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	// Check error
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	// Create new connection to db
	testQueries = New(testDB)

	// Run test
	os.Exit(m.Run())

}
