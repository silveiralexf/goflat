package pages

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silveiralexf/goflat/internal/middleware"
)

func GetErrorPageRoute(pb *pocketbase.PocketBase) func(*core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		e.Router.GET("/favicon.ico", func(ctx echo.Context) error {
			return ctx.HTML(http.StatusNotFound, "icon not found")
		})
		e.Router.GET("/sw.js", func(ctx echo.Context) error {
			return ctx.HTML(http.StatusNotFound, "icon not found")
		})

		e.Router.GET("/error/:code", func(ctx echo.Context) error {
			// get data
			code := ctx.PathParam("code")
			codeNum, err := strconv.ParseInt(code, 10, 64)
			if err != nil {
				codeNum = 500
			}
			errorText := http.StatusText(int(codeNum))
			if errorText == "" {
				codeNum = 500
				errorText = http.StatusText(500)
			}

			// and since i'm using 'base.gohtml' with Nav, i'll need Nav info
			s := middleware.NewSession(pb, ctx)

			PageData := struct {
				NavInfo   middleware.SessionData
				ErrorCode int64
				ErrorText string
			}{
				NavInfo:   *s,
				ErrorCode: codeNum,
				ErrorText: errorText,
			}

			// then render template with it
			templateName := "templates/errors/error.gohtml"
			switch codeNum {
			case 404:
				templateName = "templates/errors/404.gohtml"
			case 401:
				templateName = "templates/errors/401.gohtml"
			}
			tmpl := template.Must(template.ParseFS(templatesFS, "templates/base.gohtml", templateName))
			var instantiatedTemplate bytes.Buffer
			if err := tmpl.Execute(&instantiatedTemplate, PageData); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error parsing template"})
			}

			return ctx.HTML(int(codeNum), instantiatedTemplate.String())
		})
		return nil
	}
}
