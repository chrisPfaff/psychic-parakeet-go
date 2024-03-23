package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Psychic Parakeet"))
	})
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
