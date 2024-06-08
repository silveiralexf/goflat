package pages

import (
	"time"

	g "github.com/maragudk/gomponents"
	hx "github.com/maragudk/gomponents-htmx"
	. "github.com/maragudk/gomponents/html"

	"github.com/silveiralexf/goflat/pkg/ui/elements"
)

func photosPage() (string, g.Node) {
	now := time.Now()
	return "Photos",
		Div(Class("max-w-7xl mx-auto p-4 prose lg:prose-lg xl:prose-xl"),
			H1(g.Text(`gomponents + HTMX`)),
			elements.PartialUpdate(now),
			FormEl(Method("post"), Action("/photos"), hx.Boost("true"), hx.Target("#partial"), hx.Swap("innerHTML"),
				Button(
					Type("submit"),
					g.Text(`Update time`),
					Class(`
					bg-orange-600
					border
					border-transparent
					focus:outline-none focus:ring-2
					focus:ring-offset-2
					focus:ring-orange-500
					font-medium
					px-4 py-2
					rounded-md
					shadow-sm hover:bg-orange-700
					text-sm
					text-white
					`),
				),
			),
		)
}
