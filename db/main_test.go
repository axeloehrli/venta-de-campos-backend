package db

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "secret"
	dbname   = "venta-de-campos-backend"
)

// Create DB instance which will be used for package tests
var database *sql.DB

var err error

func TestMain(m *testing.M) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
