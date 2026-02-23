package server

import (
	"fmt"
	"log"
	"net/http"
)

func Run(port int) error {
	mux := http.NewServeMux()
	SetupRoutes(mux)

	log.Printf("Server is running on port %d...\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
