// HTTP request/response logic
package main

import (
	"net/http"
)

func RomListHandler(w http.ResponseWriter, _ *http.Request) {
    setCORSHeaders(w)

    roms, err := returnRomList()
    if err != nil {
        http.Error(w, "Failed to get ROM list", http.StatusInternalServerError)
        return
    }
    w.Write(roms)
}

func addRomHandler(w http.ResponseWriter, _ *http.Request, sourcePath string) {
    setCORSHeaders(w)

    fileMap, err := handleMoveFile(sourcePath)
    if err == ErrDuplicateROM {
        http.Error(w, "ROM already exists", http.StatusConflict)
        return
    } else if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    addRomDatabase(fileMap)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message":"ROM ` + fileMap["fileName"] + ` added successfully to ` + fileMap["type"] + `"}`))
}

func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}