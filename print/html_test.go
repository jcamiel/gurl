package print

import (
	"github.com/stretchr/testify/assert"
	"gurl/ast"
	"testing"
)

func TestHTMLFormatter(t *testing.T) {

	input := `# Sample Hurl
GET https://sample.org
User-Agent: Some Agent
[QueryStringParams]
q: dummy

HTTP/1.1 200`

	expected := `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>Hurl</title>
    </head>
    <body bgcolor="black">
		<pre>
			<code>
<span style="color:#666">1</span>  <span style="color:#8a8a8a"># Sample Hurl</span>
<span style="color:#666">2</span>  <span style="color:#ffaf00">GET</span> <span style="color:#19b2b2">https://sample.org</span>
<span style="color:#666">3</span>  <span style="color:white">User-Agent</span><span style="color:white">:</span> <span style="color:#18b218">Some Agent</span>
<span style="color:#666">4</span>  <span style="color:#b217b2">[QueryStringParams]</span>
<span style="color:#666">5</span>  <span style="color:white">q</span><span style="color:white">:</span> <span style="color:#18b218">dummy</span>
<span style="color:#666">6</span>  
<span style="color:#666">7</span>  <span style="color:white">HTTP/1.1</span> <span style="color:#00afff">200</span>

			</code>
		</pre>
    </body>
</html>`
	p := ast.NewParserFromString(input, "")
	hurl := p.Parse()
	assert.NotNil(t, hurl)
	assert.Nil(t, p.Err())

	f := NewHTMLPrinter()
	output := f.Print(hurl)
	assert.Equal(t, expected, output)
}
