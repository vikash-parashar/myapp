package handlers

import (
	"myapp/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is Home Page")
	render.RenderTemplate(w, "home", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is About Page")
	render.RenderTemplate(w, "about", nil)
}
