package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/prakharporwal/bank-server/api"
	"github.com/prakharporwal/bank-server/models/store"
	"github.com/prakharporwal/bank-server/services/klog"
)

const (
	serverAdd = "0.0.0.0:8080"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	// lambda.Start(handler)
	handler()
}

const RedisAddr = "13.233.195.130:6379"

func handler() {
	klog.Info("Hey I am creating a Bank Payment System! Will be fun to work on !", "hhh")
	//defer conn.Close()
	store := store.GetInstance()

	//store := db.GetInstance()
	server := api.NewServer(store)

	err := server.Start(serverAdd)
	defer server.Stop()
	if err != nil {
		klog.Error("cannot start server", err)
	}
}
