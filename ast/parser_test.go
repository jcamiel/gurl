package ast

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestReadRune(t *testing.T) {
	text := `Some content`
	p := NewParserFromString(text, "")

	r, err := p.readRune()
	assert.Nil(t, err)
	assert.Equal(t, 'S', r, "First rune should be equal")
	assert.Equal(t, 1, p.pos.Offset, "Parser offset should be incremented")
	assert.Equal(t, 1, p.pos.Line, "Parser line should be equal")
}

func TestLineCount(t *testing.T) {
	text := `Some multiline string:
line 2
line 3
line 4`
	p := NewParserFromString(text, "")

	for {
		_, err := p.readRune()
		if err != nil {
			break
		}
	}
	assert.Equal(t, 4, p.pos.Line, "Parser line should be equal")
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
