//HTTP request/response logic
package main

import (
    "net/http"
)

func addRomHandler(w http.ResponseWriter, r *http.Request, sourcePath string) {
    setCORSHeaders(w)
    
    handleMoveFile(sourcePath) // Call business logic
}

func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}