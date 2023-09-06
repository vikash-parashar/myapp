package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vikash-parashar/myapp/pkg/config"
	"github.com/vikash-parashar/myapp/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	tmpl = tmpl + ".page.tmpl"
	var tc map[string]*template.Template
	var err error
	// get the template cache from the app config
	// tc, err := CreateTemplateCache()
	// if err != nil {
	// 	log.Fatalln("failed to create template cache")
	// }
	// fmt.Println(tc)
	// Create Template Cache
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatalln(err)
		}
	}

	// get requested template from Cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("failed to get template from template cache")
	}

	//optional code if you need
	buf := new(bytes.Buffer)
	data = AddDefaultData(data)
	err = t.Execute(buf, data)
	if err != nil {
		log.Println(err)
	}
	// Render Template-
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser")
	}
	// t.Execute(w, tmpl)

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)

		//render pages first
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// then layout , and so on ...
		layoutMatches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(layoutMatches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		fmt.Println("Adding Templates into TemplatesCache For Browser")
		myCache[name] = ts

	}
	return myCache, nil
}
