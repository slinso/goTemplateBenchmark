package golang

import (
	"bytes"
	"html"
	"reflect"
	"strconv"
	"unsafe"

	"github.com/SlinSo/goTemplateBenchmark/model"
	"github.com/valyala/bytebufferpool"
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

func Escape(bb *bytebufferpool.ByteBuffer, b []byte) {
	write := bb.Write
	j := 0
	for i, c := range b {
		switch c {
		case '<':
			write(b[j:i])
			write(strLT)
			j = i + 1
		case '>':
			write(b[j:i])
			write(strGT)
			j = i + 1
		case '"':
			write(b[j:i])
			write(strQuot)
			j = i + 1
		case '\'':
			write(b[j:i])
			write(strApos)
			j = i + 1
		case '&':
			write(b[j:i])
			write(strAmp)
			j = i + 1
		}
	}
	write(b[j:])
}

var (
	strLT   = []byte("&lt;")
	strGT   = []byte("&gt;")
	strQuot = []byte("&quot;")
	strApos = []byte("&#39;")
	strAmp  = []byte("&amp;")
)

func hyper(bb *bytebufferpool.ByteBuffer, s string) {
	classFound := false
	idFound := false
	bb.WriteByte('<')

	b := UnsafeStrToBytes(s)

	for _, c := range b {
		switch c {
		case '#':
			bb.WriteString(" id=\"")
			idFound = true
		case '.':
			if idFound {
				bb.WriteByte('"')
			}
			if !classFound {
				bb.WriteString(" class=\"")
				classFound = true
			} else {
				bb.WriteByte(' ')
			}
		default:
			bb.WriteByte(c)
		}
	}
	if idFound || classFound {
		bb.WriteByte('"')
	}
	bb.WriteByte('>')
}

func Index3(bb *bytebufferpool.ByteBuffer, u *model.User, nav []*model.Navigation, title string) {
	bb.WriteString(`<!DOCTYPE html>
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

func Index2(bb *bytebufferpool.ByteBuffer, u *model.User, nav []*model.Navigation, title string) {
	bb.WriteString(`<!DOCTYPE html>`)
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
	bb.WriteString(`Hello `)
	bb.WriteString(html.EscapeString(u.FirstName))
	hyper(bb, "/h4")
	hyper(bb, "div.raw")
	bb.WriteString(u.RawContent)
	hyper(bb, "/div")
	hyper(bb, "div.enc")
	Escape(bb, UnsafeStrToBytes(u.EscapedContent))
	hyper(bb, "/div></div")

	for i := 1; i <= 5; i++ {
		hyper(bb, "p")
		bb.WriteString(html.EscapeString(u.FirstName))
		bb.WriteString(` has `)
		bb.B = strconv.AppendInt(bb.B, int64(i), 10)
		if i == 1 {
			bb.WriteString(` message</p>`)
		} else {
			bb.WriteString(` messages</p>`)
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
func EscapeHTML(html string, buffer *bytebufferpool.ByteBuffer) {
	var i, j, k int

	for i < len(html) {
		for j = i; j < len(html); j++ {
			k = bytes.IndexByte(escapedKeys, html[j])
			if k != -1 {
				break
			}
		}

		buffer.WriteString(html[i:j])
		if k != -1 {
			buffer.WriteString(escapedValues[k])
		}
		i = j + 1
	}
}

func Index(bb *bytebufferpool.ByteBuffer, u *model.User, nav []*model.Navigation, title string) {
	bb.WriteString(`
<!DOCTYPE html>
<html>
<body>

<header>
`)
	Header(bb, &title)
	bb.WriteString(`
</header>

<nav>
`)
	Navigation(bb, nav)
	bb.WriteString(`
</nav>

<section>
<div class="content">
	<div class="welcome">
		<h4>Hello `)
	Escape(bb, UnsafeStrToBytes(u.FirstName))
	bb.WriteString(`</h4>
		
		<div class="raw">`)
	bb.WriteString(u.RawContent)
	bb.WriteString(`</div>
		<div class="enc">`)
	Escape(bb, UnsafeStrToBytes(u.EscapedContent))
	bb.WriteString(`</div>
	</div>

`)
	for i := 1; i <= 5; i++ {
		bb.WriteString(`
			<p>`)
		bb.WriteString(html.EscapeString(u.FirstName))
		bb.WriteString(` has `)
		if i == 1 {
			bb.WriteString(`1 message</p>`)
		} else {
			bb.B = strconv.AppendInt(bb.B, int64(i), 10)
			bb.WriteString(` messages</p>`)
		}

	}
	bb.WriteString(`
</div>
</section>

<footer>
`)
	WriteFooter(bb)
	bb.WriteString(`
</footer>

</body>
</html>
`)
}
