package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/templates"
)

// TemplateData holds the data that we pass to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
}

// HomePageHandler displays the home page
func HomePageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "home.page.tmpl", nil, app)
	}
}

// AboutPageHandler displays the about page
func AboutPageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "about.page.tmpl", nil, app)
	}
}

// ContactPageHandler displays the contact page
func ContactPageHandler(app config.AppConfig) http.HandlerFunc {
	stringMap := make(map[string]string)
	stringMap["phone"] = "+19025551212"

	td := TemplateData{
		StringMap: stringMap,
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, r, "contact.page.tmpl", &td, app)
	}
}

func render(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData, app config.AppConfig) {
	var templateCache map[string]*template.Template
	if !app.UseCache {
		_ = templates.NewTemplateCache(&app)
	}
	templateCache = app.TemplateCache

	ts, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Cannot retrieve template")
		return
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, AddDefaultData(td, r, w))

	if err != nil {
		log.Fatal(w, err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// AddDefaultData adds default data to the template
func AddDefaultData(td *TemplateData, r *http.Request, w http.ResponseWriter) *TemplateData {
	if td == nil {
		td = &TemplateData{}
	}
	return td
}
