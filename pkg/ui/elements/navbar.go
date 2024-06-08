package elements

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type PageLink struct {
	Path string
	Name string
}

func Navbar(currentPath string, links []PageLink) g.Node {
	return Nav(Class("bg-gray-700 mb-4"),
		Container(
			Div(
				Class("flex items-center space-x-4 h-16"),
				NavbarLink("/", "Home", currentPath == "/"),

				// We can Map custom slices to Nodes
				g.Group(
					g.Map(links, func(l PageLink) g.Node { return NavbarLink(l.Path, l.Name, currentPath == l.Path) }),
				),
			),
		),
	)
}

// NavbarLink is a link in the Navbar.
func NavbarLink(path, text string, active bool) g.Node {
	return A(Href(path), g.Text(text),
		// Apply CSS classes conditionally
		c.Classes{
			"px-3 py-2 rounded-md text-sm font-medium focus:outline-none focus:text-white focus:bg-gray-700": true,
			"text-white bg-gray-900":                           active,
			"text-gray-300 hover:text-white hover:bg-gray-700": !active,
		},
	)
}
