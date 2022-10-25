package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/countBreadedDice/booking_go_try/pkg/config"
	"github.com/countBreadedDice/booking_go_try/pkg/handlers"
	"github.com/countBreadedDice/booking_go_try/pkg/render"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager
func main() {
	

	//Change this to true when in production
	app.InProduction = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false
	app.Session =session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/sum", handlers.Repo.Sum)
	// http.HandleFunc("/divide", handlers.Repo.Divide)
	fmt.Println(fmt.Sprintln("Starting app on port ", portNumber))
	// http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
