package layout

import (
	"bytes"
)

func Base(body string, header string, footer string, navigation string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!DOCTYPE html>\n<html>\n<body>\n\n<header>")
	_buffer.WriteString((header))
	_buffer.WriteString("\n</header>\n\n<nav>")
	_buffer.WriteString((navigation))
	_buffer.WriteString("\n</nav>\n\n<section>")
	_buffer.WriteString((body))
	_buffer.WriteString("\n</section>\n\n<footer>")
	_buffer.WriteString((footer))
	_buffer.WriteString("\n</footer>\n\n</body>\n</html>")

	return _buffer.String()
}
