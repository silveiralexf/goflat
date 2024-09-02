package pages

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silveiralexf/goflat/internal/middleware"
)

// sometest
// render and return some page
func GetIndexPageRoute(pb *pocketbase.PocketBase) func(*core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		e.Router.GET("/", func(ctx echo.Context) error {
			// first collect data
			s := middleware.NewSession(pb, ctx)
			indexPageData := struct {
				BackendMessage string
				NavInfo        middleware.SessionData
			}{
				BackendMessage: "Hello from the backend!",
				NavInfo:        *s,
			}

			// then render template with it
			templateName := "templates/index.gohtml"
			tmpl := template.Must(template.ParseFS(templatesFS, "templates/base.gohtml", templateName))
			var instantiatedTemplate bytes.Buffer
			if err := tmpl.Execute(&instantiatedTemplate, indexPageData); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error parsing template"})
			}

			return ctx.HTML(http.StatusOK, instantiatedTemplate.String())
		})
		return nil
	}
}
