package handlers

import (
	"github.com/simpleittools/simplepersonallibrary/pkg/config"
	"github.com/simpleittools/simplepersonallibrary/pkg/render"
	"net/http"
)

var Repo *Repository

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
	render.TemplateRenderer(w, "home.page.gohtml")
}
