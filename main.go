package main

import (
	"database/sql"
	"github.com/ironsoul0/simplebank/api"
	db "github.com/ironsoul0/simplebank/db/sqlc"
	"github.com/ironsoul0/simplebank/util"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Can not read config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(config, store)

	log.Fatalln(server.Start(config.ServerAddress))
}
