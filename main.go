package main

import (
	"log"
	"net/http"
	"psychic-parakeet-go/init/middleware"
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

func getTheHostName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Psychic Parakeet Hostname is " + r.Host))
}

func main() {
	router := http.NewServeMux()
	// Adding a method param to a route can be done by adding a space after the method name
	// and then the route path
	router.HandleFunc("GET /item/{id}", getTheId)
	router.HandleFunc("POST /data/", getTheData)
	// Hostname can be added to the route by adding a space after the hostname
	// spoof this in curl with -H "Host: psychicparakeet.com"
	router.HandleFunc("psychicparakeet.com/", getTheHostName)

	server := http.Server{
		Addr: ":8080",
		// The middleware.Logging function is called with the router as an argument
		// this just logs the request method, but can be handled in any other way
		// by changing the function in the middleware package
		// examples of other middleware functions are in the middleware package
		// cache, cors, and logging
		Handler: middleware.Logging(router),
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
