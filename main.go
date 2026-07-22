package main

import (
//	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// 1. Create a new http.ServeMux to route requests
	mux := http.NewServeMux()

	// Add routes
//	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello, World! You're at: %s", r.URL.Path)
//	})
//
//	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
//		w.Write([]byte("pong"))
//	})

	// 2. Create a new http.Server struct with additional configuration
	server := &http.Server{
		// 2.1 Use the new "ServeMux" as the server's handler
		Handler: mux,
		// 2.2 Set the .Addr field to ":8080"
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("Server starting on :8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
