package handlers

import (
	"github.com/simpleittools/simplepersonallibrary/pkg/render"
	"net/http"
)

// Home is the home page function
func Home(w http.ResponseWriter, r *http.Request) {
	render.TemplateRenderer(w, "home.page.gohtml")
}
