package main_test

import (
	"bytes"
	"context"
	"html"
	"html/template"
	"path/filepath"
	"testing"
	text "text/template"

	"github.com/SlinSo/goTemplateBenchmark/golang"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/SlinSo/goTemplateBenchmark/templbench"
	"github.com/valyala/bytebufferpool"

	"github.com/SlinSo/goTemplateBenchmark/ego"
	"github.com/SlinSo/goTemplateBenchmark/ftmpl"
	goh "github.com/SlinSo/goTemplateBenchmark/goh"
	"github.com/SlinSo/goTemplateBenchmark/gorazor"
	herotmpl "github.com/SlinSo/goTemplateBenchmark/hero"
	"github.com/SlinSo/goTemplateBenchmark/jade"
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

	testComplexNav = []*model.Navigation{
		{
			Item: "Link 1",
			Link: "http://www.mytest.com/",
		}, {
			Item: "Link 2",
			Link: "http://www.mytest.com/",
		}, {
			Item: "Link 3",
			Link: "http://www.mytest.com/",
		},
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

func TestComplexGolangText(t *testing.T) {
	var buf bytes.Buffer

	funcMap := text.FuncMap{
		"safehtml": func(s string) string { return s },
	}

	templates := make(map[string]*text.Template)
	templatesDir := "go/"

	layouts, err := filepath.Glob(templatesDir + "layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	tempData := testComplexData.User.EscapedContent
	testComplexData.User.EscapedContent = text.HTMLEscapeString(testComplexData.User.EscapedContent)

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = text.Must(text.New("").Funcs(funcMap).ParseFiles(files...))
	}
	templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
	testComplexData.User.EscapedContent = tempData
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
		buf.Reset()
	}
}

func BenchmarkComplexGolangText(b *testing.B) {
	var buf bytes.Buffer

	funcMap := text.FuncMap{
		"safehtml": func(s string) string { return s },
	}

	templates := make(map[string]*text.Template)
	templatesDir := "go/"

	layouts, err := filepath.Glob(templatesDir + "layout/*.tmpl")
	if err != nil {
		panic(err)
	}

	includes, err := filepath.Glob(templatesDir + "includes/*.tmpl")
	if err != nil {
		panic(err)
	}

	tempData := testComplexData.User.EscapedContent
	testComplexData.User.EscapedContent = text.HTMLEscapeString(testComplexData.User.EscapedContent)

	// Generate our templates map from our layouts/ and includes/ directories
	for _, layout := range layouts {
		files := append(includes, layout)
		templates[filepath.Base(layout)] = text.Must(text.New("").Funcs(funcMap).ParseFiles(files...))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		templates["index.tmpl"].ExecuteTemplate(&buf, "base", testComplexData)
		buf.Reset()
	}
	testComplexData.User.EscapedContent = tempData
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
		buf.Reset()
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
		buf.Reset()
	}
}

/******************************************************************************
** templ
******************************************************************************/
func TestComplexTempl(t *testing.T) {
	var buf bytes.Buffer
	templbench.Index(testComplexUser, testComplexNav, testComplexTitle).Render(context.Background(), &buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexTempl(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		templbench.Index(testComplexUser, testComplexNav, testComplexTitle).Render(context.Background(), &buf)
		buf.Reset()
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

/******************************************************************************
** Jet
******************************************************************************/

func TestComplexJetHTML(t *testing.T) {
	var buf bytes.Buffer

	tmpl, err := jetSet.GetTemplate("index.jet")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.Execute(&buf, nil, testComplexData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexJetHTML(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := jetSet.GetTemplate("index.jet")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, nil, testComplexData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Goh
******************************************************************************/
func TestComplexGoh(t *testing.T) {
	var buf bytes.Buffer

	goh.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexGoh(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		goh.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** Hero
******************************************************************************/
func TestComplexHero(t *testing.T) {
	var buf bytes.Buffer

	herotmpl.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexHero(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		herotmpl.Index(testComplexUser, testComplexNav, testComplexTitle, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** Jade
******************************************************************************/
func TestComplexJade(t *testing.T) {
	buf := bytebufferpool.Get()

	jade.Index(testComplexUser, testComplexNav, testComplexTitle, buf)

	if msg, ok := linesEquals(buf.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkComplexJade(b *testing.B) {
	buf := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		jade.Index(testComplexUser, testComplexNav, testComplexTitle, buf)
		buf.Reset()
	}
}

/******************************************************************************
** Go func
******************************************************************************/
func TestComplexGoFunc(t *testing.T) {
	bb := bytebufferpool.Get()

	golang.Index(bb, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(bb.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
	bb.Reset()

	golang.Index2(bb, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(bb.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
	bb.Reset()

	golang.Index3(bb, testComplexUser, testComplexNav, testComplexTitle)

	if msg, ok := linesEquals(bb.String(), expectedtComplexResult); !ok {
		t.Error(msg)
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexGoDirectBuffer(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.Index(bb, testComplexUser, testComplexNav, testComplexTitle)
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexGoHyperscript(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.Index2(bb, testComplexUser, testComplexNav, testComplexTitle)
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexGoStaticString(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.Index3(bb, testComplexUser, testComplexNav, testComplexTitle)
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscapeHTML(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.EscapeHTML(testComplexData.User.EscapedContent, bb)
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscape(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.Escape(bb, golang.UnsafeStrToBytes(testComplexData.User.EscapedContent))
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscapeGo(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		bb.WriteString(html.EscapeString(testComplexData.User.EscapedContent))
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscapeHTMLNoop(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.EscapeHTML(testComplexData.User.FirstName, bb)
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscapeNoop(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.Escape(bb, golang.UnsafeStrToBytes(testComplexData.User.FirstName))
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}

func BenchmarkComplexEscapeGoNoop(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		bb.WriteString(html.EscapeString(testComplexData.User.FirstName))
		bb.Reset()
	}
	bytebufferpool.Put(bb)
}
