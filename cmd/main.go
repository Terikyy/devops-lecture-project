package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Terikyy/devops-lecture-project/internal/auth"
	"github.com/Terikyy/devops-lecture-project/internal/checkout"
	"github.com/Terikyy/devops-lecture-project/internal/products"
)

func main() {
	mux := http.NewServeMux()

	// Auth
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)

	// Products
	mux.HandleFunc("/products", products.ListHandler)
	mux.HandleFunc("/products/{id}", products.DetailHandler)

	// Checkout
	mux.HandleFunc("/checkout/placeorder", checkout.PlaceOrderHandler)

	port := 8080
	log.Printf("Server is running on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
