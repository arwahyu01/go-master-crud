package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"

	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/routes"
	"github.com/gorilla/mux"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	database.Connect()

	route := mux.NewRouter()
	routes.RegisterRoutes(route)
	log.Fatal(http.ListenAndServe(":8080", route))
}
