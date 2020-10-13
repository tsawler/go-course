package handlers

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/forms"
	"tsawler/go-course/pkg/templates"
)

var app *config.AppConfig

func NewHandlers(a *config.AppConfig, db *sql.DB) {
	app = a
}

// TemplateData holds the data that we pass to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}

// HomePageHandler displays the home page
func HomePageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		app.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
		render(w, r, "home.page.tmpl", nil, app)
	}
}

// AboutPageHandler displays the about page
func AboutPageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["remote_ip"] = app.Session.GetString(r.Context(), "remote_ip")
		render(w, r, "about.page.tmpl", &TemplateData{
			StringMap: stringMap,
		}, app)
	}
}

// ContactPageHandler displays the contact page
func ContactPageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["phone"] = "+19025551212"

		td := TemplateData{
			StringMap: stringMap,
			Form:      forms.New(nil),
		}
		render(w, r, "contact.page.tmpl", &td, app)
	}
}

// PostContactPageHandler handles posting of the contact page form
func PostContactPageHandler(app config.AppConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := forms.New(r.PostForm)

		form.Required("name", "email")
		form.IsEmail("email")
		form.MinLength("name", 3)
		userName := r.Form.Get("name")
		userEmail := r.Form.Get("email")

		if !form.Valid() {
			render(w, r, "contact.page.tmpl", &TemplateData{
				Form: form,
			}, app)
			return
		}

		app.Session.Put(r.Context(), "flash", fmt.Sprintf("The user entered %s abd %s", userName, userEmail))
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
	td.CSRFToken = nosurf.Token(r)
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	return td
}
