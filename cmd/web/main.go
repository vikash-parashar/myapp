package main

import (
	"log"
	"myapp/pkg/handlers"
	"net/http"
)

const (
	port = ":8080"
)

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Println("starting our application on port : ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("failed to start application")
		return
	}
}
