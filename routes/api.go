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

	user := api.PathPrefix("/user").Subrouter()
	user.Use(middleware.AuthMiddleware)
	user.HandleFunc("", UserController.GetAllUser).Methods("GET")
	user.HandleFunc("/{id}", UserController.UpdateUser).Methods("PUT")
	user.HandleFunc("/profile", UserController.GetProfile).Methods("GET")
}
