package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	_ "github.com/stretchr/testify/require"
)

var testqueries *Queries
var testdb *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root1:ananya@localhost:5432/SimpleBank?sslmode=disable"
)

//var testqueries *Queries

func TestMain(m *testing.M) {
	var err error
	testdb, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testqueries = New(testdb)

	os.Exit(m.Run())
}
