package main

import (
	"fmt"
	"log"

	"github.com/prakharporwal/back-server/api"
)

const (
	serverAdd = "0.0.0.0:8080"
)

func main() {
	// lambda.Start(handler)
	handler()
}

func handler() {
	fmt.Println("Hey I am creating a Bank Payment System! Will be fun to work on !")

	server := api.NewServer()

	err := server.Start(serverAdd)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
