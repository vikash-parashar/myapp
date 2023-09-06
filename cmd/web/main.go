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
	app.UseCache = true // keep false if you are in developer mod
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplate(&app)
	log.Println("starting our application on port ", port)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println("failed to start application")
		return
	}
}
