package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/axeloehrli/venta-de-campos-backend/util"
	_ "github.com/lib/pq"
)

// Create DB instance which will be used for package tests
var database *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

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
