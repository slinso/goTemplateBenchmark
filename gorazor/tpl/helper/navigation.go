package helper

import (
	"bytes"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/sipin/gorazor/gorazor"
)

func Navigation(nav []*model.Navigation) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<ul class=\"navigation\">")

	for _, item := range nav {

		_buffer.WriteString("<li><a href=\"")
		_buffer.WriteString(gorazor.HTMLEscape(item.Link))
		_buffer.WriteString("\">")
		_buffer.WriteString(gorazor.HTMLEscape(item.Item))
		_buffer.WriteString("</a></li>")

	}

	_buffer.WriteString("\n</ul>")

	return _buffer.String()
}
