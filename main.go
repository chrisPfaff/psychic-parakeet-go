package main

import (
	"log"
	"net/http"
)

func getTheId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Psychic Parakeet id is " + id))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/item/{id}", getTheId)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
