package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	handler "github.com/annguyen0511/bookings/pkg/handlers"
	"github.com/annguyen0511/bookings/pkg/render"

	"github.com/annguyen0511/bookings/pkg/config"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handler.NewRepo(&app)

	handler.NewHandler(repo)

	fmt.Printf("Starting application on port %s\n", portNumber)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
