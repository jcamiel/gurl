package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMethodSucceed(t *testing.T) {
	text := "POST http://google.com"
	p := NewParserFromString(text, "")
	node, err := p.parseMethod()

	assert.Nil(t, err)
	assert.Equal(t, &Method{
		Position{0, 1},
		Position{4, 1},
		"POST",
	}, node, "POST should be parsed")
}

func TestParseMethodFailed(t *testing.T) {
	text := "ABCEDFGHIJKLM"
	p := NewParserFromString(text, "")
	node, err := p.parseMethod()
	assert.Nil(t, node)
	assert.NotNil(t, err)
}

func TestParseRequest(t *testing.T) {

	var text string
	var node *Request
	var p *Parser

	text = "GET	http://www.example.org"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET\u0020http://www.example.org"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET	http://www.example.org	# Some comment"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET http://www.example.org/foo.html#bar # Some comment"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

}


func TestParseFailed(t *testing.T) {
	text := "\n\nPOSThttp://google.com"
	p := NewParserFromString(text, "")
	_, err := p.parseRequest()
	fmt.Println(err)
	assert.NotNil(t, err)
}