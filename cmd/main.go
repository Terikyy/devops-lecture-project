package main

import (
	"log"

	"github.com/Terikyy/devops-lecture-project/internal/server"
)

func main() {
	if err := server.Run(8080); err != nil {
		log.Fatal(err)
	}
}
