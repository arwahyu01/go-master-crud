package routes

import (
	AuthController "github.com/arwahyu01/go-jwt/app/controllers/auth"
	UserController "github.com/arwahyu01/go-jwt/app/controllers/user"
	"github.com/arwahyu01/go-jwt/app/middleware"
	"github.com/gorilla/mux"
)

func ApiRoutes(r *mux.Router) {
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", AuthController.Login).Methods("POST")
	api.HandleFunc("/register", AuthController.Register).Methods("POST")
	api.HandleFunc("/logout", AuthController.Logout).Methods("GET")

	// Routes yang butuh autentikasi
	api.HandleFunc("/user", middleware.AuthMiddleware(UserController.GetAllUser)).Methods("GET")
	api.HandleFunc("/profile", middleware.AuthMiddleware(UserController.GetProfile)).Methods("GET")
}
