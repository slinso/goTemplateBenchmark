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

// TMPLERRbase evaluates a template base.tmpl
func TMPLERRbase(u *model.User, nav []*model.Navigation, title string) (string, error) {
	_template := "base.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<!DOCTYPE html>
<html>
<body>

<header>
`)
	_w(`</header>

<nav>
`)
	_w(`</nav>

<section>
`)
	_w(`</section>

<footer>
`)
	_w(`</footer>

</body>
</html>`)

	return _ftmpl.String(), nil
}

// TMPLbase evaluates a template base.tmpl
func TMPLbase(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TMPLERRbase(u, nav, title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}
