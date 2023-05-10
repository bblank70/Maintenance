package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bblank70/bookings/internal/config"
	"github.com/bblank70/bookings/internal/handlers"
	"github.com/bblank70/bookings/internal/render"
)

const port = ":8080"

// expiration time
const expirationInverval = 2 * time.Hour

var app config.AppConfig
var session *scs.SessionManager

func main() {
	fmt.Printf("Starting applicaiton on %s \n", port)

	app.InProduction = false
	session = scs.New()

	// sets the session cookie to expiration interval
	session.Lifetime = expirationInverval
	session.Cookie.Persist = false
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc

	//........allows user to enter debug mode for rendering testing......///
	app.UseCache = false

	//repo will build a new repo to the app reference
	repo := handlers.NewRepo(&app)
	//instantiates the handlers
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// _ = http.ListenAndServe(port, nil)

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
