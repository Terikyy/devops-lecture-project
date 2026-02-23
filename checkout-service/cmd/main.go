package main

import (
	"log"

	"github.com/Terikyy/devops-lecture-project/checkout-service/internal/server"
)

func main() {
	if err := server.Run(8083); err != nil {
		log.Fatal(err)
	}
}
