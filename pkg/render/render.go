package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// TemplateRenderer will render templates
func TemplateRenderer(w http.ResponseWriter, gohtml string) {
	templateCache, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	tmpl, ok := templateCache[gohtml]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	err = tmpl.Execute(buffer, nil)
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
