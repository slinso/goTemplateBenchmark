package helper

import (
	"bytes"
)

func Footer() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("<div class=\"footer\">copyright 2016</div>")

	return _buffer.String()
}
