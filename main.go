package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/prakharporwal/back-server/api"
	"github.com/prakharporwal/back-server/db"
)

const (
	serverAdd = "0.0.0.0:8080"
	dbSource  = "postgresql://admin:password@localhost:5432/default_db?sslmode=disable"
	dbDriver  = "postgres"
)

func main() {
	// lambda.Start(handler)
	handler()
}

func handler() {
	fmt.Println("Hey I am creating a Bank Payment System! Will be fun to work on !")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("connect to db failed !", err)
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("\nSuccessfully connected to database!\n")

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAdd)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
