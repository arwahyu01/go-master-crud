package routes

import "github.com/gorilla/mux"

func RegisterRoutes(r *mux.Router) {
	WebRoutes(r)
	ApiRoutes(r)
}
