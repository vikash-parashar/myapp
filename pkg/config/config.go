package config

import (
	"database/sql"
	"html/template"
	"log"
)

// AppConfig holds the application configuration
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	Db            *sql.DB
}
