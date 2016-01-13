package ftmpl

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"html"
	"os"
)

func init() {
	_ = fmt.Sprintf
	_ = errors.New
	_ = os.Stderr
	_ = html.EscapeString
}

// Generated code, do not edit!!!!
func TE__simple(u *model.User) (string, error) {
	__template__ := "simple.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <html> */
	result.WriteString(`<html>
`)
	/* <body> */
	result.WriteString(`    <body>
`)
	/* <h1>{{s u.FirstName }}</h1> */
	result.WriteString(fmt.Sprintf(`        <h1>%s</h1>
`, __escape__(u.FirstName)))
	/*  */
	result.WriteString(`
`)
	/* <p>Here's a list of your favorite colors:</p> */
	result.WriteString(`        <p>Here's a list of your favorite colors:</p>
`)
	/* <ul> */
	result.WriteString(`        <ul>
`)
	/*  */
	result.WriteString(`        `)
	/* !for _, colorName := range u.FavoriteColors{ */
	for _, colorName := range u.FavoriteColors {
		/*  */
		result.WriteString(`
`)
		/* <li>{{s colorName }}</li> */
		result.WriteString(fmt.Sprintf(`            <li>%s</li>`, __escape__(colorName)))
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* </ul> */
	result.WriteString(`        </ul>
`)
	/* </body> */
	result.WriteString(`    </body>
`)
	/* </html> */
	result.WriteString(`</html>`)

	return result.String(), nil
}

func T__simple(u *model.User) string {
	html, err := TE__simple(u)
	if err != nil {
		os.Stderr.WriteString("Error running template simple.tmpl:" + err.Error())
	}
	return html
}
