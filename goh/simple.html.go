// DO NOT EDIT!
// Generate By Goh

package template

import (
	"bytes"

	"github.com/SlinSo/goTemplateBenchmark/model"

	Goh "github.com/OblivionOcean/Goh/utils"
)

func SimpleQtc(u *model.User, buffer *bytes.Buffer) {
	buffer.Grow(194)
	buffer.WriteString(`



<html>
    <body>
        <h1>`)
	Goh.EscapeHTML(u.FirstName, buffer)
	buffer.WriteString(`</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
            `)
	for _, colorName := range u.FavoriteColors {
		buffer.WriteString(`
                <li>`)
		Goh.EscapeHTML(colorName, buffer)
		buffer.WriteString(`</li>
            `)
	}
	buffer.WriteString(`
        </ul>
    </body>
</html>
`)
}
