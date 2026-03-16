package server

import (
	"net/http"

	"github.com/Terikyy/devops-lecture-project/checkout-service/internal/checkout"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/checkout/placeorder", checkout.PlaceOrderHandler)
}
