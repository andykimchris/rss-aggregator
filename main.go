package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file :", envErr)
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not set")
	}
	fmt.Println("Server is running on port:", portStr)

	router := chi.NewRouter()

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}

}
