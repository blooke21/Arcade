// HTTP request/response logic, handles interactions with frontend
package main

import (
    "encoding/json"
	"net/http"
)

// RomListHandler handles fetching the list of ROMs
func RomListHandler(w http.ResponseWriter, _ *http.Request) {
    setCORSHeaders(w)

    roms, err := returnRomList()
    if err != nil {
        http.Error(w, "Failed to get ROM list", http.StatusInternalServerError)
        return
    }
    w.Write(roms)
}

// addRomHandler handles adding a new ROM file
func addRomHandler(w http.ResponseWriter, _ *http.Request, sourcePath string) {
    setCORSHeaders(w)

    fileMap, err := buildMoveFile(sourcePath)
    if err == ErrDuplicateROM {
        http.Error(w, "ROM already exists", http.StatusConflict)
        return
    } else if err != nil {
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    romID, err := addRomDatabase(fileMap)
    if err != nil {
        http.Error(w, "Failed to add ROM to database", http.StatusInternalServerError)
        return
    }

    resp := struct {
        Message string `json:"message"`
        RomID   int    `json:"romID"`
    }{
        Message: "ROM " + fileMap["fileName"] + " added successfully to " + fileMap["type"],
        RomID:   romID,
    }
    body, err := json.Marshal(resp)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(body)
}

// editRomHandler handles changes the details of passed ROM
func editRomHandler(w http.ResponseWriter, r *http.Request) {
    setCORSHeaders(w)
    var req struct {
        RomID   string `json:"romID"`
        NewName string `json:"newName"`
        ImgPath string `json:"imgPath"`
    }
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    err := editRomDatabase(req.RomID, req.NewName, req.ImgPath)
    if err != nil {
        http.Error(w, "Failed to edit ROM in database", http.StatusInternalServerError)
        return
    }

    resp := struct {
        Message string `json:"message"`
    }{
        Message: "ROM " + req.RomID + " edited successfully",
    }
    body, err := json.Marshal(resp)
    if err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(body)
}

// setCORSHeaders sets the necessary CORS headers for the response
func setCORSHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
}