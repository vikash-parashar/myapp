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
	var templates []string
	var tmp *template.Template
	var err error

	tmp, ok := tc[t]
	templates = []string{
		fmt.Sprintf("./templates/%s.page.tmpl", t),
		"./templates/base.layout.tmpl",
	}
	if !ok {
		tmp, err = createTemplateCache(templates)
		if err != nil {
			log.Println("failed to parse templates . . . . !")
			return
		}
		log.Println("adding template into templates cache")
		tc[t] = tmp
	} else {
		log.Println("using cache template")

	}
	err = tmp.Execute(w, data)
	if err != nil {
		log.Fatalln("failed to execute", strings.ToUpper(t), "template")
		return
	}

}

func createTemplateCache(tmpls []string) (*template.Template, error) {
	tmp, err := template.ParseFiles(tmpls...)
	if err != nil {
		return tmp, err
	}
	return tmp, nil
}
