package handlers

import (
	"net/http"

	"github.com/vikash-parashar/myapp/pkg/config"
	"github.com/vikash-parashar/myapp/pkg/models"
	"github.com/vikash-parashar/myapp/pkg/render"
)

// Repo the Repository used by handlers
var Repo *Repository

// Repository is the type Repository
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home", &models.TemplateData{})
}
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringmap := make(map[string]string)

	stringmap["test"] = "Hello Again"

	render.RenderTemplate(w, "about", &models.TemplateData{StringMap: stringmap})
}
