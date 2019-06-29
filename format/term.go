package format

import (
	"github.com/logrusorgru/aurora"
	"gurl/ast"
)

type TermFormatter struct {
	text string
}

func NewTermFormatter() *TermFormatter {
	f := TermFormatter{}
	return &f
}

func (p *TermFormatter) ToText(hurlFile *ast.HurlFile) string {
	p.text = ""
	ast.Walk(p, hurlFile)
	return p.text
}

func (p *TermFormatter) Visit(node ast.Noder) ast.Visitor {

	switch n := node.(type) {
	case *ast.Eol:
		p.text += n.Text
		return nil
	case *ast.Comment:
		p.text += aurora.Gray(11, n.Text).String()
		return nil
	case *ast.Whitespaces:
		p.text += n.Text
		return nil
	case *ast.Url:
		p.text += aurora.Yellow(n.Text).String()
		return nil
	case *ast.Method:
		p.text += aurora.Magenta(n.Value).String()
		return nil
	case *ast.KeyString:
		p.text += aurora.Green(n.Text).String()
		return nil
	case *ast.JsonString:
		p.text += aurora.Green(n.Text).String()
		return nil
	case *ast.Spaces:
		p.text += n.Text
		return nil
	case *ast.Colon:
		p.text += n.Text
		return nil
	}
	return p
}
