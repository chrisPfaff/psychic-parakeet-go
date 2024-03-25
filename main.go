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

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}
	name := r.Form.Get("name")
	email := r.Form.Get("email")
	message := r.Form.Get("message")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status": "ok"}`))
	log.Println("data", name, email, message)
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
	router.HandleFunc("psychicparakeet-go.com/", getTheHostName)
	// Create a middleware stack
	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.Cors)

	server := http.Server{
		Addr: ":8080",
		// The middleware.Logging function is called with the router as an argument
		// this just logs the request method, but can be handled in any other way
		// by changing the function in the middleware package
		// examples of other middleware functions are in the middleware package
		// cache, cors, and logging
		Handler: stack(router),
	}
	log.Println("Server started at :8080")
	server.ListenAndServe()
}
