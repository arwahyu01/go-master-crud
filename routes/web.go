package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

func WebRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("What are you looking for?"))
		if err != nil {
			return
		}
	}).Methods("GET")
}
