package main

import (
	"log"

	"github.com/Terikyy/devops-lecture-project/auth-service/internal/server"
)

func main() {
	if err := server.Run(8081); err != nil {
		log.Fatal(err)
	}
}
