package handlers

import (
	"github.com/simpleittools/simplepersonallibrary/internal/config"
	"github.com/simpleittools/simplepersonallibrary/internal/models"
	"github.com/simpleittools/simplepersonallibrary/internal/render"
	"net/http"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page function
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.TemplateRenderer(
		w, r, "home.page.gohtml", &models.TemplateData{
			StringMap: stringMap,
		},
	)
}

// About this is the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some business logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again."

	// send the data to the template
	render.TemplateRenderer(
		w, r, "about.page.gohtml", &models.TemplateData{
			StringMap: stringMap,
		},
	)
}
