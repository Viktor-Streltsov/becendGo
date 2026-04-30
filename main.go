package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("Hello from Go 🚀"))
	})

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}