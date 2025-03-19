package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/arwahyu01/go-jwt/database"
	"github.com/arwahyu01/go-jwt/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}

	database.Connect()

	route := mux.NewRouter()
	routes.RegisterRoutes(route)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Ganti dengan domain jika sudah deploy
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("🚀 Server running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, corsHandler.Handler(route)))
}
