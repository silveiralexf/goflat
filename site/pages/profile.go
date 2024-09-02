package pages

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silveiralexf/goflat/internal/middleware"
)

func GeProfileRoute(pb *pocketbase.PocketBase) func(*core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		e.Router.GET("/profile", func(ctx echo.Context) error {
			// get data
			// and since i'm using 'base.gohtml' with Nav, i'll need Nav info
			s := middleware.NewSession(pb, ctx)

			profileData := struct {
				NavInfo middleware.SessionData
			}{
				NavInfo: *s,
			}

			// then render template with it
			templateName := "templates/profile.gohtml"
			tmpl := template.Must(template.ParseFS(templatesFS, "templates/base.gohtml", templateName))
			var instantiatedTemplate bytes.Buffer
			if err := tmpl.Execute(&instantiatedTemplate, profileData); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error parsing template"})
			}

			return ctx.HTML(http.StatusOK, instantiatedTemplate.String())
		}, apis.RequireAdminOrRecordAuth())
		return nil
	}
}
