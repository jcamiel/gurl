package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWhitespaces(t *testing.T) {

	var parser *Parser
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
			parser = NewParserFromString(test.text, "")
			node, err = parser.parseWhitespaces()
			assert.Nil(t, err)
			assert.Equal(t, &Whitespaces{test.start, test.end, test.expectedText}, node, "Whitespaces should be parsed")
		})
	}
}
