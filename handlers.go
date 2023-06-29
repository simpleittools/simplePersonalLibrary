package main

import (
	"net/http"
)

// Home is the home page function
func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home.page.gohtml")
}
