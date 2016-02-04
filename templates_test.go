package main_test

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"

	"github.com/SlinSo/goTemplateBenchmark/model"

	"github.com/SlinSo/goTemplateBenchmark/ego"
	"github.com/SlinSo/goTemplateBenchmark/egon"
	"github.com/SlinSo/goTemplateBenchmark/egonslinso"
	"github.com/SlinSo/goTemplateBenchmark/ftmpl"
	"github.com/SlinSo/goTemplateBenchmark/gorazor"
	"github.com/aymerick/raymond"
	"github.com/dskinner/damsel"
	"github.com/eknkc/amber"
	"github.com/flosch/pongo2"
	"github.com/hoisie/mustache"
	"github.com/robfig/soy"
	"github.com/robfig/soy/data"
	"github.com/yosssi/ace"
	"github.com/ziutek/kasia.go"

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
	tmpl.Execute(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGolang(b *testing.B) {
	var buf bytes.Buffer

	tmpl, _ := template.ParseFiles("go/simple.tmpl")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmpl.Execute(&buf, testData)
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
	}
}

/******************************************************************************
** Egon
******************************************************************************/
func TestEgon(t *testing.T) {
	var buf bytes.Buffer
	egon.SimpleTemplate(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkEgon(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egon.SimpleTemplate(&buf, testData)
	}
}
func BenchmarkEgonFooter(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egon.FooterTemplate(&buf)
	}
}

/******************************************************************************
** EgonSlinSo
******************************************************************************/
func TestEgonSlinso(t *testing.T) {
	var buf bytes.Buffer
	egonslinso.SimpleTemplate(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkEgonSlinso(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egonslinso.SimpleTemplate(&buf, testData)
	}
}
func BenchmarkEgonSlinsoFooter(b *testing.B) {
	var buf bytes.Buffer

	for i := 0; i < b.N; i++ {
		egonslinso.FooterTemplate(&buf)
	}
}

/******************************************************************************
** ftmpl
******************************************************************************/
func TestFtmpl(t *testing.T) {
	result := ftmpl.T__simple(testData)

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkFtmpl(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = ftmpl.T__simple(testData)
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

	tpl, _ := ace.Load("ace/simple", "", nil)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tpl.Execute(&buf, testData)
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
		tpl.Execute(&buf, testData)
	}
}

/******************************************************************************
** Damsel
******************************************************************************/
func TestDamsel(t *testing.T) {
	dmsl, err := damsel.ParseFile("damsel/simple.dmsl")
	if err != nil {
		t.Error(err)
	}

	tpl := damsel.NewHtmlTemplate(dmsl)
	result, err := tpl.Execute(testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkDamsel(b *testing.B) {
	dmsl, _ := damsel.ParseFile("damsel/simple.dmsl")

	tpl := damsel.NewHtmlTemplate(dmsl)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tpl.Execute(testData)
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
		tpl.ExecuteWriterUnbuffered(pongo2.Context{"u": testData}, &buf)
	}
}

/******************************************************************************
** Handlebars
******************************************************************************/
func TestHandlebars(t *testing.T) {
	tpl, err := raymond.ParseFile("handlebars/simple.handle")
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
	tpl, _ := raymond.ParseFile("handlebars/simple.handle")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tpl.Exec(testData)
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
** Kasia
******************************************************************************/
func TestKasia(t *testing.T) {
	var buf bytes.Buffer

	tpl, err := kasia.ParseFile("kasia/simple.kt")
	if err != nil {
		t.Error(err)
	}

	err = tpl.Execute(&buf, testData)
	if err != nil {
		t.Error(err)
	}

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkKasia(b *testing.B) {
	var buf bytes.Buffer

	tpl, _ := kasia.ParseFile("kasia/simple.kt")
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tpl.Execute(&buf, testData)
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
		tofu.Render(&buf, "soy.simple", soyData)
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

	if bytes.Compare(b1, b2) != 0 {
		return fmt.Sprintf("Lines don't match \n1:\"%s\"\n2:\"%s\"", b1, b2), false
	}

	return "", true
}
