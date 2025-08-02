// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)


func main() {

	//register handlers
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/test", testHandler)
	http.HandleFunc("/api/move-file", func(w http.ResponseWriter, r *http.Request) {
		log.Println("move-file endpoint hit")
		var req struct {
        	SourcePath string `json:"sourcePath"`
    	}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        	http.Error(w, "Invalid JSON", http.StatusBadRequest)
        	return
    	}
		moveFileHandler(w, r, req.SourcePath)
	})
	
	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}