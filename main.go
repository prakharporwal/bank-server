package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/prakharporwal/bank-server/api"
	"github.com/prakharporwal/bank-server/db"
	"github.com/prakharporwal/bank-server/services/klog"
)

const (
	serverAdd = "0.0.0.0:8080"
	dbDriver  = "postgres"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

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
