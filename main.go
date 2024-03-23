package main

import (
	"log"
	"net/http"
)

func getTheId(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Psychic Parakeet id is " + id))
}

func getTheData(w http.ResponseWriter, r *http.Request) {
	data := r.FormValue("data")
	if data == "" {
		http.Error(w, "No data provided", http.StatusBadRequest)
		return
	} else if data == "parakeet" {
		http.Error(w, "Parakeet is not allowed", http.StatusBadRequest)
		return
	} else {
		w.Write([]byte("Psychic Parakeet data is " + data))
	}
}

func main() {
	router := http.NewServeMux()
	// Adding a method param to a route can be done by adding a space after the method name
	// and then the route path
	router.HandleFunc("GET /item/{id}", getTheId)
	router.HandleFunc("POST /data/", getTheData)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
