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

func TestEmptyString(t *testing.T) {
	text := ""
	parser := NewParserFromString(text, "")
	_, err := parser.readRune()
	assert.Equal(t, io.EOF, err, "Error should be end of file")
}

func TestParseWhiteSpace(t *testing.T) {

	var parser *Parser
	var node Node
	var err error

	var tests = []struct {
		text         string
		expectedText string
		skipNewLine  bool
		start        Position
		end          Position
	}{
		{"\u0020\u0020\n\n\u0020ABCD", "\u0020\u0020", false, Position{0, 1}, Position{2, 1}},
		{"\u0020\u0020\n\n\u0020ABCD", "\u0020\u0020\n\n\u0020", true, Position{0, 1}, Position{5, 3}},
		{"\n", "\n", true, Position{0, 1}, Position{1, 2}},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			parser = NewParserFromString(test.text, "")
			node, err = parser.parseWhiteSpace(test.skipNewLine)
			assert.Nil(t, err)
			assert.Equal(t, &Whitespace{test.start, test.end, test.expectedText}, node, "Whitespace should be parsed")
		})
	}
}
