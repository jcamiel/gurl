package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWhitespaces(t *testing.T) {

	var p *Parser
	var node Node
	var err error

	var tests = []struct {
		text         string
		expectedText string
		start        Position
		end          Position
	}{
		{
			"\u0020\t\u0020\t\n\t\n\u0020ABCD",
			"\u0020\t\u0020\t\n\t\n\u0020",
			Position{0, 1},
			Position{8, 3},
		},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node, err = p.parseWhitespaces()
			assert.Nil(t, err)
			assert.Equal(t, &Whitespaces{test.start, test.end, test.expectedText}, node, "Whitespaces should be parsed")
		})
	}
}

func TestParseOptionalWhitespaces(t *testing.T) {

	text := "\u0020\u0020\nABCDEF"
	p := NewParserFromString(text, "")

	node, err := p.tryParse(p.parseWhitespaces)
	assert.Nil(t, err)
	assert.Equal(t, &Whitespaces{Position{0, 1}, Position{2, 1}, "\u0020\u0020\n"}, node, "Whitespaces should be parsed")
	assert.Equal(t, 3, p.Current, "Offset should be equal")
	assert.Equal(t, 2, p.Line, "Line should be equal")

	node, err = p.tryParse(p.parseWhitespaces)
	assert.NotNil(t, err)
	assert.Nil(t, node, "Whitespace should not be parsed")
	assert.Equal(t, 3, p.Current, "Offset should be equal")
	assert.Equal(t, 2, p.Line, "Offset should be equal")
}

func TestParseWhitespacesFailed(t *testing.T) {
	text := ""
	p := NewParserFromString(text, "")
	node, err := p.parseWhitespaces()
	assert.Nil(t, node)
	assert.IsType(t, &SyntaxError{}, err)
}