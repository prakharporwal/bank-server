package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/prakharporwal/bank-server/api"
	"github.com/prakharporwal/bank-server/db"
	"github.com/prakharporwal/bank-server/services"
)

const (
	serverAdd = "0.0.0.0:8080"
	dbSource  = "postgresql://admin:password@localhost:5432/default_db?sslmode=disable"
	dbDriver  = "postgres"
)

var klog services.Logger

func main() {
	// lambda.Start(handler)
	handler()
}

func handler() {
	klog.Info("Hey I am creating a Bank Payment System! Will be fun to work on !", "hhh")

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		klog.Info("connect to db failed !", err)
		panic(err)
	}
	defer conn.Close()

	klog.Debug("\nSuccessfully connected to database!\n")

	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(serverAdd)
	if err != nil {
		klog.Error("cannot sta rt server", err)
	}
}
