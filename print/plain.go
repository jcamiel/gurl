package print

import (
	"gurl/ast"
)

type PlainPrinter struct {
	text string
}

func NewPlainPrinter() *PlainPrinter {
	return &PlainPrinter{}
}


func (p *PlainPrinter) Print(hurlFile *ast.HurlFile) string {
	p.text = ""
	ast.Walk(p, hurlFile)
	return p.text
}

func (p *PlainPrinter) Visit(node ast.Noder) ast.Visitor {

	switch n := node.(type) {
	case *ast.Eol:
		p.text += n.Value
		return nil
	case *ast.Comment:
		p.text += n.Value
		return nil
	case *ast.Whitespaces:
		p.text += n.Value
		return nil
	case *ast.Method:
		p.text += n.Value
		return nil
	case *ast.Spaces:
		p.text += n.Value
		return nil
	}
	return p
}
