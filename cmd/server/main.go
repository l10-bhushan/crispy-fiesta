package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/l10-bhushan/crispy-fiesta/internal/handlers"
	"github.com/l10-bhushan/crispy-fiesta/internal/middleware"
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
	// Adding middleware to the router
	// Each middleware should wrap the previous handler so the chain is preserved.
	// Order: RequestId -> Logger -> Recovery
	handler := middleware.Logger(mux)
	handler = middleware.Recovery(handler)
	handler = middleware.RequestId(handler)

	// The sequence of above middleware will be
	// RequestID - will fetch the x-request-id from the header if available, or generate one and store it in request context.
	// Logger - Logs the request infromation such as method, path , requestID.
	// Recovery - uses defer and recover to tackle panics in our code.

	// Adding a route to our router
	// A simple health router
	mux.HandleFunc("GET /health", handlers.HealthHandler)
	// A simple version router
	mux.HandleFunc("GET /version", handlers.VersionHandler)
	// A simple panic handler to test "Recovery" middleware
	mux.HandleFunc("GET /panic", handlers.PanicHandler)

	// Configuring the server, server has many different properties as well.
	// But for now we will only use Addr and Handler
	// Addr: takes the port no
	// Handler takes the router
	server := http.Server{
		Addr:    port,
		Handler: handler,
	}

	// Starting the server as well as cheking for errors.
	// Here we are using a go routine because we don't want our server to run on the main function
	// for eg: if we have two statements
	// server.ListenAndServe
	// print("hello")
	// The program would never reach the hello part because our main function is occupied by the server
	// So, if we want a graceful shutdown in that case, we separate the server onto a different go routine
	// then the main with that we can catch the SIGINT and SIGTERM signals and shutdown our server.
	go func() {
		// Information for the user.
		fmt.Printf("Server up and running on PORT: %s\n", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// This creates a channel with type os.Signal, and the 1 denotes the capacity of the channel
	// to hold only one signal at a time.
	shutdownSignal := make(chan os.Signal, 1)

	// This basically means that whenever the program receives one of the signals store it in shutdownSignal
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM)

	// This means "Receive a signal from the channels"
	// So, it's basically
	// 	main
	//  ↓
	// server started
	//  ↓
	// register signals
	//  ↓
	// <-shutdownSignal
	//  ↓
	// WAITING...
	// Meanwhile, the server is running in it's separate go routine
	<-shutdownSignal

	log.Printf("Shutdown signal received: %v", shutdownSignal)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	// This tells the server don't accept new requests and gracefully shutdown.
	// If a request takes more than the context time limit, the graceful shutdown returns an error
	// but that does not mean it didn't shutdown, it means it exceeded the limit of the context.
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("server failed to shutdown : %v", err)
	}

	log.Println("server stopped gracefully")
}
