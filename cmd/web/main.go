package main

import (
	"fmt"
	"github.com/simpleittools/simplepersonallibrary/pkg/config"
	"github.com/simpleittools/simplepersonallibrary/pkg/handlers"
	"github.com/simpleittools/simplepersonallibrary/pkg/render"
	"log"
	"net/http"
)

// TODO: move port to environmental variable
const port = ":3000"

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot creat template cache")
	}

	app.TemplateCache = templateCache

	render.NewTemplates(&app)
	// TODO: move this to a routes handler
	http.HandleFunc(
		"/", handlers.Home,
	)
	fmt.Println(fmt.Sprintf("The server is running at http://localhost%s", port))
	_ = http.ListenAndServe(port, nil)

}
