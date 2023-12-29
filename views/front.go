package views

import (
	g "github.com/maragudk/gomponents"
	h "github.com/maragudk/gomponents/html"
)

func FrontPage() g.Node {
	return Page(
		"Canvas",
		"/",
		h.H1(g.Text("Solutions to problems")),
		h.P(g.Text("Do you have problems? We also had problems.")),
		h.P(g.Raw("Then we created the <em>canvas</em> app, and now we don't! 😬")),
	)
}
