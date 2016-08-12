// Package ftmpl is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
package ftmpl

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRsimple evaluates a template simple.tmpl
func TMPLERRsimple(u *model.User) (string, error) {
	_template := "simple.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<html>
    <body>
        <h1>`)
	_w(fmt.Sprintf(`%s`, _escape(u.FirstName)))
	_w(`</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        `)
	for _, colorName := range u.FavoriteColors {
		_w(`
            <li>`)
		_w(fmt.Sprintf(`%s`, _escape(colorName)))
		_w(`</li>`)
	}
	_w(`
        </ul>
    </body>
</html>`)

	return _ftmpl.String(), nil
}

// TMPLsimple evaluates a template simple.tmpl
func TMPLsimple(u *model.User) string {
	html, err := TMPLERRsimple(u)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template simple.tmpl:" + err.Error())
	}
	return html
}
