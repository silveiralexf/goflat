package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func aboutPage() (string, g.Node) {
	return "About", Div(
		H1(g.Text("About this site")),
		P(g.Text("This is a site showing off gomponents.")),
	)
}
