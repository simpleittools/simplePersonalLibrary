package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/simpleittools/simplepersonallibrary/internal/config"
	"github.com/simpleittools/simplepersonallibrary/internal/handlers"
	"github.com/simpleittools/simplepersonallibrary/internal/render"
	"log"
	"net/http"
	"time"
)

// TODO: move port to environmental variable
const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
