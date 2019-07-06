package ast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseWhitespaces(t *testing.T) {
	var p *Parser
	var node interface{}

	var tests = []struct {
		text         string
		expectedText string
		begin        Position
		end          Position
	}{
		{
			"\u0020\t\u0020\t\n\t\n\u0020ABCD",
			"\u0020\t\u0020\t\n\t\n\u0020",
			Position{0, 1, 1},
			Position{8, 3, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseWhitespaces()
			assert.Nil(t, p.Err())
			assert.Equal(t, &Whitespaces{Node{test.begin, test.end}, test.expectedText}, node, "Whitespaces should be parsed")
		})
	}
}

func TestParseOptionalWhitespaces(t *testing.T) {

	text := "\u0020\u0020\nABCDEF"
	p := NewParserFromString(text, "")

	node := p.tryParseWhitespaces()
	assert.Nil(t, p.Err())
	assert.Equal(t, &Whitespaces{Node{Position{0, 1, 1}, Position{3, 2, 1}}, "\u0020\u0020\n"}, node, "Whitespaces should be parsed")
	assert.Equal(t, 3, p.pos.Offset, "offset should be equal")
	assert.Equal(t, 2, p.pos.Line, "line should be equal")

	node = p.tryParseWhitespaces()
	assert.Nil(t, p.Err())
	assert.Nil(t, node, "Whitespace should not be parsed")
	assert.Equal(t, 3, p.pos.Offset, "offset should be equal")
	assert.Equal(t, 2, p.pos.Line, "line should be equal")
}

func TestParseWhitespacesFailed(t *testing.T) {
	text := ""
	p := NewParserFromString(text, "")
	node := p.parseWhitespaces()
	assert.NotNil(t, p.Err())
	assert.Nil(t, node)
}

func TestParseComment(t *testing.T) {
	var text string
	var node *Comment
	var p *Parser

	text = "# Some comments\nBla bla bal"
	p = NewParserFromString(text, "")
	node = p.parseComment()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())

	text = "abcedf"
	p = NewParserFromString(text, "")
	node = p.parseComment()
	assert.Nil(t, node)
	assert.NotNil(t, p.Err())
}

func TestParseCommentLine(t *testing.T) {
	var text string
	var node *CommentLine
	var p *Parser

	text = "# Some comments\n\n\n\t\t\tBla bla bal"
	p = NewParserFromString(text, "")
	node = p.parseCommentLine()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())

	text = "abcedf"
	p = NewParserFromString(text, "")
	node = p.parseCommentLine()
	assert.Nil(t, node)
	assert.NotNil(t, p.Err())

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
	node = p.parseComments()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestParseJsonString(t *testing.T) {
	var node *JsonString
	var p *Parser

	var tests = []struct {
		text          string
		expectedValue string
		error         bool
	}{
		{text: `"abcdef 012345"`, expectedValue: "abcdef 012345"},
		{text: `"abcd"toto`, expectedValue: "abcd"},
		{text: `"abc\ndef"`, expectedValue: "abc\ndef"},
		{text: `"abc\"def"`, expectedValue: "abc\"def"},
		{text: `"abc\ud83d\udca9"`, expectedValue: "abc💩"},
		{text: `abc`, error: true},
		{text: `"abc`, error: true},
		{text: `{"id":"123"}`, error: true},
		{text: `{"id": 0,"selected": true}`, error: true},
		{text: `true`, error: true},
		{text: "\"abcdef\n012345\"", error: true},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseJsonString()
			if !test.error {
				assert.NotNil(t, node)
				assert.Equal(t, test.expectedValue, node.Value)
				assert.Nil(t, p.Err())
			} else {
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestParseKeyString(t *testing.T) {
	var node *KeyString
	var p *Parser

	var tests = []struct {
		text          string
		expectedValue string
		error         bool
	}{
		{text: `abcedf`, expectedValue: "abcedf"},
		{text: `key:value`, expectedValue: "key"},
		{text: `fruit : banana"`, expectedValue: "fruit"},
		{text: `: kiwi"`, error: true},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseKeyString()
			if !test.error {
				assert.Equal(t, test.expectedValue, node.Value)
				assert.Nil(t, p.Err())
			} else {
				assert.Nil(t, node)
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestParseKey(t *testing.T) {
	var node *Key
	var p *Parser

	var tests = []struct {
		text          string
		expectedValue string
		error         bool
	}{
		{text: `key:value`, expectedValue: "key"},
		{text: `"key":012345678`, expectedValue: "key"},
		{text: `"key1:key2":012345678`, expectedValue: "key1:key2"},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseKey()
			if !test.error {
				assert.Equal(t, test.expectedValue, node.Value)
				assert.Nil(t, p.Err())
			} else {
				assert.Nil(t, node)
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestParseKeyValue(t *testing.T) {
	var text string
	var node *KeyValue
	var p *Parser

	/*	text = `# Some comments on header
		ABCEDF : "uyfgze fuzy uyezfgezuy " # some comment on eol
	`*/
	text = "X-WASSUP-ULV: 0x400007b220d105f228acb76e   # identifiant Wassup oidval"
	p = NewParserFromString(text, "")
	node = p.parseKeyValue()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestParseValueString(t *testing.T) {
	var node *ValueString
	var p *Parser

	var tests = []struct {
		text          string
		expectedValue string
		error         bool
	}{
		{text: `abcdef`, expectedValue: "abcdef"},
		{text: `abcdef   `, expectedValue: "abcdef"},
		{text: `abcdef 0123456`, expectedValue: "abcdef 0123456"},
		{text: `abcdef#0123456`, expectedValue: "abcdef"},
		{text: `abcdef    #0123456`, expectedValue: "abcdef"},
		{text: `012 345 678   `, expectedValue: "012 345 678"},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseValueString()
			if !test.error {
				assert.Equal(t, test.expectedValue, node.Value)
				assert.Nil(t, p.Err())
			} else {
				assert.Nil(t, node)
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestParseSectionHeader(t *testing.T) {
	var text string
	var node *SectionHeader
	var p *Parser

	text = "[Cookies]"
	p = NewParserFromString(text, "")
	node = p.parseSectionHeader("Cookies")
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestJson(t *testing.T) {
	var node Json
	var p *Parser

	var tests = []struct {
		text  string
		error bool
	}{
		{text: `{
	"id": 0,
    "name": "Frieda",
    "picture": "images/scottish-terrier.jpeg",
    "age": 3,
    "breed": "Scottish Terrier",
    "location": "Lisco, Alabama"} xxxxxxxx`},
		{text: `{"id": 0,"selected": true}`},
		{text: `true xxxxx`},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node, _ = p.parseJson()
			if !test.error {
				assert.NotNil(t, node)
				assert.Nil(t, p.Err())
			} else {
				assert.Nil(t, node)
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestXml(t *testing.T) {
	var p *Parser
	var xml string

	var tests = []struct {
		text  string
		error bool
	}{
		{text: `<catalog>
   <book id="bk101">
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications 
      with XML.</description>
   </book>
</catalog>-----`},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			xml = p.parseXml()
			if !test.error {
				assert.Equal(t, 334, len(xml))
				assert.Nil(t, p.Err())
			} else {
				assert.NotNil(t, p.Err())
			}
		})
	}
}

func TestParseBase64(t *testing.T) {
	var p *Parser

	var tests = []struct {
		encoded string
		decoded string
		error   bool
	}{
		{encoded: `base64,V2VsY29tZSBodXJsIQ==;`, decoded: "Welcome hurl!"},
		{encoded: `base64,V2VsY29tZSBodXJsIQ==;XXXX`, decoded: "Welcome hurl!"},
		{encoded: `base64,TWFuIGlzIGRpc3Rpbmd1aXNoZWQsIG5vdCBvbmx5IGJ5IGhpcyByZWFzb24sIGJ1dCBieSB0aGlz
IHNpbmd1bGFyIHBhc3Npb24gZnJvbSBvdGhlciBhbmltYWxzLCB3aGljaCBpcyBhIGx1c3Qgb2Yg
dGhlIG1pbmQsIHRoYXQgYnkgYSBwZXJzZXZlcmFuY2Ugb2YgZGVsaWdodCBpbiB0aGUgY29udGlu
dWVkIGFuZCBpbmRlZmF0aWdhYmxlIGdlbmVyYXRpb24gb2Yga25vd2xlZGdlLCBleGNlZWRzIHRo
ZSBzaG9ydCB2ZWhlbWVuY2Ugb2YgYW55IGNhcm5hbCBwbGVhc3VyZS4=;`, decoded:"Man is distinguished, not only by his reason," +
			" but by this singular passion from other animals, which is a lust of the mind, that by a perseverance" +
			" of delight in the continued and indefatigable generation of knowledge, exceeds the short vehemence of" +
			" any carnal pleasure."},
		{encoded: `base64,V2VsY29tZSBodXJsIQ==`, error: true},
	}
	for _, test := range tests {
		t.Run(test.encoded, func(t *testing.T) {
			p = NewParserFromString(test.encoded, "")
			value, _ := p.parseBase64()
			if !test.error {
				assert.Equal(t, test.decoded, string(value))
				assert.Nil(t, p.Err())
			} else {
				assert.Nil(t, value)
				assert.NotNil(t, p.Err())
			}
		})
	}

}
