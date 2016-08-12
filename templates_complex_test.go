package main_test

import (
	"bytes"
	"html/template"
	"path/filepath"
	"testing"

	"github.com/SlinSo/goTemplateBenchmark/model"

	"github.com/SlinSo/goTemplateBenchmark/ego"
	"github.com/SlinSo/goTemplateBenchmark/egon"
	"github.com/SlinSo/goTemplateBenchmark/egonslinso"
	"github.com/SlinSo/goTemplateBenchmark/ftmpl"
	"github.com/SlinSo/goTemplateBenchmark/gorazor"
	"github.com/SlinSo/goTemplateBenchmark/quicktemplate"
	"github.com/hoisie/mustache"
)

type tmplData struct {
	User     *model.User
	Nav      []*model.Navigation
	Title    string
	Messages []struct {
		I      int
		Plural bool
	}
}

var (
	testComplexUser = &model.User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
		RawContent:     "<div><p>Raw Content to be displayed</p></div>",
		EscapedContent: "<div><div><div>Escaped</div></div></div>",
	}

	testComplexNav = []*model.Navigation{{
		Item: "Link 1",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 2",
		Link: "http://www.mytest.com/"}, {
		Item: "Link 3",
		Link: "http://www.mytest.com/"},
	}
	testComplexTitle = testComplexUser.FirstName

	testComplexData = tmplData{

		User:  testComplexUser,
		Nav:   testComplexNav,
		Title: testComplexTitle,
		Messages: []struct {
			I      int
			Plural bool
		}{{1, false}, {2, true}, {3, true}, {4, true}, {5, true}},
	}

	expectedtComplexResult = `<!DOCTYPE html>
<html>
<body>

<header><title>Bob's Home Page</title>
<div class="header">Page Header</div>
</header>

<nav>
<ul class="navigation"><li><a href="http://www.mytest.com/">Link 1</a></li>
<li><a href="http://www.mytest.com/">Link 2</a></li>
<li><a href="http://www.mytest.com/">Link 3</a></li>
</ul>
</nav>

<section>

<div class="content">
        <div class="welcome">
                <h4>Hello Bob</h4>

                <div class="raw"><div><p>Raw Content to be displayed</p></div></div>
                <div class="enc">&lt;div&gt;&lt;div&gt;&lt;div&gt;Escaped&lt;/div&gt;&lt;/div&gt;&lt;/div&gt;</div>
        </div><p>Bob has 1 message</p><p>Bob has 2 messages</p><p>Bob has 3 messages</p><p>Bob has 4 messages</p><p>Bob has 5 messages</p>
</div>
</section>

<footer><div class="footer">copyright 2016</div>
</footer>

</body>
</html>`

	expectedtComplexResultMinified = "<html><body><h1>Bob</h1><p>Here's a list of your favorite colors:</p><ul><li>blue</li><li>green</li><li>mauve</li></ul></body></html>"
)

/******************************************************************************
** Go
******************************************************************************/
func TestComplexGolang(t *testing.T) {

	var buf bytes.Buffer

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}

	templates := make(map[string]*template.Template)
	templatesDir := "go/"

	layouts, err := filepath.Glob(templatesDir + "layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = template.Must(template.New("").Funcs(funcMap).ParseFiles(files...))
	}
	templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexGolang(b *testing.B) {
	var buf bytes.Buffer

	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}

	templates := make(map[string]*template.Template)
	templatesDir := "go/"

	layouts, err := filepath.Glob(templatesDir + "layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = template.Must(template.New("").Funcs(funcMap).ParseFiles(files...))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)
	}
}

/******************************************************************************
** Ego
******************************************************************************/
func TestComplexEgo(t *testing.T) {
	var buf bytes.Buffer
	ego.EgoComplex(&buf, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexEgo(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		ego.EgoComplex(&buf, testComplexUser, testComplexNav, testComplexTitle)
	}
}

/******************************************************************************
** Quicktemplate
******************************************************************************/
func TestComplexQuicktemplate(t *testing.T) {
	var buf bytes.Buffer
	quicktemplate.WriteIndex(&buf, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexQuicktemplate(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		quicktemplate.WriteIndex(&buf, testComplexUser, testComplexNav, testComplexTitle)
	}
}

/******************************************************************************
** Egon
******************************************************************************/
func TestComplexEgon(t *testing.T) {
	var buf bytes.Buffer
	egon.IndexTemplate(&buf, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexEgon(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egon.IndexTemplate(&buf, testComplexUser, testComplexNav, testComplexTitle)
	}
}

/******************************************************************************
** EgoSlinso
******************************************************************************/
func TestComplexEgoSlinso(t *testing.T) {
	var buf bytes.Buffer
	egonslinso.IndexTemplate(&buf, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexEgoSlinso(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egonslinso.IndexTemplate(&buf, testComplexUser, testComplexNav, testComplexTitle)
	}
}

/******************************************************************************
** ftmpl
******************************************************************************/
func TestComplexFtmpl(t *testing.T) {
	result := ftmpl.TMPLindex(testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(result, expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexFtmpl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ftmpl.TMPLindex(testComplexUser, testComplexNav, testComplexTitle)
	}
}

func TestComplexFtmplInclude(t *testing.T) {
	result := ftmpl.TMPLindex2(testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(result, expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexFtmplInclude(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ftmpl.TMPLindex2(testComplexUser, testComplexNav, testComplexTitle)
	}
}

/******************************************************************************
** Mustache
******************************************************************************/
func TestComplexMustache(t *testing.T) {
	layoutTmpl, err := mustache.ParseFile("mustache/base.mustache")
	if err != nil {
		t.Error(err)
	}
	tmpl, err := mustache.ParseFile("mustache/index.mustache")
	if err != nil {
		t.Error(err)
	}

	result := tmpl.RenderInLayout(layoutTmpl, testComplexData)

	if msg, ok := linesEquals(result, expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexMustache(b *testing.B) {
	layoutTmpl, _ := mustache.ParseFile("mustache/base.mustache")
	tmpl, _ := mustache.ParseFile("mustache/index.mustache")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmpl.RenderInLayout(layoutTmpl, testComplexData)
	}
}

/******************************************************************************
** gorazor
******************************************************************************/
func TestComplexGorazor(t *testing.T) {
	result := gorazor.Index(testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(result, expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexGorazor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gorazor.Index(testComplexUser, testComplexNav, testComplexTitle)
	}
}
