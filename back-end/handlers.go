// HTTP request/response logic
package main

import (
	"net/http"
)

func addRomHandler(w http.ResponseWriter, r *http.Request) {
    setCORSHeaders(w)
    sourcePath := r.URL.Query().Get("sourcePath")
    
    fileMap, err := handleMoveFile(sourcePath)
    if err == ErrDuplicateROM {
        http.Error(w, "ROM already exists", http.StatusConflict)
        return
    } else if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    updateRomDatabase(fileMap)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"ROM ` + fileMap["fileName"] + ` added successfully"}`))
}

func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}