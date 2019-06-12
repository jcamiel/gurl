package gurl

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestReadRune(t *testing.T) {
	text := `Some content`
	parser := NewParserFromString(text, "")

	r, err := parser.readRune()
	assert.Nil(t, err)
	assert.Equal(t, r, 'S', "First rune should be equal")
}

func TestEndOfFile(t *testing.T) {
	text := `123`
	parser := NewParserFromString(text, "")

	for range text {
		_, _ = parser.readRune()
	}

	_, err := parser.readRune()
	assert.Equal(t, err, io.EOF, "Error should be end of file")
}

func TestEmptyString(t *testing.T) {
	text := ""
	parser := NewParserFromString(text, "")
	_, err := parser.readRune()
	assert.Equal(t, err, io.EOF, "Error should be end of file")
}

func TestParseWhiteSpace(t *testing.T) {

	prefix := "start"
	suffix := "end"
	spaces := "\u0020\u0020\u0020"
	text := prefix + spaces + suffix

	parser := NewParserFromString(text, "")

	for i:=0; i < len(prefix); i++ {
		_, _ = parser.readRune()
	}

	token, err := parser.parseWhiteSpace(false)
	assert.Nil(t, err)
	assert.Equal(t, token, &WhitespaceToken{len(prefix), spaces}, "WhitespaceToken should be parsed")
}
