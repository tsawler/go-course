package handlers

import (
	"html/template"
	"log"
	"net/http"
)

// TemplateData holds the data that we pass to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
}

// HomePageHandler displays the home page
func HomePageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl", nil)
}

// AboutPageHandler displays the about page
func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl", nil)
}

// ContactPageHandler displays the contact page
func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["phone"] = "+19025551212"

	renderTemplate(w, "contact.page.tmpl", &TemplateData{
		StringMap: stringMap,
	})
}

// renderTemplate renders a page template
func renderTemplate(w http.ResponseWriter, tmpl string, td *TemplateData) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, td)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
