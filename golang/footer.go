package golang

import (
	"bytes"
)

func WriteFooter(bb *bytes.Buffer) {
	_, _ = bb.WriteString(`<div class="footer">copyright 2016</div>`)
}
