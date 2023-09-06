package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file :", err)
	}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not set")
	}

	fmt.Println("App is running on port: ", portStr)
}
