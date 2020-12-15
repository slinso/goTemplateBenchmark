package golang

import (
	"html"

	"github.com/valyala/bytebufferpool"
)

func Header(bb *bytebufferpool.ByteBuffer, title *string) {
	_, _ = bb.WriteString(`<title>`)
	_, _ = bb.WriteString(html.EscapeString(*title))
	_, _ = bb.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
`)
}
