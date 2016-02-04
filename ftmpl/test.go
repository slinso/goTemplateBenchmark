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
func TE__base(u *model.User, nav []*model.Navigation, title string) (string, error) {
	__template__ := "base.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <!DOCTYPE html> */
	result.WriteString(`<!DOCTYPE html>
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <body> */
	result.WriteString(`<body>
`)
	/*  */
	result.WriteString(`
`)
	/* <header> */
	result.WriteString(`<header>
`)
	/* !#include header */
	/* </header> */
	result.WriteString(`</header>
`)
	/*  */
	result.WriteString(`
`)
	/* <nav> */
	result.WriteString(`<nav>
`)
	/* !#include navigation */
	/* </nav> */
	result.WriteString(`</nav>
`)
	/*  */
	result.WriteString(`
`)
	/* <section> */
	result.WriteString(`<section>
`)
	/* !#include body */
	/* </section> */
	result.WriteString(`</section>
`)
	/*  */
	result.WriteString(`
`)
	/* <footer> */
	result.WriteString(`<footer>
`)
	/* !#include footer */
	/* </footer> */
	result.WriteString(`</footer>
`)
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`</body>
`)
	/* </html> */
	result.WriteString(`</html>`)

	return result.String(), nil
}

func T__base(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TE__base(u, nav, title)
	if err != nil {
		os.Stderr.WriteString("Error running template base.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__basecontent(u *model.User, nav []*model.Navigation, title string) (string, error) {
	__template__ := "basecontent.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* !#extends base */
	/*  */
	result.WriteString(`
`)
	/* <!DOCTYPE html> */
	result.WriteString(`<!DOCTYPE html>
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <body> */
	result.WriteString(`<body>
`)
	/*  */
	result.WriteString(`
`)
	/* <header> */
	result.WriteString(`<header>
`)
	/* <title>{{s title}}'s Home Page</title> */
	result.WriteString(fmt.Sprintf(`<title>%s's Home Page</title>
`, __escape__(title)))
	/* <div class="header">Page Header</div> */
	result.WriteString(`<div class="header">Page Header</div>
`)
	/*  */
	result.WriteString(`
`)
	/* </header> */
	result.WriteString(`</header>
`)
	/*  */
	result.WriteString(`
`)
	/* <nav> */
	result.WriteString(`<nav>
`)
	/* <ul class="navigation"> */
	result.WriteString(`<ul class="navigation">
`)
	/*  */
	result.WriteString(`    `)
	/* !for _, item := range nav{ */
	for _, item := range nav {
		/*  */
		result.WriteString(`
`)
		/* <li><a href="{{s item.Link }}">{{s item.Item }}</a></li> */
		result.WriteString(fmt.Sprintf(`        	<li><a href="%s">%s</a></li>
`, __escape__(item.Link), __escape__(item.Item)))
		/*  */
		result.WriteString(`    `)
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* </ul> */
	result.WriteString(`</ul>`)
	/* </nav> */
	result.WriteString(`</nav>
`)
	/*  */
	result.WriteString(`
`)
	/* <section> */
	result.WriteString(`<section>
`)
	/* </section> */
	result.WriteString(`</section>
`)
	/*  */
	result.WriteString(`
`)
	/* <footer> */
	result.WriteString(`<footer>
`)
	/* <div class="footer">copyright 2016</div> */
	result.WriteString(`<div class="footer">copyright 2016</div>
`)
	/*  */
	result.WriteString(`
`)
	/* </footer> */
	result.WriteString(`</footer>
`)
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`</body>
`)
	/* </html> */
	result.WriteString(`</html>`)

	return result.String(), nil
}

func T__basecontent(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TE__basecontent(u, nav, title)
	if err != nil {
		os.Stderr.WriteString("Error running template basecontent.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__footer() (string, error) {
	__template__ := "footer.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <div class="footer">copyright 2016</div> */
	result.WriteString(`<div class="footer">copyright 2016</div>`)

	return result.String(), nil
}

func T__footer() string {
	html, err := TE__footer()
	if err != nil {
		os.Stderr.WriteString("Error running template footer.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__header(title string) (string, error) {
	__template__ := "header.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <title>{{s title}}'s Home Page</title> */
	result.WriteString(fmt.Sprintf(`<title>%s's Home Page</title>
`, __escape__(title)))
	/* <div class="header">Page Header</div> */
	result.WriteString(`<div class="header">Page Header</div>`)

	return result.String(), nil
}

func T__header(title string) string {
	html, err := TE__header(title)
	if err != nil {
		os.Stderr.WriteString("Error running template header.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__index(u *model.User, nav []*model.Navigation, title string) (string, error) {
	__template__ := "index.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* !#extends base */
	/*  */
	result.WriteString(`
`)
	/* <!DOCTYPE html> */
	result.WriteString(`<!DOCTYPE html>
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <body> */
	result.WriteString(`<body>
`)
	/*  */
	result.WriteString(`
`)
	/* <header> */
	result.WriteString(`<header>
`)
	/* <title>{{s title}}'s Home Page</title> */
	result.WriteString(fmt.Sprintf(`<title>%s's Home Page</title>
`, __escape__(title)))
	/* <div class="header">Page Header</div> */
	result.WriteString(`<div class="header">Page Header</div>
`)
	/*  */
	result.WriteString(`
`)
	/* </header> */
	result.WriteString(`</header>
`)
	/*  */
	result.WriteString(`
`)
	/* <nav> */
	result.WriteString(`<nav>
`)
	/* <ul class="navigation"> */
	result.WriteString(`<ul class="navigation">
`)
	/*  */
	result.WriteString(`    `)
	/* !for _, item := range nav{ */
	for _, item := range nav {
		/*  */
		result.WriteString(`
`)
		/* <li><a href="{{s item.Link }}">{{s item.Item }}</a></li> */
		result.WriteString(fmt.Sprintf(`        	<li><a href="%s">%s</a></li>
`, __escape__(item.Link), __escape__(item.Item)))
		/*  */
		result.WriteString(`    `)
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* </ul> */
	result.WriteString(`</ul>
`)
	/*  */
	result.WriteString(`
`)
	/* </nav> */
	result.WriteString(`</nav>
`)
	/*  */
	result.WriteString(`
`)
	/* <section> */
	result.WriteString(`<section>
`)
	/* <div class="content"> */
	result.WriteString(`<div class="content">
`)
	/* <div class="welcome"> */
	result.WriteString(`	<div class="welcome">
`)
	/* <h4>Hello {{s u.FirstName }}</h4> */
	result.WriteString(fmt.Sprintf(`		<h4>Hello %s</h4>
`, __escape__(u.FirstName)))
	/*  */
	result.WriteString(`		
`)
	/* <div class="raw">{{=s u.RawContent }}</div> */
	result.WriteString(fmt.Sprintf(`		<div class="raw">%s</div>
`, u.RawContent))
	/* <div class="enc">{{s u.EscapedContent }}</div> */
	result.WriteString(fmt.Sprintf(`		<div class="enc">%s</div>
`, __escape__(u.EscapedContent)))
	/* </div> */
	result.WriteString(`	</div>
`)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(`
`)
	/* !	for i := 1; i <= 5; i++ { */
	for i := 1; i <= 5; i++ {

		/* !		if i == 1 { */
		if i == 1 {

			/* <p>{{s u.FirstName}} has {{d i}} message</p> */
			result.WriteString(fmt.Sprintf(`			<p>%s has %d message</p>
`, __escape__(u.FirstName), i))
			/* !		} else { */
		} else {

			/* <p>{{s u.FirstName}} has {{d i}} messages</p> */
			result.WriteString(fmt.Sprintf(`			<p>%s has %d messages</p>
`, __escape__(u.FirstName), i))
			/* !		} */
		}

		/* !	} */
	}

	/* </div> */
	result.WriteString(`</div>
`)
	/*  */
	result.WriteString(``)
	/* </section> */
	result.WriteString(`</section>
`)
	/*  */
	result.WriteString(`
`)
	/* <footer> */
	result.WriteString(`<footer>
`)
	/* <div class="footer">copyright 2016</div> */
	result.WriteString(`<div class="footer">copyright 2016</div>
`)
	/*  */
	result.WriteString(`
`)
	/* </footer> */
	result.WriteString(`</footer>
`)
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`</body>
`)
	/* </html> */
	result.WriteString(`</html>`)

	return result.String(), nil
}

func T__index(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TE__index(u, nav, title)
	if err != nil {
		os.Stderr.WriteString("Error running template index.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__index2(u *model.User, nav []*model.Navigation, title string) (string, error) {
	__template__ := "index2.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* !#extends base */
	/*  */
	result.WriteString(`
`)
	/* <!DOCTYPE html> */
	result.WriteString(`<!DOCTYPE html>
`)
	/* <html> */
	result.WriteString(`<html>
`)
	/* <body> */
	result.WriteString(`<body>
`)
	/*  */
	result.WriteString(`
`)
	/* <header> */
	result.WriteString(`<header>
`)
	/* ! result.WriteString(T__header(title)) */
	result.WriteString(T__header(title))

	/*  */
	result.WriteString(`
`)
	/* </header> */
	result.WriteString(`</header>
`)
	/*  */
	result.WriteString(`
`)
	/* <nav> */
	result.WriteString(`<nav>
`)
	/* ! result.WriteString(T__navigation(nav)) */
	result.WriteString(T__navigation(nav))

	/*  */
	result.WriteString(`
`)
	/* </nav> */
	result.WriteString(`</nav>
`)
	/*  */
	result.WriteString(`
`)
	/* <section> */
	result.WriteString(`<section>
`)
	/* <div class="content"> */
	result.WriteString(`<div class="content">
`)
	/* <div class="welcome"> */
	result.WriteString(`	<div class="welcome">
`)
	/* <h4>Hello {{s u.FirstName }}</h4> */
	result.WriteString(fmt.Sprintf(`		<h4>Hello %s</h4>
`, __escape__(u.FirstName)))
	/*  */
	result.WriteString(`		
`)
	/* <div class="raw">{{=s u.RawContent }}</div> */
	result.WriteString(fmt.Sprintf(`		<div class="raw">%s</div>
`, u.RawContent))
	/* <div class="enc">{{s u.EscapedContent }}</div> */
	result.WriteString(fmt.Sprintf(`		<div class="enc">%s</div>
`, __escape__(u.EscapedContent)))
	/* </div> */
	result.WriteString(`	</div>
`)
	/*  */
	result.WriteString(`
`)
	/*  */
	result.WriteString(`
`)
	/* !	for i := 1; i <= 5; i++ { */
	for i := 1; i <= 5; i++ {

		/* !		if i == 1 { */
		if i == 1 {

			/* <p>{{s u.FirstName}} has {{d i}} message</p> */
			result.WriteString(fmt.Sprintf(`			<p>%s has %d message</p>
`, __escape__(u.FirstName), i))
			/* !		} else { */
		} else {

			/* <p>{{s u.FirstName}} has {{d i}} messages</p> */
			result.WriteString(fmt.Sprintf(`			<p>%s has %d messages</p>
`, __escape__(u.FirstName), i))
			/* !		} */
		}

		/* !	} */
	}

	/* </div> */
	result.WriteString(`</div>
`)
	/*  */
	result.WriteString(``)
	/* </section> */
	result.WriteString(`</section>
`)
	/*  */
	result.WriteString(`
`)
	/* <footer> */
	result.WriteString(`<footer>
`)
	/* ! result.WriteString(T__footer()) */
	result.WriteString(T__footer())

	/*  */
	result.WriteString(`
`)
	/* </footer> */
	result.WriteString(`</footer>
`)
	/*  */
	result.WriteString(`
`)
	/* </body> */
	result.WriteString(`</body>
`)
	/* </html> */
	result.WriteString(`</html>`)

	return result.String(), nil
}

func T__index2(u *model.User, nav []*model.Navigation, title string) string {
	html, err := TE__index2(u, nav, title)
	if err != nil {
		os.Stderr.WriteString("Error running template index2.tmpl:" + err.Error())
	}
	return html
}

// Generated code, do not edit!!!!
func TE__navigation(nav []*model.Navigation) (string, error) {
	__template__ := "navigation.tmpl"
	_ = __template__
	__escape__ := html.EscapeString
	_ = __escape__
	var result bytes.Buffer
	/* <ul class="navigation"> */
	result.WriteString(`<ul class="navigation">
`)
	/*  */
	result.WriteString(`    `)
	/* !for _, item := range nav{ */
	for _, item := range nav {
		/*  */
		result.WriteString(`
`)
		/* <li><a href="{{s item.Link }}">{{s item.Item }}</a></li> */
		result.WriteString(fmt.Sprintf(`        	<li><a href="%s">%s</a></li>
`, __escape__(item.Link), __escape__(item.Item)))
		/*  */
		result.WriteString(`    `)
		/* !} */
	}
	/*  */
	result.WriteString(`
`)
	/* </ul> */
	result.WriteString(`</ul>`)

	return result.String(), nil
}

func T__navigation(nav []*model.Navigation) string {
	html, err := TE__navigation(nav)
	if err != nil {
		os.Stderr.WriteString("Error running template navigation.tmpl:" + err.Error())
	}
	return html
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
