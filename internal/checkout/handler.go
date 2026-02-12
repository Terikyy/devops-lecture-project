package checkout

import (
	"net/http"
	"strings"

	"github.com/Terikyy/devops-lecture-project/pkg/jwt"
)

func PlaceOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Missing Authorization header"}`))
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Authorization header must use Bearer scheme"}`))
		return
	}

	tokenString := jwt.ExtractBearerToken(authHeader)
	if !jwt.VerifyToken(tokenString) {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(`{"error":"Invalid token"}`))
		return
	}

	w.Write([]byte(`{"message":"Order placed successfully"}`))
}
