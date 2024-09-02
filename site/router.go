package site

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silveiralexf/goflat/site/pages"
)

// registers site pages, to be served by pocketbase
// passes `pb` to allow access to `DAO` and other apis
// each page will get auth data in request context
// and will be able to create all necessary info for page render:
// user data, external api calls, calculations
func AddPageRoutes(pb *pocketbase.PocketBase) {
	pb.OnBeforeServe().Add(pages.GetDataPageRoute(pb))
	pb.OnBeforeServe().Add(pages.GeProfileRoute(pb))
	pb.OnBeforeServe().Add(pages.GetErrorPageRoute(pb))
	pb.OnBeforeServe().Add(pages.GetIndexPageRoute(pb))

	pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.StaticFS("/static", staticFilesFS)
		// this path works : http://127.0.0.1:8090/static/static/public/htmx.min.js
		return nil
	})
}
