package server

import (
	"net/http"

	"github.com/Terikyy/devops-lecture-project/internal/auth"
	"github.com/Terikyy/devops-lecture-project/internal/checkout"
	"github.com/Terikyy/devops-lecture-project/internal/products"
)

func SetupRoutes(mux *http.ServeMux) {
	// Auth
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)

	// Products
	mux.HandleFunc("/products", products.ListHandler)
	mux.HandleFunc("/products/{id}", products.DetailHandler)

	// Checkout
	mux.HandleFunc("/checkout/placeorder", checkout.PlaceOrderHandler)
}
