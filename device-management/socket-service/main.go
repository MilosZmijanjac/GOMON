package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/micro/plugins/v5/registry/consul"
	"go-micro.dev/v5/web"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Error loading .env file: %v\n", err)
		return
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		log.Println("SERVER_ADDRESS is not set in .env file!")
		return
	}

	r := consul.NewRegistry()
	service := web.NewService(
		web.Registry(r),
		web.Name("socket-service"),
		web.Address(serverAddress),
	)
	// Initialize the service
	service.Init()
	hub := NewHub()
	go hub.run()
	// Set up a route and handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	// Assign the handler to the service
	service.Handle("/", http.DefaultServeMux)

	// Start the service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
