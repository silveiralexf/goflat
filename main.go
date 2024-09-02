package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/silveiralexf/goflat/internal/middleware"
	"github.com/silveiralexf/goflat/site"
)

func main() {
	app := pocketbase.New()

	middleware.AddCookieSessionMiddleware(app)
	middleware.AddErrorsMiddleware(app)
	site.AddPageRoutes(app)

	// starts the pocketbase backend
	// parses cli arguments for hostname and data dir
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
