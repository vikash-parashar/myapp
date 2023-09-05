package main

import (
	"html/template"
	"log"
	"net/http"
)

func renderTemplate(w http.ResponseWriter, name string, data any) {
	tpl, err := template.ParseFiles("./templates/" + name)
	if err != nil {
		log.Fatalln("failed to parse files")
		return
	}
	if err = tpl.Execute(w, data); err != nil {
		log.Println("failed to execute template")
		return
	}
	log.Println(name + "template executed successfully")

}
