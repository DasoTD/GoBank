package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	// "github.com/dasotd/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	DBDriver ="postgres"
	DBSource = "postgresql://root:secret@localhost:5432/test?sslmode=disable"
)

func TestMain(m *testing.M) {
	// config, err := util.LoadConfig("../..")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }

	testDB, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}