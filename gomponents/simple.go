package gomponents

import (
	"github.com/SlinSo/goTemplateBenchmark/model"
	g "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
)

func Page(u *model.User) g.Node {
	return HTML(
		Body(
			H1(g.Text(u.FirstName)),
			P(g.Text("Here's a list of your favorite colors:")),
			Ul(
				g.Map(u.FavoriteColors, func(colorname string) g.Node {
					return Li(g.Text(colorname))
				}),
			),
		),
	)
}
