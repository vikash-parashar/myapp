package handlers

import (
	"net/http"

	"github.com/vikash-parashar/myapp/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is Home Page")
	render.RenderTemplate(w, "home", nil)
}
func About(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "This is About Page")
	render.RenderTemplate(w, "about", nil)
}
