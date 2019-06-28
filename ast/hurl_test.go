package ast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMethodSucceed(t *testing.T) {
	text := "POST http://google.com"
	p := NewParserFromString(text, "")
	node, err := p.parseMethod()

	assert.Nil(t, err)
	assert.Equal(t, &Method{
		Position{0, 1},
		Position{4, 1},
		"POST",
	}, node, "POST should be parsed")
}

func TestParseMethodFailed(t *testing.T) {
	text := "ABCEDFGHIJKLM"
	p := NewParserFromString(text, "")
	node, err := p.parseMethod()
	assert.Nil(t, node)
	assert.NotNil(t, err)
}

func TestParseRequest(t *testing.T) {

	var text string
	var node *Request
	var p *Parser

	text = "GET	http://www.example.org"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET\u0020http://www.example.org"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET	http://www.example.org	# Some comment"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = "GET http://www.example.org/foo.html#bar # Some comment"
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)

	text = `# Some comment on request
# ----------
	GET http://www.example.org/foo.html#bar # Some comment
`
	p = NewParserFromString(text, "")
	node, _ = p.parseRequest()
	assert.NotNil(t, node)
}

func TestParseFailed(t *testing.T) {

	var text string
	var err error
	var p *Parser

	text = "\n\nPOSThttp://google.com"
	p = NewParserFromString(text, "")
	_, err = p.parseRequest()
	fmt.Println(err)
	assert.NotNil(t, err)
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
	node, _ = p.parseHurlFile()
	assert.NotNil(t, node)
}