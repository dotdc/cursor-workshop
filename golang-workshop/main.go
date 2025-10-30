package main

import (
	"log"
	"os"

	"golang-workshop/app"
)

func main() {
	server, err := app.NewServer()
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := server.ListenAndServe(":" + port); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
