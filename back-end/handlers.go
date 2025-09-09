// HTTP request/response logic
package main

import (
	"net/http"
)

func addRomHandler(w http.ResponseWriter, r *http.Request, sourcePath string) {
    setCORSHeaders(w)
    
    fileMap := handleMoveFile(sourcePath)
    updateRomDatabase(fileMap)
}

func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}