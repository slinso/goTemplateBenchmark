package golang

import (
	"bytes"
	"html"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/SlinSo/goTemplateBenchmark/model"
)

var ()

func UnsafeStrToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func Escape(bb *bytes.Buffer, b []byte) {
	write := bb.Write
	j := 0
	for i, c := range b {
		switch c {
		case '<':
			_, _ = write(b[j:i])
			_, _ = write(strLT)
			j = i + 1
		case '>':
			_, _ = write(b[j:i])
			_, _ = write(strGT)
			j = i + 1
		case '"':
			_, _ = write(b[j:i])
			_, _ = write(strQuot)
			j = i + 1
		case '\'':
			_, _ = write(b[j:i])
			_, _ = write(strApos)
			j = i + 1
		case '&':
			_, _ = write(b[j:i])
			_, _ = write(strAmp)
			j = i + 1
		}
	}
	_, _ = write(b[j:])
}

var (
	strLT   = []byte("&lt;")
	strGT   = []byte("&gt;")
	strQuot = []byte("&quot;")
	strApos = []byte("&#39;")
	strAmp  = []byte("&amp;")
)

func hyper(bb *bytes.Buffer, s string) {
	classFound := false
	idFound := false
	_ = bb.WriteByte('<')

	b := UnsafeStrToBytes(s)

	for _, c := range b {
		switch c {
		case '#':
			_, _ = bb.WriteString(" id=\"")
			idFound = true
		case '.':
			if idFound {
				_ = bb.WriteByte('"')
			}
			if !classFound {
				_, _ = bb.WriteString(" class=\"")
				classFound = true
			} else {
				_ = bb.WriteByte(' ')
			}
		default:
			_ = bb.WriteByte(c)
		}
	}
	if idFound || classFound {
		_ = bb.WriteByte('"')
	}
	_ = bb.WriteByte('>')
}

func Index3(bb *bytes.Buffer, u *model.User, nav []*model.Navigation, title string) {
	_, _ = bb.WriteString(`<!DOCTYPE html>
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
</html>`)
}

func Index2(bb *bytes.Buffer, u *model.User, nav []*model.Navigation, title string) {
	_, _ = bb.WriteString(`<!DOCTYPE html>`)
	hyper(bb, "html")
	hyper(bb, "body")
	hyper(bb, "header")
	Header(bb, &title)
	hyper(bb, "/header")
	hyper(bb, "nav")
	Navigation(bb, nav)
	hyper(bb, "/nav")
	hyper(bb, "section")
	hyper(bb, "div.content")
	hyper(bb, "div.welcome")
	hyper(bb, "h4")
	_, _ = bb.WriteString(`Hello `)
	_, _ = bb.WriteString(html.EscapeString(u.FirstName))
	hyper(bb, "/h4")
	hyper(bb, "div.raw")
	_, _ = bb.WriteString(u.RawContent)
	hyper(bb, "/div")
	hyper(bb, "div.enc")
	Escape(bb, UnsafeStrToBytes(u.EscapedContent))
	hyper(bb, "/div></div")

	for i := 1; i <= 5; i++ {
		hyper(bb, "p")
		_, _ = bb.WriteString(html.EscapeString(u.FirstName))
		_, _ = bb.WriteString(` has `)
		bb.WriteString(strconv.FormatInt(int64(i), 10))
		if i == 1 {
			_, _ = bb.WriteString(` message</p>`)
		} else {
			_, _ = bb.WriteString(` messages</p>`)
		}

	}
	hyper(bb, "/div></section")
	hyper(bb, "footer")
	WriteFooter(bb)
	hyper(bb, "/footer></body></html")
}

var (
	escapedKeys   = []byte{'&', '\'', '<', '>', '"'}
	escapedValues = []string{"&amp;", "&#39;", "&lt;", "&gt;", "&#34;"}
)

// EscapeHTML escapes the html and then put it to the buffer.
func EscapeHTML(html string, buffer *bytes.Buffer) {
	var i, j, k int

	for i < len(html) {
		for j = i; j < len(html); j++ {
			k = bytes.IndexByte(escapedKeys, html[j])
			if k != -1 {
				break
			}
		}

		_, _ = buffer.WriteString(html[i:j])
		if k != -1 {
			_, _ = buffer.WriteString(escapedValues[k])
		}
		i = j + 1
	}
}

func Index(bb *bytes.Buffer, u *model.User, nav []*model.Navigation, title string) {
	_, _ = bb.WriteString(`
<!DOCTYPE html>
<html>
<body>

<header>
`)
	Header(bb, &title)
	_, _ = bb.WriteString(`
</header>

<nav>
`)
	Navigation(bb, nav)
	_, _ = bb.WriteString(`
</nav>

<section>
<div class="content">
	<div class="welcome">
		<h4>Hello `)
	Escape(bb, UnsafeStrToBytes(u.FirstName))
	_, _ = bb.WriteString(`</h4>
		
		<div class="raw">`)
	_, _ = bb.WriteString(u.RawContent)
	_, _ = bb.WriteString(`</div>
		<div class="enc">`)
	Escape(bb, UnsafeStrToBytes(u.EscapedContent))
	_, _ = bb.WriteString(`</div>
	</div>

`)
	for i := 1; i <= 5; i++ {
		_, _ = bb.WriteString(`
			<p>`)
		_, _ = bb.WriteString(html.EscapeString(u.FirstName))
		_, _ = bb.WriteString(` has `)
		if i == 1 {
			_, _ = bb.WriteString(`1 message</p>`)
		} else {
			bb.WriteString(strconv.FormatInt(int64(i), 10))
			_, _ = bb.WriteString(` messages</p>`)
		}

	}
	_, _ = bb.WriteString(`
</div>
</section>

<footer>
`)
	WriteFooter(bb)
	_, _ = bb.WriteString(`
</footer>

</body>
</html>
`)
}
