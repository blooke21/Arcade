//HTTP request/response logic
package main

import (
    "encoding/json"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    setCORSHeaders(w)
    
    message := getHelloMessage() // Call business logic
    
    response := map[string]string{"message": message}
    json.NewEncoder(w).Encode(response)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	setCORSHeaders(w)
	
	message := getTestMessage() // Call business logic
	
	response := map[string]string{"message": message}
	json.NewEncoder(w).Encode(response)
}

func moveFileHandler(w http.ResponseWriter, r *http.Request, sourcePath string) {
    setCORSHeaders(w)
    
    handleMoveFile(sourcePath) // Call business logic
}

func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}