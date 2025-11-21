package main

import (
	"log"

	"github.com/EYOB123695/ecommerce-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8081", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
