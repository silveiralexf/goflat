package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func contactPage() (string, g.Node) {
	return "Contact", Div(
		H1(g.Text("Contact us")),
		P(g.Text("Just do it.")),
	)
}
