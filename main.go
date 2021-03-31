package simplebank

import (
	"database/sql"
	"github.com/ironsoul0/simplebank/api"
	db "github.com/ironsoul0/simplebank/db/sqlc"
	"log"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Can not connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	log.Fatalln(server.Start(":8080"))
}
