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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("The server is running at http://localhost%s", port))
	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	log.Fatal(err)
}
