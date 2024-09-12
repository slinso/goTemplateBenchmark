package main_test

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"testing"
	text "text/template"

	"github.com/SlinSo/goTemplateBenchmark/golang"
	"github.com/SlinSo/goTemplateBenchmark/gomponents"
	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/SlinSo/goTemplateBenchmark/templbench"
	"github.com/valyala/bytebufferpool"

	"github.com/SlinSo/goTemplateBenchmark/ego"
	"github.com/SlinSo/goTemplateBenchmark/ftmpl"
	"github.com/SlinSo/goTemplateBenchmark/gorazor"
	herotmpl "github.com/SlinSo/goTemplateBenchmark/hero"
	"github.com/SlinSo/goTemplateBenchmark/jade"
	"github.com/SlinSo/goTemplateBenchmark/quicktemplate"
	"github.com/aymerick/raymond"
	"github.com/eknkc/amber"
	"github.com/flosch/pongo2"
	"github.com/hoisie/mustache"
	"github.com/robfig/soy"
	"github.com/robfig/soy/data"
	"github.com/yosssi/ace"

	"github.com/CloudyKit/jet"
	"github.com/dchest/htmlmin"
)

var (
	testData = &model.User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}

	expectedtResult = `<html>
	<body>
		<h1>Bob</h1>

		<p>Here's a list of your favorite colors:</p>
		<ul>

			<li>blue</li>
			<li>green</li>
			<li>mauve</li>
		</ul>
	</body>
</html>`

	expectedtResultMinified = "<html><body><h1>Bob</h1><p>Here's a list of your favorite colors:</p><ul><li>blue</li><li>green</li><li>mauve</li></ul></body></html>"
)

/******************************************************************************
** Go
******************************************************************************/
func TestGolang(t *testing.T) {
	var buf bytes.Buffer

	tmpl, err := template.ParseFiles("go/simple.tmpl")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.Execute(&buf, testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func TestGolangText(t *testing.T) {
	var buf bytes.Buffer

	tmpl, err := text.ParseFiles("go/simple.tmpl")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.Execute(&buf, testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGolang(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := template.ParseFiles("go/simple.tmpl")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func BenchmarkGolangText(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := text.ParseFiles("go/simple.tmpl")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** GoFunctions
******************************************************************************/
func TestGoFunc(t *testing.T) {
	bb := bytebufferpool.Get()
	golang.WriteSimpleGolang(bb, testData)

	if msg, ok := linesEquals(bb.String(), expectedtResult); !ok {
		t.Error(msg)
	}
	bb.Reset()
	g := golang.NewGoFunc(bb)

	golang.GoFuncElem(g, testData)
	if msg, ok := linesEquals(bb.String(), expectedtResult); !ok {
		t.Error(msg)
	}
	bb.Reset()
	g = golang.NewGoFunc(bb)

	golang.GoFuncFunc(g, testData)
	if msg, ok := linesEquals(bb.String(), expectedtResult); !ok {
		t.Error(msg)
	}

	bytebufferpool.Put(bb)
}

func BenchmarkGoDirectBuffer(b *testing.B) {
	bb := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		golang.WriteSimpleGolang(bb, testData)
		bb.Reset()
	}
}

func BenchmarkGoCustomHtmlAPI(b *testing.B) {
	buf := bytebufferpool.Get()
	g := golang.NewGoFunc(buf)

	for i := 0; i < b.N; i++ {
		golang.GoFuncElem(g, testData)
		buf.Reset()
	}
}

func BenchmarkGoFunc3(b *testing.B) {
	buf := bytebufferpool.Get()
	g := golang.NewGoFunc(buf)

	for i := 0; i < b.N; i++ {
		golang.GoFuncFunc(g, testData)
		buf.Reset()
	}
}

/******************************************************************************
** Ego
******************************************************************************/
func TestEgo(t *testing.T) {
	var buf bytes.Buffer
	ego.EgoSimple(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkEgo(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		ego.EgoSimple(&buf, testData)
		buf.Reset()
	}
}

/******************************************************************************
** Quicktemplate
******************************************************************************/
func TestQuicktemplate(t *testing.T) {
	var buf bytes.Buffer
	quicktemplate.WriteSimpleQtc(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkQuicktemplate(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		quicktemplate.WriteSimpleQtc(&buf, testData)
		buf.Reset()
	}
}

/******************************************************************************
** ftmpl
******************************************************************************/
func TestFtmpl(t *testing.T) {
	result, err := ftmpl.TMPLERRsimple(testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkFtmpl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := ftmpl.TMPLERRsimple(testData)
		if err != nil {
			b.Fatal(err)
		}

	}
}

/******************************************************************************
** Ace
******************************************************************************/
func TestAce(t *testing.T) {
	var buf bytes.Buffer

	tpl, err := ace.Load("ace/simple", "", nil)
	if err != nil {
		t.Error(err)
	}

	if err := tpl.Execute(&buf, testData); err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResultMinified); !ok {
		t.Error(msg)
	}
}

func BenchmarkAce(b *testing.B) {
	var buf bytes.Buffer

	tpl, err := ace.Load("ace/simple", "", nil)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tpl.Execute(&buf, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Amber
******************************************************************************/
func TestAmber(t *testing.T) {
	var buf bytes.Buffer

	tpl, err := amber.CompileFile("amber/simple.amber", amber.DefaultOptions)
	if err != nil {
		t.Error(err)
	}

	if err := tpl.Execute(&buf, testData); err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkAmber(b *testing.B) {
	var buf bytes.Buffer

	tpl, _ := amber.CompileFile("amber/simple.amber", amber.DefaultOptions)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tpl.Execute(&buf, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Mustache
******************************************************************************/
func TestMustache(t *testing.T) {
	tpl, err := mustache.ParseFile("mustache/simple.mustache")
	if err != nil {
		t.Error(err)
	}

	result := tpl.Render(testData)

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkMustache(b *testing.B) {
	tpl, _ := mustache.ParseFile("mustache/simple.mustache")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tpl.Render(testData)
	}
}

/******************************************************************************
** pongo2
******************************************************************************/
func TestPongo2(t *testing.T) {
	var buf bytes.Buffer

	tpl, err := pongo2.FromFile("pongo2/simple.pongo")
	if err != nil {
		t.Error(err)
	}

	err = tpl.ExecuteWriter(pongo2.Context{"u": testData}, &buf)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkPongo2(b *testing.B) {
	var buf bytes.Buffer

	tpl, _ := pongo2.FromFile("pongo2/simple.pongo")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tpl.ExecuteWriter(pongo2.Context{"u": testData}, &buf)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Handlebars
******************************************************************************/
func TestHandlebars(t *testing.T) {
	tpl, err := raymond.ParseFile("raymond/simple.handle")
	if err != nil {
		t.Error(err)
	}

	result, err := tpl.Exec(testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkHandlebars(b *testing.B) {
	tpl, _ := raymond.ParseFile("raymond/simple.handle")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := tpl.Exec(testData)
		if err != nil {
			b.Fatal(err)
		}
	}
}

/******************************************************************************
** gorazor
******************************************************************************/
func TestGorazor(t *testing.T) {
	result := gorazor.Simple(testData)

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGorazor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gorazor.Simple(testData)
	}
}

/******************************************************************************
** Soy
******************************************************************************/
func TestSoy(t *testing.T) {
	var buf bytes.Buffer

	tofu, err := soy.NewBundle().AddTemplateDir("soy").CompileToTofu()
	if err != nil {
		t.Error(err)
	}

	err = tofu.Render(&buf, "soy.simple", map[string]interface{}{
		"user": data.New(testData),
	})
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkSoy(b *testing.B) {
	var buf bytes.Buffer

	tofu, _ := soy.NewBundle().AddTemplateDir("soy").CompileToTofu()
	soyData := map[string]interface{}{
		"user": data.New(testData),
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tofu.Render(&buf, "soy.simple", soyData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Jet
******************************************************************************/
var jetSet = jet.NewHTMLSet("./jet")

func TestJetHTML(t *testing.T) {
	var buf bytes.Buffer

	tmpl, err := jetSet.GetTemplate("simple.jet")
	if err != nil {
		t.Error(err)
	}
	err = tmpl.Execute(&buf, nil, testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkJetHTML(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := jetSet.GetTemplate("simple.jet")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := tmpl.Execute(&buf, nil, testData)
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

/******************************************************************************
** Hero
******************************************************************************/
func TestHero(t *testing.T) {
	var buf bytes.Buffer

	herotmpl.SimpleQtc(testData, &buf)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkHero(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		herotmpl.SimpleQtc(testData, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** Jade
******************************************************************************/
func TestJade(t *testing.T) {
	buf := bytebufferpool.Get()

	jade.Simple(testData, buf)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkJade(b *testing.B) {
	buf := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		jade.Simple(testData, buf)
		buf.Reset()
	}
}

/******************************************************************************
** templ
******************************************************************************/
func TestTempl(t *testing.T) {
	var buf bytes.Buffer
	templbench.SimpleTempl(testData).Render(context.Background(), &buf)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkTempl(b *testing.B) {
	var buf bytes.Buffer

	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		templbench.SimpleTempl(testData).Render(ctx, &buf)
		buf.Reset()
	}
}

/******************************************************************************
** Gomponents
******************************************************************************/
func TestGomponents(t *testing.T) {
	buf := bytebufferpool.Get()

	err := gomponents.Page(testData).Render(buf)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGomponents(b *testing.B) {
	buf := bytebufferpool.Get()

	for i := 0; i < b.N; i++ {
		gomponents.Page(testData).Render(buf)
		buf.Reset()
	}
}

/******************************************************************************
** helpers
******************************************************************************/
func linesEquals(str1, str2 string) (explanation string, equals bool) {
	if str1 == str2 {
		return "", true
	}

	// Minify removes whitespace infront of the first tag
	b1, err := htmlmin.Minify([]byte(str1), nil)
	if err != nil {
		panic(err)
	}

	b2, err := htmlmin.Minify([]byte(str2), nil)
	if err != nil {
		panic(err)
	}

	b1 = bytes.Replace(b1, []byte(" "), []byte("[space]"), -1)
	b1 = bytes.Replace(b1, []byte("\t"), []byte("[tab]"), -1)
	b1 = bytes.Replace(b1, []byte("\n"), []byte(""), -1)

	b2 = bytes.Replace(b2, []byte(" "), []byte("[space]"), -1)
	b2 = bytes.Replace(b2, []byte("\t"), []byte("[tab]"), -1)
	b2 = bytes.Replace(b2, []byte("\n"), []byte(""), -1)

	if !bytes.Equal(b1, b2) {
		return fmt.Sprintf("Lines don't match \n1:\"%s\"\n2:\"%s\"", b1, b2), false
	}

	return "", true
}
