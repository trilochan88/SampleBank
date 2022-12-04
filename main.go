package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/trilochan88/simplebank/api"
	db "github.com/trilochan88/simplebank/db/sqlc"
	"github.com/trilochan88/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Printf("Config: driver: %s, source: %s,address: %s", config.DBDriver, config.DBSource, config.ServerAddress)
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannont start server: ", err)
	}

}
