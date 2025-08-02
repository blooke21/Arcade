// main.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Response structure
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Single API endpoint
	http.HandleFunc("/api/hello", helloHandler)
	
	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	fmt.Println("Try: http://localhost:8080/api/hello")
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers for Electron
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	
	// Create response
	response := Response{
		Message: "Hello World from Go backend!",
	}
	
	// Send JSON response
	json.NewEncoder(w).Encode(response)
}