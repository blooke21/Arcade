// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
)

var ErrDuplicateROM = fmt.Errorf("ROM already exists in database")

var romImgs string

var romDB string

func main() {

	initalizeEnv();

	//register handlers

	http.HandleFunc("/api/roms", func(w http.ResponseWriter, r *http.Request) {
		log.Println("roms endpoint hit")
		RomListHandler(w, r)
	})

	http.HandleFunc("/api/add-rom", func(w http.ResponseWriter, r *http.Request) {
		log.Println("add-rom endpoint hit")
		var req struct {
        	SourcePath string `json:"sourcePath"`
    	}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        	http.Error(w, "Invalid JSON", http.StatusBadRequest)
        	return
    	}
		addRomHandler(w, r, req.SourcePath)
	})

	http.HandleFunc("/api/edit-rom", func(w http.ResponseWriter, r *http.Request) {
		log.Println("edit-rom endpoint hit")
		editRomHandler(w, r)
	})

	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initalizeEnv() {
	romImgs = os.Getenv("ROM_IMG_DIR")
    if romImgs == "" {
        romImgs = "/app/rom/img"
    }

	romDB = os.Getenv("ROM_DB_DIR")
    if romDB == "" {
        romDB = "/app/rom/rom_database.json"
    }
}
