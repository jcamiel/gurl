package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMethodSucceed(t *testing.T) {
	text := "POST http://google.com"
	parser := NewParserFromString(text, "")
	method, err := parser.parseMethod()

	assert.Nil(t, err)
	assert.Equal(t, &Method{
		Position{0, 1},
		Position{4, 1},
		"POST",
	}, method, "POST should be parsed")
}

func TestParseMethodFailed(t *testing.T) {
	text := "ABCEDFGHIJKLM"
	parser := NewParserFromString(text, "")
	method, err := parser.parseMethod()
	assert.Nil(t, method)
	assert.NotNil(t, err)
}
