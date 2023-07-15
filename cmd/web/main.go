package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/simpleittools/simplepersonallibrary/internal/config"
	"github.com/simpleittools/simplepersonallibrary/internal/handlers"
	"github.com/simpleittools/simplepersonallibrary/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

// TODO: move port to environmental variable
const port = ":3000"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("The server is running at http://localhost%s", port))
	server := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = server.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

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
	return nil
}
