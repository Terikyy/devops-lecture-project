package server

import (
	"net/http"

	"github.com/Terikyy/devops-lecture-project/product-service/internal/products"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/products", products.ListHandler)
	mux.HandleFunc("/products/{id}", products.DetailHandler)
}
