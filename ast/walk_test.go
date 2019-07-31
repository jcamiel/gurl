package ast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DummyVisitor struct {
	text string
}

func (v *DummyVisitor) Visit(node Noder) Visitor {
	if node != nil {
		v.text += fmt.Sprintf("%T begin=%+v end=%+v\n",
			node, node.GetBegin(), node.GetEnd())
	}

	return v
}

func TestVisit(t *testing.T) {

	text := `# Some hurl file
GET http://sample.org
User-Agent : Something
[Cookies]
a: 12
[QueryStringParams]
q: test
[FormParams]
fruit: "apple"

HTTP/1.1 200
[Captures]
csrf: xpath //div/text()
[Asserts]
xpath count(//body) equals 1
xpath boolean(count(//tag)) equals false
`
	p := NewParserFromString(text, "")
	hurl := p.Parse()

	v := DummyVisitor{}
	Walk(&v, hurl)

	assert.True(t, len(v.text) > 0)
}
