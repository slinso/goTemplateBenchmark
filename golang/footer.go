package golang

import (
	"github.com/valyala/bytebufferpool"
)

func WriteFooter(bb *bytebufferpool.ByteBuffer) {
	_, _ = bb.WriteString(`<div class="footer">copyright 2016</div>`)
}
