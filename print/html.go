package print

import (
	"fmt"
	"gurl/ast"
	"math"
	"strings"
)

type HTMLPrinter struct {
	text string
}

func NewHTMLPrinter() *HTMLPrinter {
	return &HTMLPrinter{}
}

func (p *HTMLPrinter) Print(hurlFile *ast.HurlFile) string {

	const template = `<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8">
        <title>Hurl</title>
    </head>
    <body bgcolor="black">
		<pre>
			<code>
%s
			</code>
		</pre>
    </body>
</html>`

	ast.Walk(p, hurlFile)
	text := addLineNumber(p.text)
	return fmt.Sprintf(template, text)
}

func (p *HTMLPrinter) Visit(node ast.Noder) ast.Visitor {
	switch n := node.(type) {
	case *ast.Body:
		p.text += white(n.Text)
		return nil
	case *ast.Eol:
		p.text += n.Value
		return nil
	case *ast.Whitespaces:
		p.text += n.Value
		return nil
	case *ast.Spaces:
		p.text += n.Value
		return nil
	case *ast.Comment:
		p.text += gray(n.Value)
		return nil
	case *ast.Url:
		p.text += cyan(n.Value)
		return nil
	case *ast.Method:
		p.text += orange(n.Value)
		return nil
	case *ast.Key:
		if n.KeyString != nil {
			p.text += white(n.Value)
		}
		if n.JsonString != nil {
			p.text += white(n.Value)
		}
		return nil
	case *ast.KeyString:
		p.text += green(n.Value)
		return nil
	case *ast.ValueString:
		p.text += green(n.Value)
		return nil
	case *ast.JsonString:
		p.text += green(n.Text)
		return nil
	case *ast.QueryString:
		p.text += green(n.Value)
		return nil
	case *ast.CookieValue:
		p.text += green(n.Value)
		return nil
	case *ast.Colon:
		p.text += white(n.Value)
		return nil
	case *ast.SectionHeader:
		p.text += magenta(n.Value)
		return nil
	case *ast.Version:
		p.text += white(n.Value)
		return nil
	case *ast.QueryType:
		p.text += cyan(n.Value)
		return nil
	case *ast.PredicateType:
		p.text += orange( n.Value)
		return nil
	case *ast.Natural:
		p.text += lightBlue(n.Text)
		return nil
	case *ast.Integer:
		p.text += lightBlue(n.Text)
		return nil
	case *ast.Bool:
		p.text += lightBlue(n.Text)
		return nil
	case *ast.Float:
		p.text += lightBlue(n.Text)
		return nil
	}
	return p
}


func white(text string) string {
	return span(text, "white")
}

func cyan(text string) string {
	return span(text, "#19b2b2")
}

func gray(text string) string {
	return span(text, "#8a8a8a")
}

func darkGray(text string) string {
	return span(text, "#666")
}

func orange(text string) string {
	return span(text, "#ffaf00")
}

func green(text string) string {
	return span(text, "#18b218")
}

func magenta(text string) string {
	return span(text, "#b217b2")
}

func lightBlue(text string) string {
	return span(text, "#00afff")
}

func span(text string, color string) string {
	return fmt.Sprintf("<span style=\"color:%s\">%s</span>", color, text)
}

func addLineNumber(text string) string {
	lines := strings.Split(text, "\n")
	width := 1 + int(math.Log10(float64(len(lines))))
	var out string
	for i, l := range lines {
		out += darkGray(fmt.Sprintf("%*d", width, i+1))
		out += "  "
		out += l
		out += "\n"
	}
	return out
}