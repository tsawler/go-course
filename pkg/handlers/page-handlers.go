package handlers

import (
	"bytes"
	"fmt"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"tsawler/go-course/pkg/config"
	"tsawler/go-course/pkg/driver"
	"tsawler/go-course/pkg/forms"
	"tsawler/go-course/pkg/repository"
	"tsawler/go-course/pkg/repository/dbrepo"
	"tsawler/go-course/pkg/templates"
)

// DBRepo holds the repository
type DBRepo struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

var Repo *DBRepo

// NewHandlers creates the handlers with a repo
func NewHandlers(repo *DBRepo) {
	Repo = repo
}

// NewDatabaseRepo returns a new database repository and app concif
func NewDatabaseRepo(db *driver.DB, a *config.AppConfig) *DBRepo {
	return &DBRepo{
		App: a,
		DB:  dbrepo.NewMySQLRepo(db.SQL, a),
	}
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
func (m *DBRepo) HomePageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := m.DB.AllUsers()
		if err != nil {
			log.Println(err)
		}

		for _, x := range users {
			log.Println(x.FirstName)
		}
		m.App.Session.Put(r.Context(), "remote_ip", r.RemoteAddr)
		m.Render(w, r, "home.page.tmpl", nil)
	}
}

// AboutPageHandler displays the about page
func (m *DBRepo) AboutPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["remote_ip"] = m.App.Session.GetString(r.Context(), "remote_ip")
		m.Render(w, r, "about.page.tmpl", &TemplateData{
			StringMap: stringMap,
		})
	}
}

// ContactPageHandler displays the contact page
func (m *DBRepo) ContactPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stringMap := make(map[string]string)
		stringMap["phone"] = "+19025551212"

		td := TemplateData{
			StringMap: stringMap,
			Form:      forms.New(nil),
		}
		m.Render(w, r, "contact.page.tmpl", &td)
	}
}

// PostContactPageHandler handles posting of the contact page form
func (m *DBRepo) PostContactPageHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := forms.New(r.PostForm)

		form.Required("name", "email")
		form.IsEmail("email")
		form.MinLength("name", 3)
		userName := r.Form.Get("name")
		userEmail := r.Form.Get("email")

		if !form.Valid() {
			m.Render(w, r, "contact.page.tmpl", &TemplateData{
				Form: form,
			})
			return
		}

		m.App.Session.Put(r.Context(), "flash", fmt.Sprintf("The user entered %s abd %s", userName, userEmail))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// Render renders a go template with data
func (m *DBRepo) Render(w http.ResponseWriter, r *http.Request, tmpl string, td *TemplateData) {
	var templateCache map[string]*template.Template
	if !m.App.UseCache {
		_ = templates.NewTemplateCache(m.App)
	}
	templateCache = m.App.TemplateCache

	ts, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Cannot retrieve template")
		return
	}

	buf := new(bytes.Buffer)
	err := ts.Execute(buf, m.AddDefaultData(td, r, w))

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
func (m *DBRepo) AddDefaultData(td *TemplateData, r *http.Request, w http.ResponseWriter) *TemplateData {
	if td == nil {
		td = &TemplateData{}
	}
	td.CSRFToken = nosurf.Token(r)
	td.Flash = m.App.Session.PopString(r.Context(), "flash")
	td.Warning = m.App.Session.PopString(r.Context(), "warning")
	td.Error = m.App.Session.PopString(r.Context(), "error")

	return td
}
