package pages

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func indexPage() (string, g.Node) {
	return "Welcome!", Div(
		H1(g.Text("Welcome to this example page")),
		P(g.Text("I hope it will make you happy. 😄 It's using TailwindCSS for styling.")),
	)
}
