package main

import (
	"log"
	"net/http"

	"github.com/vikash-parashar/myapp/pkg/config"
	"github.com/vikash-parashar/myapp/pkg/handlers"
	"github.com/vikash-parashar/myapp/pkg/render"
)

const (
	port = ":8080"
)

func main() {
	app := config.AppConfig{}

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	log.Println("starting our application on port ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Println("failed to start application")
		return
	}
}
