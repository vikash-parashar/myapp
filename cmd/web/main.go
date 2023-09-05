package main

import (
	"log"
	"net/http"

	"github.com/vikash-parashar/myapp/pkg/handlers"
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
