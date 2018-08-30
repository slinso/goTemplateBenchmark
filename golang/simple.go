package golang

import (
	"html"
	"net/http"

	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/valyala/bytebufferpool"
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
	bb *bytebufferpool.ByteBuffer
}

func NewGoFunc(buf *bytebufferpool.ByteBuffer) *GoFunc {
	return &GoFunc{bb: buf}
}

type hyperfunc func(string, hyperfunc)

func (g *GoFunc) H(elem string, fn func(*string)) func(*string) {
	return func(s *string) {

		g.bb.WriteString("<")
		g.bb.WriteString(elem)
		g.bb.WriteString(">")

		if fn != nil {
			fn(s)
		}

		g.bb.WriteString("</")
		g.bb.WriteString(elem)
		g.bb.WriteString(">")
	}
}

type Attr map[string]string

func (g *GoFunc) html() {
	g.bb.WriteString("<html>")
}
func (g *GoFunc) htmlEnd() {
	g.bb.WriteString("</html>")
}

func (g *GoFunc) body() {
	g.bb.WriteString("<body>")
}
func (g *GoFunc) bodyEnd() {
	g.bb.WriteString("</body>")
}

func (g *GoFunc) h1() {
	g.bb.WriteString("<h1>")
}
func (g *GoFunc) h1End() {
	g.bb.WriteString("</h1>")
}
func (g *GoFunc) p() {
	g.bb.WriteString("<p>")
}
func (g *GoFunc) pEnd() {
	g.bb.WriteString("</p>")
}
func (g *GoFunc) li() {
	g.bb.WriteString("<li>")
}
func (g *GoFunc) liEnd() {
	g.bb.WriteString("</li>")
}
func (g *GoFunc) ul() {
	g.bb.WriteString("<ul>")
}
func (g *GoFunc) ulEnd() {
	g.bb.WriteString("</ul>")
}
func (g *GoFunc) escape(s string) {
	g.bb.WriteString(html.EscapeString(s))
}
func (g *GoFunc) s(s string) {
	g.bb.WriteString(s)
}
func (g *GoFunc) elem(s []string) {
	for _, e := range s {
		g.bb.WriteString("<")
		g.bb.WriteString(e)
		g.bb.WriteString(">")
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
func WriteSimpleGolang(bb *bytebufferpool.ByteBuffer, u *model.User) {
	bb.WriteString(`
<html>
    <body>
        <h1>`)
	bb.WriteString(html.EscapeString(u.FirstName))
	bb.WriteString(`</h1>
        <p>Here's a list of your favorite colors:</p>
        <ul>
        `)
	for _, colorName := range u.FavoriteColors {
		bb.WriteString(`<li>`)
		bb.WriteString(html.EscapeString(colorName))
		bb.WriteString(`</li>`)
	}
	bb.WriteString(`
        </ul>
    </body>
</html>
`)
}
