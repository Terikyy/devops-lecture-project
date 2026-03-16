package server

import (
	"net/http"

	"github.com/Terikyy/devops-lecture-project/auth-service/internal/auth"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth/login", auth.LoginHandler)
	mux.HandleFunc("/auth/logout", auth.LogoutHandler)
}
