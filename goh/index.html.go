// DO NOT EDIT!
// Generate By Goh

package template

import (
	"bytes"

	Goh "github.com/OblivionOcean/Goh/utils"
	"github.com/SlinSo/goTemplateBenchmark/model"
)

func Index(u *model.User, nav []*model.Navigation, title string, buffer *bytes.Buffer) {
	buffer.Grow(543)
	buffer.WriteString(`



<!DOCTYPE html>
<html>
<body>

<header>
    <title>`)
	Goh.EscapeHTML(title, buffer)
	buffer.WriteString(`'s Home Page</title>
<div class="header">Page Header</div>

</header>

<nav>
    <ul class="navigation">
`)
	for _, item := range nav {
		buffer.WriteString(`
	<li><a href="`)
		Goh.EscapeHTML(item.Link, buffer)
		buffer.WriteString(`">`)
		Goh.EscapeHTML(item.Item, buffer)
		buffer.WriteString(`</a></li>
`)
	}
	buffer.WriteString(`
</ul>

</nav>

<section>
<div class="content">
	<div class="welcome">
		<h4>Hello `)
	Goh.EscapeHTML(u.FirstName, buffer)
	buffer.WriteString(`</h4>

		<div class="raw">`)
	buffer.WriteString(u.RawContent)
	buffer.WriteString(`</div>
		<div class="enc">`)
	Goh.EscapeHTML(u.EscapedContent, buffer)
	buffer.WriteString(`</div>
	</div>

    `)
	for i := 1; i <= 5; i++ {
		buffer.WriteString(`
        `)
		if i == 1 {
			buffer.WriteString(`
            <p>`)
			Goh.EscapeHTML(u.FirstName, buffer)
			buffer.WriteString(` has `)
			Goh.FormatInt(int64(i), buffer)
			buffer.WriteString(` message</p>
        `)
		} else {
			buffer.WriteString(`
            <p>`)
			Goh.EscapeHTML(u.FirstName, buffer)
			buffer.WriteString(` has `)
			Goh.FormatInt(int64(i), buffer)
			buffer.WriteString(` messages</p>
        `)
		}
		buffer.WriteString(`
    `)
	}
	buffer.WriteString(`
</div>
</section>

<footer>
    <div class="footer">copyright 2016</div>

</footer>

</body>
</html>
`)
}
