package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/users", ListOfUser)
	log.Println("starting our application on port : ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("failed to start application")
		return
	}
}
