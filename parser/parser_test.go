package parser

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
	assert.Equal(t, 'S', r, "First rune should be equal")
	assert.Equal(t, 1, parser.Current, "Parser index should be incremented")
	assert.Equal(t, 1, parser.Line, "Parser line should be equal")
}

func TestLineCount(t *testing.T) {
	text := `Some multiline string:
Line 2
Line 3
Line 4`
	parser := NewParserFromString(text, "")

	for {
		_, err := parser.readRune()
		if err != nil {
			break
		}
	}
	assert.Equal(t, 4, parser.Line, "Parser line should be equal")
}

func TestEndOfFile(t *testing.T) {
	text := `123`
	parser := NewParserFromString(text, "")

	for range text {
		_, _ = parser.readRune()
	}

	_, err := parser.readRune()
	assert.Equal(t, io.EOF, err, "Error should be end of file")
}

func TestReadEmptyString(t *testing.T) {
	text := ""
	parser := NewParserFromString(text, "")
	_, err := parser.readRune()
	assert.Equal(t, io.EOF, err, "Error should be end of file")
}
