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

// TMPLERRfooter evaluates a template footer.tmpl
func TMPLERRfooter() (string, error) {
	_template := "footer.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<div class="footer">copyright 2016</div>`)

	return _ftmpl.String(), nil
}

// TMPLfooter evaluates a template footer.tmpl
func TMPLfooter() string {
	html, err := TMPLERRfooter()
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template footer.tmpl:" + err.Error())
	}
	return html
}
