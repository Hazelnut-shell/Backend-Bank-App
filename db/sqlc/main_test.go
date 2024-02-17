package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/yyht/simplebank/util"
)

// _ "github.com/lib/pq", to make go formatter keep this import.
// Otherwise, since we didn't directly use pq, this import would be deleted by go formatter

// global
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run()) // m.Run()  to start running the unit test, return exit code

}
