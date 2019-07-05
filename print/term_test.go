package print

import (
	"github.com/stretchr/testify/assert"
	"gurl/ast"
	"testing"
)

func TestTermFormatter(t *testing.T) {

	input := `# GENERATED - DO NOT MODIFY
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

	p := ast.NewParserFromString(input, "")
	hurl := p.Parse()
	assert.NotNil(t, hurl)
	assert.Nil(t, p.Err())

	formatter := NewTermPrinter()
	output := formatter.Print(hurl)
	assert.NotNil(t, output)
}
