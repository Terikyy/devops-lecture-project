package checkout

import (
	"net/http"
	"strings"

	"github.com/Terikyy/devops-lecture-project/checkout-service/pkg/httputil"
	"github.com/Terikyy/devops-lecture-project/checkout-service/pkg/jwt"
)

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		httputil.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		httputil.WriteError(w, http.StatusUnauthorized, "Authorization header must use Bearer scheme")
		return
	}

	tokenString := jwt.ExtractBearerToken(authHeader)
	if !jwt.VerifyToken(tokenString) {
		httputil.WriteError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	httputil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Order placed successfully"})
}
