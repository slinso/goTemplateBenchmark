// Package ftmpl is generated with ftmpl {{{v0.3.1}}}, do not edit!!!! */
package ftmpl

import (
	"bytes"
	"errors"
	"fmt"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// TMPLERRheader evaluates a template header.tmpl
func TMPLERRheader(title string) (string, error) {
	_template := "header.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<title>`)
	_w(fmt.Sprintf(`%s`, _escape(title)))
	_w(`'s Home Page</title>
<div class="header">Page Header</div>`)

	return _ftmpl.String(), nil
}

// TMPLheader evaluates a template header.tmpl
func TMPLheader(title string) string {
	html, err := TMPLERRheader(title)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template header.tmpl:" + err.Error())
	}
	return html
}
