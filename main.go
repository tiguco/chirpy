package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// 1. Create a new http.ServeMux to route requests
	mux := http.NewServeMux()

	// 2. Use http.Dir to convert the current directory (".") to a directory for the FileServer
	dir := http.Dir(".")
	
	// 3. Create a standard http.FileServer using the directory
	fileServer := http.FileServer(dir)
	
	// 4. Use the ServeMux's .Handle() method to add a handler for the root path (/)
	//    The file server will serve index.html when someone visits the root
	mux.Handle("/", fileServer)

	// Create the HTTP server with the mux as the handler
	server := &http.Server{
		Handler:      mux,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server starting on http://localhost:8080")
	log.Println("Serving files from current directory (.)")
	log.Println("Try visiting: http://localhost:8080/index.html")
	
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
