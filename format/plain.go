package format

import (
	"gurl/ast"
)

type PlainFormatter struct {
	text string
}

func (p *PlainFormatter) ToText(hurlFile *ast.HurlFile) string {
	p.text = ""
	ast.Walk(p, hurlFile)
	return p.text
}

func (p *PlainFormatter) Visit(node ast.Node) ast.Visitor {

	switch n := node.(type) {
	case *ast.Eol:
		p.text += n.Text
		return nil
	case *ast.Comment:
		p.text += n.Text
		return nil
	case *ast.Whitespaces:
		p.text += n.Text
		return nil
	case *ast.Method:
		p.text += n.Value
		return nil
	case *ast.Spaces:
		p.text += n.Text
		return nil
	}
	return p
}
