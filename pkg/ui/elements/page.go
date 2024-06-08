package elements

import (
	"time"

	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

// Page generates an HTML5 boilerplate document
func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(
		c.HTML5Props{
			Title:    title,
			Language: "en",
			Head: []g.Node{
				Script(Src("https://cdn.tailwindcss.com?plugins=forms,typography")),
				Script(Src("https://unpkg.com/htmx.org")),
			},
			Body: []g.Node{
				Navbar(path, []PageLink{
					{Path: "/contact", Name: "Contact"},
					{Path: "/about", Name: "About"},
					{Path: "/photos", Name: "Photos"},
				}),
				Container(
					Prose(body),
					PageFooter(),
				),
			},
		},
	)
}

func Container(children ...g.Node) g.Node {
	return Div(
		Class("max-w-7xl mx-auto px-2 sm:px-6 lg:px-8"),
		g.Group(children),
	)
}

func Prose(children ...g.Node) g.Node {
	return Div(Class("prose"), g.Group(children))
}

func PageFooter() g.Node {
	return Footer(
		Class("prose prose-sm prose-indigo"),
		P(
			g.Textf("Rendered %v. ", time.Now().Format(time.RFC3339)),

			// Conditional inclusion
			g.If(time.Now().Second()%2 == 0, g.Text("It's an even second.")),
			g.If(time.Now().Second()%2 == 1, g.Text("It's an odd second.")),
		),
		P(
			A(
				Href("https://www.gomponents.com"), g.Text("gomponents"),
			),
		),
	)
}

func PartialUpdate(now time.Time) g.Node {
	timeFormat := "15:04:05"
	return P(ID("partial"), g.Textf(`Time was last updated at %v.`, now.Format(timeFormat)))
}
