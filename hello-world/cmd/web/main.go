package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/et-codes/lab/hello-world/pkg/config"
	"github.com/et-codes/lab/hello-world/pkg/handlers"
	"github.com/et-codes/lab/hello-world/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting application on port %s\n", portNumber)
	log.Fatal(http.ListenAndServe(portNumber, nil))
}
