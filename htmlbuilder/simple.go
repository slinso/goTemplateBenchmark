package htmlbuilder

import (
	"context"
	"fmt"
	"html"
	"io"

	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/gouniverse/hb"
)

func HbSimple(w io.Writer, u *model.User) {
	ul := hb.NewUL()
	for i := 0; i < len(u.FavoriteColors); i++ {
		ul.Child(hb.NewLI().HTML(u.FavoriteColors[i]))
	}

	page := hb.NewTag("html").
		Child(hb.NewTag("body").
			Child(hb.NewHeading1().HTML(u.FirstName)).
			Child(hb.NewParagraph().HTML("Here's a list of your favorite colors:")).
			Child(ul),
		).ToHTML()

	_, _ = io.WriteString(w, page)
}

var (
	_ fmt.Stringer
	_ io.Reader
	_ context.Context
	_ = html.EscapeString
)
