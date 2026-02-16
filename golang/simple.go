package golang

import (
	"bytes"
	"html"
	"net/http"

	"github.com/SlinSo/goTemplateBenchmark/model"
)

type TemplateFunc func(string)

func (f TemplateFunc) H(s string) {
	f(s)
}

type TemplateH interface {
	H(string)
}
type Me http.Handler
type Me2 http.HandlerFunc

type GoFunc struct {
	bb *bytes.Buffer
}

func NewGoFunc(buf *bytes.Buffer) *GoFunc {
	return &GoFunc{bb: buf}
}

func (g *GoFunc) H(elem string, fn func(*string)) func(*string) {
	return func(s *string) {

		_, _ = g.bb.WriteString("<")
		_, _ = g.bb.WriteString(elem)
		_, _ = g.bb.WriteString(">")

		if fn != nil {
			fn(s)
		}

		_, _ = g.bb.WriteString("</")
		_, _ = g.bb.WriteString(elem)
		_, _ = g.bb.WriteString(">")
	}
}

type Attr map[string]string

func (g *GoFunc) html() {
	_, _ = g.bb.WriteString("<html>")
}
func (g *GoFunc) htmlEnd() {
	_, _ = g.bb.WriteString("</html>")
}

func (g *GoFunc) body() {
	_, _ = g.bb.WriteString("<body>")
}
func (g *GoFunc) bodyEnd() {
	_, _ = g.bb.WriteString("</body>")
}

func (g *GoFunc) h1() {
	_, _ = g.bb.WriteString("<h1>")
}
func (g *GoFunc) h1End() {
	_, _ = g.bb.WriteString("</h1>")
}
func (g *GoFunc) p() {
	_, _ = g.bb.WriteString("<p>")
}
func (g *GoFunc) pEnd() {
	_, _ = g.bb.WriteString("</p>")
}
func (g *GoFunc) li() {
	_, _ = g.bb.WriteString("<li>")
}
func (g *GoFunc) liEnd() {
	_, _ = g.bb.WriteString("</li>")
}
func (g *GoFunc) ul() {
	_, _ = g.bb.WriteString("<ul>")
}
func (g *GoFunc) ulEnd() {
	_, _ = g.bb.WriteString("</ul>")
}
func (g *GoFunc) escape(s string) {
	_, _ = g.bb.WriteString(html.EscapeString(s))
}
func (g *GoFunc) s(s string) {
	_, _ = g.bb.WriteString(s)
}
func (g *GoFunc) elem(s []string) {
	for _, e := range s {
		_, _ = g.bb.WriteString("<")
		_, _ = g.bb.WriteString(e)
		_, _ = g.bb.WriteString(">")
	}
}

func GoFuncElem(g *GoFunc, u *model.User) {
	g.elem([]string{"html", "body", "h1"})
	g.escape(u.FirstName)
	g.h1End()
	g.p()
	g.s("Here's a list of your favorite colors:")
	g.pEnd()
	g.ul()
	for _, colorName := range u.FavoriteColors {
		g.li()
		g.escape(colorName)
		g.liEnd()
	}
	g.ulEnd()
	g.bodyEnd()
	g.htmlEnd()
}

func GoFuncFunc(g *GoFunc, u *model.User) {
	g.html()
	g.body()
	g.h1()
	g.escape(u.FirstName)
	g.h1End()
	g.p()
	g.s("Here's a list of your favorite colors:")
	g.pEnd()
	g.ul()
	for _, colorName := range u.FavoriteColors {
		g.li()
		g.escape(colorName)
		g.liEnd()
	}
	g.ulEnd()
	g.bodyEnd()
	g.htmlEnd()
}

// WriteSimpleGolang golang funcion based template
func WriteSimpleGolang(bb *bytes.Buffer, u *model.User) {
	_, _ = bb.WriteString(`
<html>
    <body>
        <h1>`)
	_, _ = bb.WriteString(html.EscapeString(u.FirstName))
	_, _ = bb.WriteString(`</h1>
        <p>Here's a list of your favorite colors:</p>
        <ul>
        `)
	for _, colorName := range u.FavoriteColors {
		_, _ = bb.WriteString(`<li>`)
		_, _ = bb.WriteString(html.EscapeString(colorName))
		_, _ = bb.WriteString(`</li>`)
	}
	_, _ = bb.WriteString(`
        </ul>
    </body>
</html>
`)
}
