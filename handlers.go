package main

import (
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is Home Page")
	renderTemplate(w, "home", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is About Page")
	renderTemplate(w, "about", nil)
}
func ListOfUser(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is About Page")
	renderTemplate(w, "user", users)
}
