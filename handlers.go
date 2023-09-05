package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Home Page")
}
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is About Page")
}
