package gorazor

import (
	"bytes"
	"github.com/SlinSo/goTemplateBenchmark/gorazor/tpl/helper"
	"github.com/SlinSo/goTemplateBenchmark/gorazor/tpl/layout"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/sipin/gorazor/gorazor"
)

func Index(u *model.User, nav []*model.Navigation, title string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\n<div class=\"content\">\n\t<div class=\"welcome\">\n\t\t<h4>Hello ")
	_buffer.WriteString(gorazor.HTMLEscape(u.FirstName))
	_buffer.WriteString("</h4>\n\t\t\n\t\t<div class=\"raw\">")
	_buffer.WriteString((u.RawContent))
	_buffer.WriteString("</div>\n\t\t<div class=\"enc\">")
	_buffer.WriteString(gorazor.HTMLEscape(u.EscapedContent))
	_buffer.WriteString("</div>\n\t</div>")

	for i := 1; i <= 5; i++ {
		if i == 1 {

			_buffer.WriteString("<p>")
			_buffer.WriteString(gorazor.HTMLEscape(u.FirstName))
			_buffer.WriteString(" has ")
			_buffer.WriteString(gorazor.HTMLEscape(i))
			_buffer.WriteString(" message</p>")

		} else {

			_buffer.WriteString("<p>")
			_buffer.WriteString(gorazor.HTMLEscape(u.FirstName))
			_buffer.WriteString(" has ")
			_buffer.WriteString(gorazor.HTMLEscape(gorazor.Itoa(i)))
			_buffer.WriteString(" messages</p>")

		}
	}

	_buffer.WriteString("\n</div>")

	header := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString((helper.Header(title)))

		return _buffer.String()
	}

	footer := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString((helper.Footer()))

		return _buffer.String()
	}

	navigation := func() string {
		var _buffer bytes.Buffer

		_buffer.WriteString((helper.Navigation(nav)))

		return _buffer.String()
	}

	return layout.Base(_buffer.String(), header(), footer(), navigation())
}
