package ast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWhitespaces(t *testing.T) {

	var p *Parser
	var node interface{}
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

	node, err := p.tryParseWhitespaces()
	assert.Nil(t, err)
	assert.Equal(t, &Whitespaces{Position{0, 1}, Position{3, 2}, "\u0020\u0020\n"}, node, "Whitespaces should be parsed")
	assert.Equal(t, 3, p.current, "Offset should be equal")
	assert.Equal(t, 2, p.line, "line should be equal")

	node, err = p.tryParseWhitespaces()
	assert.NotNil(t, err)
	assert.Nil(t, node, "Whitespace should not be parsed")
	assert.Equal(t, 3, p.current, "Offset should be equal")
	assert.Equal(t, 2, p.line, "Offset should be equal")
}

func TestParseWhitespacesFailed(t *testing.T) {
	text := ""
	p := NewParserFromString(text, "")
	node, err := p.parseWhitespaces()
	assert.Nil(t, node)
	assert.IsType(t, &SyntaxError{}, err)
}

func TestParseComment(t *testing.T) {

	var text string
	var node *Comment
	var p *Parser

	text = "# Some comments\nBla bla bal"
	p = NewParserFromString(text, "")
	node, _ = p.parseComment()
	assert.NotNil(t, node)

	text = "abcedf"
	p = NewParserFromString(text, "")
	node, _ = p.parseComment()
	assert.Nil(t, node)
}

func TestParseCommentLine(t *testing.T) {

	var text string
	var node *CommentLine
	var p *Parser

	text = "# Some comments\n\n\n\t\t\tBla bla bal"
	p = NewParserFromString(text, "")
	node, _ = p.parseCommentLine()
	assert.NotNil(t, node)

	text = "abcedf"
	p = NewParserFromString(text, "")
	node, _ = p.parseCommentLine()
	assert.Nil(t, node)

}

func TestParseComments(t *testing.T) {

	var text string
	var node *Comments
	var p *Parser

	text = `# Some comments on line 1
# Some comments on line 2

# Some comments on line 3
# Some comments on line 4


Bla bla bal`

	p = NewParserFromString(text, "")
	node, _ = p.parseComments()
	assert.NotNil(t, node)
}

func TestParseJsonString(t *testing.T) {

	var node *JsonString
	var p *Parser
	var err error

	var tests = []struct {
		text          string
		expectedValue string
		error            bool
	}{
		{text:`"abcdef 012345"`, expectedValue:"abcdef 012345"},
		{text: `"abc\ndef"`, expectedValue: "abc\ndef"},
		{text: `"abc\"def"`, expectedValue: "abc\"def"},
		{text: `abc`, error: true},
		{text: `"abc`, error: true},
		{text: `"abc\ud83d\udca9"`, error: true},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {

			p = NewParserFromString(test.text, "")
			node, err = p.parseJsonString()

			if !test.error {
				assert.Equal(t, test.expectedValue, node.Value)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func TestParseKeyString(t *testing.T) {

	var node *KeyString
	var p *Parser
	var err error

	var tests = []struct {
		text          string
		expectedValue string
		error            bool
	}{
		{text:`abcedf`, expectedValue:"abcedf"},
		{text:`key:value`, expectedValue:"key"},
		{text:`fruit : banana"`, expectedValue:"fruit"},
		{text:`: kiwi"`, error:true},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {

			p = NewParserFromString(test.text, "")
			node, err = p.parseKeyString()

			if !test.error {
				assert.Equal(t, test.expectedValue, node.Text)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
