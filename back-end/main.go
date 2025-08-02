// main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {

	//register handlers
	http.HandleFunc("/api/hello", helloHandler)
	http.HandleFunc("/api/test", testHandler)
	
	// Start server
	port := "8080"
	fmt.Printf("Go backend running on http://localhost:%s\n", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}