package config

import (
	"github.com/alexedwards/scs/v2"
	"html/template"
)

// AppConfig holds the app config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
