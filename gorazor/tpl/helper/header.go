package helper

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Header(title string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<title>")
	_buffer.WriteString(gorazor.HTMLEscape(title))
	_buffer.WriteString("'s Home Page</title>\n<div class=\"header\">Page Header</div>")

	return _buffer.String()
}
