package render

import (
	"bytes"
	"github.com/simpleittools/simplepersonallibrary/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

// TemplateRenderer will render templates
func TemplateRenderer(w http.ResponseWriter, gohtml string) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		// create a template cache
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	tmpl, ok := templateCache[gohtml]
	if !ok {
		log.Println("unable to get pages from cache")
	}

	buffer := new(bytes.Buffer)

	err := tmpl.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	pageCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return pageCache, err
	}

	// range through all the pages for the *.page.gohtml files
	for _, page := range pages {
		pageName := filepath.Base(page)

		templateSet, err := template.New(pageName).ParseFiles(page)
		if err != nil {
			return pageCache, err
		}

		layoutFiles, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return pageCache, err
		}

		if len(layoutFiles) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return pageCache, err
			}
		}

		pageCache[pageName] = templateSet
	}

	return pageCache, nil
}
