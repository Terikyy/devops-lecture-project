package products

import (
	"net/http"
	"strconv"

	"github.com/Terikyy/devops-lecture-project/product-service/pkg/httputil"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, GetProducts())
}

func DetailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		httputil.WriteError(w, http.StatusBadRequest, "Product ID has wrong format")
		return
	}

	product := FindProductByID(id)
	if product == nil {
		httputil.WriteError(w, http.StatusNotFound, "Product not found")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, product)
}
