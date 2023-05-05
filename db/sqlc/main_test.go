package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5433/doubtful_bank?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("main_test.go -> Could not connect to db: ", err)
	}
	defer connection.Close()

	testQueries = New(connection)

	os.Exit(m.Run())
}
