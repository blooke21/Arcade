// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

var ErrDuplicateROM = fmt.Errorf("ROM already exists in database")

func main() {

	//register handlers

	http.HandleFunc("/api/add-rom", func(w http.ResponseWriter, r *http.Request) {
		log.Println("add-rom endpoint hit")
		var req struct {
        	SourcePath string `json:"sourcePath"`
    	}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        	http.Error(w, "Invalid JSON", http.StatusBadRequest)
        	return
    	}
		addRomHandler(w, r)
	})
	
	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}