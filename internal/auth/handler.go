package auth

import (
	"net/http"

	"github.com/Terikyy/devops-lecture-project/pkg/httputil"
	"github.com/Terikyy/devops-lecture-project/pkg/jwt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	// For simplicity, we'll use a hardcoded username and password
	// This should be replaced with a proper authentication mechanism
	if username == "user" && password == "pass" {
		token, err := jwt.CreateToken(username)
		if err != nil {
			httputil.WriteError(w, http.StatusInternalServerError, "Error generating the token")
			return
		}
		httputil.WriteJSON(w, http.StatusOK, map[string]string{"token": token})
	} else {
		httputil.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		httputil.WriteError(w, http.StatusMethodNotAllowed, "Method Not Allowed")
		return
	}

	// In this simple example, we'll just return a success message
	httputil.WriteJSON(w, http.StatusOK, map[string]string{"message": "Logout successful"})
}
