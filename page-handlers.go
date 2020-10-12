package main

import (
	"html/template"
	"log"
	"net/http"
)

// AboutPageHandler displays the about page
func AboutPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// ContactPageHandler displays the contact page
func ContactPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

// renderTemplate renders a page template
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("Error executing template :", err)
		return
	}
}
