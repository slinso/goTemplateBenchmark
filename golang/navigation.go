package golang

import (
	"bytes"
	"html"

	"github.com/SlinSo/goTemplateBenchmark/model"
)

func Navigation(bb *bytes.Buffer, nav []*model.Navigation) {
	_, _ = bb.WriteString(`<ul class="navigation">`)
	for _, item := range nav {
		_, _ = bb.WriteString(`<li><a href="`)
		_, _ = bb.WriteString(html.EscapeString(item.Link))
		_, _ = bb.WriteString(`">`)
		_, _ = bb.WriteString(html.EscapeString(item.Item))
		_, _ = bb.WriteString(`</a></li>`)

	}
	_, _ = bb.WriteString(`</ul>`)
}
