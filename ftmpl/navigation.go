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

// TMPLERRnavigation evaluates a template navigation.tmpl
func TMPLERRnavigation(nav []*model.Navigation) (string, error) {
	_template := "navigation.tmpl"
	_escape := html.EscapeString
	var _ftmpl bytes.Buffer
	_w := func(str string) { _, _ = _ftmpl.WriteString(str) }
	_, _, _ = _template, _escape, _w

	_w(`<ul class="navigation">
    `)
	for _, item := range nav {
		_w(`
        	<li><a href="`)
		_w(fmt.Sprintf(`%s`, _escape(item.Link)))
		_w(`">`)
		_w(fmt.Sprintf(`%s`, _escape(item.Item)))
		_w(`</a></li>
    `)
	}
	_w(`
</ul>`)

	return _ftmpl.String(), nil
}

// TMPLnavigation evaluates a template navigation.tmpl
func TMPLnavigation(nav []*model.Navigation) string {
	html, err := TMPLERRnavigation(nav)
	if err != nil {
		_, _ = os.Stderr.WriteString("Error running template navigation.tmpl:" + err.Error())
	}
	return html
}
