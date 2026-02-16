package golang

import (
	"bytes"
	"html"
)

func Header(bb *bytes.Buffer, title *string) {
	_, _ = bb.WriteString(`<title>`)
	_, _ = bb.WriteString(html.EscapeString(*title))
	_, _ = bb.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)
}
