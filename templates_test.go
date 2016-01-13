package main_test

import (
	"bytes"
	"fmt"
	//	"log"
	"html/template"
	"strings"
	"testing"

	"github.com/SlinSo/goTemplateBenchmark/model"

	"github.com/SlinSo/goTemplateBenchmark/ego"
	"github.com/SlinSo/goTemplateBenchmark/ftmpl"
)

var (
	testData = &model.User{
		FirstName:      "Bob",
		FavoriteColors: []string{"blue", "green", "mauve"},
	}
	golangTemplate = template.Must(template.New("").Parse(`
<html>
    <body>
        <h1>{{ .FirstName }}</h1>

        <p>Here's a list of your favorite colors:</p>
        <ul>
        {{ range .FavoriteColors }}
            <li>{{ . }}</li>{{ end }}
        </ul>
    </body>
</html>
`))

	expectedtResult = `
<html>
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
)

func TestGolang(t *testing.T) {
	var buf bytes.Buffer
	golangTemplate.Execute(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func TestEgo(t *testing.T) {
	var buf bytes.Buffer
	ego.EgoSimple(&buf, testData)

	if msg, ok := linesEquals(buf.String(), expectedtResult); !ok {
		t.Error(msg)
	}
}

func TestFtmpl(t *testing.T) {
	result := ftmpl.T__simple(testData)

	if msg, ok := linesEquals(result, expectedtResult); !ok {
		t.Error(msg)
	}
}

func BenchmarkGolang(b *testing.B) {
    var buf bytes.Buffer
    
    for i := 0; i < b.N; i++ {
        golangTemplate.Execute(&buf, testData)
    }
}

func BenchmarkEgo(b *testing.B) {
    var buf bytes.Buffer
    
    for i := 0; i < b.N; i++ {
        ego.EgoSimple(&buf, testData)
    }
}

func BenchmarkFtmpl(b *testing.B) {
    for i := 0; i < b.N; i++ {
       _ =ftmpl.T__simple(testData)
    }
}

func linesEquals(str1, str2 string) (explanation string, equals bool) {
	if str1 == str2 {
		return "", true
	}

	lines1 := strings.Split(strings.TrimSpace(str1), "\n")
	lines2 := strings.Split(strings.TrimSpace(str2), "\n")

	if len(lines1) != len(lines2) {
		return fmt.Sprintf("Lines count don't match %d!=%d", len(lines1), len(lines2)), false
	}

	for i := 0; i < len(lines1); i++ {
		line1 := lines1[i]
		line2 := lines2[i]
		if line1 != line2 {
			return fmt.Sprintf("Line #%d don't match \"%s\"!=\"%s\"", i, line1, line2), false
		}
	}

	return "", true
}
