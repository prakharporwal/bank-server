package main

import (
	_ "github.com/lib/pq"
	"github.com/prakharporwal/bank-server/api"
	"github.com/prakharporwal/bank-server/db"
	"github.com/prakharporwal/bank-server/services/klog"
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
	klog.Info("Hey I am creating a Bank Payment System! Will be fun to work on !", "hhh")

	//defer conn.Close()

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered in f", r)
	//	}
	//}()

	store := db.GetInstance()

	//store := db.GetInstance()
	server := api.NewServer(store)

	err := server.Start(serverAdd)
	defer server.Stop()
	if err != nil {
		klog.Error("cannot start server", err)
	}

}
