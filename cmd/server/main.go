package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	// We are using a package godotenv to load environment variables from our .env file
	// os.Getenv simple won't give you the environment variables.

	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found, falling back to environment variables")
	}

	// Loading the port using os.Getenv
	port := os.Getenv("PORT")

	// If we don't get any value from port we add ourselves
	if port == "" {
		port = "8000"
	}

	// Appending : to port
	port = ":" + port

	// Initialising router using http.NewServeMux
	// http.NewServeMux is built in router provided by net/http package of go
	mux := http.NewServeMux()

	// Adding a route to our router
	// A simple health router
	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "ok",
		})
	}))

	// Configuring the server, server has many different properties as well.
	// But for now we will only use Addr and Handler
	// Addr: takes the port no
	// Handler takes the router
	server := http.Server{
		Addr:    port,
		Handler: mux,
	}

	// Information for the user.
	fmt.Printf("Server up and running on PORT: %s\n", port)

	// Starting the server as well as cheking for errors.
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
