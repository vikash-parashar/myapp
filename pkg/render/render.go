package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string, data any) {
	var tmp *template.Template
	var err error

	tmp, ok := tc[t]

	if !ok {
		err = createTemplateCache(t)
		if err != nil {
			log.Println("failed to parse templates . . . . !")
			return
		}
		log.Println("adding template into templates cache")
	} else {
		log.Println("using cache template")

	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Fatalln("failed to execute", strings.ToUpper(t), "template")
		return
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s.page.tmpl", t),
		"./templates/base.layout.tmpl",
	}
	tmp, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmp
	return nil
}
