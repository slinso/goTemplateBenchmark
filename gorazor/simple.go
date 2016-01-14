package gorazor

import (
	"bytes"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/sipin/gorazor/gorazor"
)

func Simple(u *model.User) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<html>\n    <body>\n        <h1>")
	_buffer.WriteString(gorazor.HTMLEscape(u.FirstName))
	_buffer.WriteString("</h1>\n\n        <p>Here's a list of your favorite colors:</p>\n        <ul>\n        ")
	for _, colorName := range u.FavoriteColors {

		_buffer.WriteString("<li>")
		_buffer.WriteString(gorazor.HTMLEscape(colorName))
		_buffer.WriteString("</li>")
	}

	_buffer.WriteString("\n        </ul>\n    </body>\n</html>")

	return _buffer.String()
}
