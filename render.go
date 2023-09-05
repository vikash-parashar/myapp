package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
)

func renderTemplate(w http.ResponseWriter, name string, data any) {
	tpl, err := template.ParseFiles("./templates/" + name + ".html")
	if err != nil {
		log.Fatalln("failed to parse files")
		return
	}
	if err = tpl.Execute(w, data); err != nil {
		log.Println("failed to execute template")
		return
	}
	log.Println(strings.ToUpper(name) + " template executed successfully")

}
