package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/axeloehrli/venta-de-campos-backend/api"
	"github.com/axeloehrli/venta-de-campos-backend/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	database, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	server, err := api.NewServer(config, database)

	if err != nil {
		log.Fatal("cannot create server: ", err)
	}
	server.Start()
}
