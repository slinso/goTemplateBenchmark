package gomponents

import (
	"github.com/SlinSo/goTemplateBenchmark/model"
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Page(u *model.User) g.Node {
	return HTML(
		Body(
			H1(g.Raw(u.FirstName)),
			P(g.Raw("Here's a list of your favorite colors:")),
			Ul(
				g.Group(g.Map(u.FavoriteColors, func(colorname string) g.Node {
					return Li(g.Raw(colorname))
				})),
			),
		),
	)
}
