// DO NOT EDIT!
// Generate By Goh

package template

import (
	"bytes"

	"github.com/SlinSo/goTemplateBenchmark/model"

	Goh "github.com/OblivionOcean/Goh/utils"
)

func Index(u *model.User, nav []*model.Navigation, title string, buffer *bytes.Buffer) {
	buffer.Grow(239)
	buffer.WriteString(`<title>`)
	Goh.EscapeHTML(title, buffer)
	buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
<title>age</title>
<div class="header">Page Header</div>
`)
	Goh.EscapeHTML(title, buffer)
	buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>
<title>age</title>
<div class="header">Page Header</div>
`)
}
