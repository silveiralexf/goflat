package middleware

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

var redirectTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="refresh" content="0; url='/error/{{ . }}'" />
  </head>
  <body>
    <p>Redirecting to error page</p>
  </body>
</html>
`
var tmpl = template.Must(template.New("redirect-to-error").Parse(redirectTemplate))

func AddErrorsMiddleware(app *pocketbase.PocketBase) {
	app.OnBeforeApiError().Add(func(e *core.ApiErrorEvent) error {
		if e.HttpContext.Response().Status == http.StatusOK {
			err := e.HttpContext.Redirect(302, "/")
			if err != nil {
				panic(fmt.Errorf("Invalid respond on redirect: %v", err)) // TODO: this should not happen, right?
			}
		}
		return renderErrorPage(e)
	})
}

func renderErrorPage(e *core.ApiErrorEvent) error {
	errorMessage := e.Error.Error()
	errorCode := e.HttpContext.Response().Status

	switch errorMessage {
	case "Not Found.":
		errorCode = http.StatusNotFound
	case "The request requires admin or record authorization token to be set.":
		errorCode = http.StatusUnauthorized
	}

	var instantiatedTemplate bytes.Buffer
	if err := tmpl.Execute(&instantiatedTemplate, errorCode); err != nil {
		return e.HttpContext.HTML(http.StatusOK, "Error 500")
	}

	return e.HttpContext.HTML(http.StatusOK, instantiatedTemplate.String())
}
