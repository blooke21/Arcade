// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"os"
	"path/filepath"
)

var ErrDuplicateROM = fmt.Errorf("ROM already exists in database")

func main() {

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

	// Resolve image directory (env var preferred)
    imgDir := os.Getenv("ROM_IMG_DIR")
    if imgDir == "" {
        exePath, err := os.Executable()
        if err != nil {
            log.Printf("failed to get executable path: %v", err)
            // fallback to relative path from working dir
            imgDir = filepath.Join("..", "rom", "img")
        } else {
            exeDir := filepath.Dir(exePath)
            // rom is sibling to back-end: ../rom/img relative to executable dir
            imgDir = filepath.Clean(filepath.Join(exeDir, "..", "rom", "img"))
        }
    }

    // make absolute and check existence (log but continue)
    if abs, err := filepath.Abs(imgDir); err == nil {
        imgDir = abs
    }
    if _, err := os.Stat(imgDir); err != nil {
        log.Printf("warning: image directory %q not found: %v", imgDir, err)
    }
    
    fs := http.FileServer(http.Dir(imgDir))
    // Wrap to set CORS (your setCORSHeaders sets JSON content-type; for images we want default content-type)
    http.HandleFunc("/rom/img/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        // strip the prefix so FileServer sees only the filename
        http.StripPrefix("/rom/img/", fs).ServeHTTP(w, r)
    })

	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}