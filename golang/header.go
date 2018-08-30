package golang

import (
	"html"

	"github.com/valyala/bytebufferpool"
)

func Header(bb *bytebufferpool.ByteBuffer, title *string) {
	bb.WriteString(`<title>`)
	bb.WriteString(html.EscapeString(*title))
	bb.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)
}
