package pages

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/silveiralexf/goflat/internal/middleware"
	"github.com/silveiralexf/goflat/internal/utils"
)

func GetDataPageRoute(pb *pocketbase.PocketBase) func(*core.ServeEvent) error {
	return func(e *core.ServeEvent) error {
		e.Router.GET("/data", func(ctx echo.Context) error {
			// get data
			// and since i'm using 'base.gohtml' with Nav, i'll need Nav info
			s := middleware.NewSession(pb, ctx)

			pageData := struct {
				RandomNumber int
				RandomString string
				NavInfo      middleware.SessionData
				Content      []template.HTML
			}{
				RandomNumber: rand.Int(),
				RandomString: utils.StringWithCharset(25, utils.GetCharset()),
				NavInfo:      *s,
				Content:      GetDataCollection(pb, s),
			}

			// then render template with it
			templateName := "templates/data.gohtml"
			tmpl := template.Must(template.ParseFS(templatesFS, "templates/base.gohtml", templateName))

			var instantiatedTemplate bytes.Buffer
			if err := tmpl.Execute(&instantiatedTemplate, pageData); err != nil {
				return ctx.JSON(http.StatusInternalServerError, map[string]string{"message": "error parsing template"})
			}

			return ctx.HTML(http.StatusOK, instantiatedTemplate.String())
		}, apis.RequireAdminOrRecordAuth())
		return nil
	}
}

func GetDataCollection(pb *pocketbase.PocketBase, s *middleware.SessionData) []template.HTML {
	// retrieve multiple "articles" collection records by a string filter expression
	// (use "{:placeholder}" to safely bind untrusted user input parameters)
	records, err := pb.Dao().FindRecordsByFilter(
		"data", // collection
		fmt.Sprintf("status = 'public' && author = '%s'", s.ApisAuthRecord.Id),
		"-created", // sort
		10,         // limit
		0,          // offset
	)
	if err != nil {
		return nil
	}

	data := []template.HTML{}
	for _, v := range records {
		contentMarkdown := v.Get("content").(string)
		contentStr := mdToHTML([]byte(contentMarkdown))

		data = append(data, template.HTML(contentStr))
	}

	return data
}

func mdToHTML(md []byte) string {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer))
}
