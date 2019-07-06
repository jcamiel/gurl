package ast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMethodSucceed(t *testing.T) {
	text := "POST http://google.com"
	p := NewParserFromString(text, "")
	node := p.parseMethod()
	assert.Nil(t, p.Err())
	assert.Equal(t, &Method{
		Node{Position{0, 1, 1}, Position{4, 1, 5}},
		"POST",
	}, node, "POST should be parsed")
}

func TestParseMethodFailed(t *testing.T) {
	text := "ABCEDFGHIJKLM"
	p := NewParserFromString(text, "")
	node := p.parseMethod()
	assert.Nil(t, node)
	assert.NotNil(t, p.Err())
}

func TestParseRequest(t *testing.T) {
	var node *Request
	var p *Parser

	var tests = []struct {
		text string
	}{
		{"GET http://www.example.org"},
		{"GET\u0020http://www.example.org"},
		{"GET http://www.example.org\t# Some comment"},
		{"GET http://www.example.org/foo.html#bar # Some comment"},
		{`# Some comment on request
# ----------
	GET http://www.example.org/foo.html#bar # Some comment
`},
		{`GET {{orange_url}}/demenagement/planifier
User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1
	X-WASSUP-AOL: 10
	X-WASSUP-UIT: 1
	X-WASSUP-ULV: 0x7125a9223bae00010000073f
	X-WASSUP-DSN: STANY AISSAOUI
	X-WASSUP-SAI: 115651101
`},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			p = NewParserFromString(test.text, "")
			node = p.parseRequest()
			assert.NotNil(t, node)
			assert.Nil(t, p.Err())
		})
	}
}

func TestParseFailed(t *testing.T) {
	var text string
	var p *Parser

	text = "\n\nPOSThttp://google.com"
	p = NewParserFromString(text, "")
	_ = p.parseRequest()
	assert.NotNil(t, p.Err())
}

func TestParseHurlFile(t *testing.T) {
	var text string
	var node *HurlFile
	var p *Parser

	text = `# GENERATED - DO NOT MODIFY
# =========================
# 
# On teste un parcours de déménagement avec un client Internet Sosh Mobile + Livebox.
# On vérifie que l'on est bien redirigé vers la page correspondante
# sur l'espace client mobile.
# ---------------------------------------

# 
# Login Sur Wt Proxy.
GET https://auth.orange.localhost:3443/r/Oid_identification
User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1
[QueryStringParams]
wassup: Sylvie_Caniou

HTTP/1.1 302


# 
# On vérifie que les clients Sosh sont redirigés depuis la page planifier
# vers l'espace client avec le bon numéro de contrat.
GET {{orange_url}}/demenagement/planifier
User-Agent: Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1

HTTP/1.1 302
[Asserts]
header Location equals "https://sso.orange.fr/espace-client/m/?page=demenagement-demande&MCO=SOH&idContrat=9003384900"`
	p = NewParserFromString(text, "")
	node = p.parseHurlFile()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestParseHeaders(t *testing.T) {
	var text string
	var node *Headers
	var p *Parser

	text = `key0 : value0
key1 : value1
key2 : value2
key3 : value3`
	p = NewParserFromString(text, "")
	node = p.parseHeaders()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestCookies( t *testing.T) {
	var text string
	var node *Cookies
	var p *Parser

	text = `[Cookies]
	cookieA : valueA # Some comment on value A
	cookieB : valueB
	cookieC : valueC
`
	p = NewParserFromString(text, "")
	node = p.parseCookies()
	assert.NotNil(t, node)
	assert.Nil(t, p.Err())
}

func TestBody(t *testing.T) {
	var node *Body
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
			node = p.parseBody()
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