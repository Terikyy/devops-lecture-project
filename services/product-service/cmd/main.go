package main

import (
	"log"

	"github.com/Terikyy/devops-lecture-project/product-service/internal/server"
)

func main() {
	if err := server.Run(8082); err != nil {
		log.Fatal(err)
	}
}
