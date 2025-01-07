package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/silveiralexf/goflat/internal/middleware"
	"github.com/silveiralexf/goflat/internal/utils"
	"github.com/silveiralexf/goflat/site"
)

func main() {
	pb := pocketbase.New()

	middleware.AddCookieSessionMiddleware(pb)
	middleware.AddErrorsMiddleware(pb)
	site.AddPageRoutes(pb)

	// starts the pocketbase backend
	// parses cli arguments for hostname and data dir
	if err := pb.Start(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	utils.InitCollections()
}
