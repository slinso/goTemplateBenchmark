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

	htmlTag := &hb.Tag{
		TagName: "html",
	}

	bodyTag := &hb.Tag{
		TagName: "body",
	}

	page := htmlTag.Children([]*hb.Tag{
		bodyTag.Children([]*hb.Tag{
			hb.NewHeading1().HTML(u.FirstName),
			hb.NewParagraph().HTML("Here's a list of your favorite colors:"),
			ul,
		}),
	})

	_, _ = io.WriteString(w, page.ToHTML())
}

var _ fmt.Stringer
var _ io.Reader
var _ context.Context
var _ = html.EscapeString
