package print

import (
	"github.com/logrusorgru/aurora"
	"gurl/ast"
	"strings"
)

type TermPrinter struct {
	text string
}

func NewTermPrinter() *TermPrinter {
	f := TermPrinter{}
	return &f
}

func (p *TermPrinter) Print(hurlFile *ast.HurlFile) string {
	p.text = ""
	ast.Walk(p, hurlFile)
	return p.text
}

func (p *TermPrinter) Visit(node ast.Noder) ast.Visitor {

	switch n := node.(type) {
	case *ast.Body:
		p.text += n.Text
	case *ast.Eol:
		p.text += aurora.Gray(3, visualizeWhitespaces(n.Value)).String()
		return nil
	case *ast.Whitespaces:
		p.text += aurora.Gray(3, visualizeWhitespaces(n.Value)).String()
		return nil
	case *ast.Spaces:
		p.text += aurora.Gray(3, visualizeWhitespaces(n.Value)).String()
		return nil
	case *ast.Comment:
		p.text += aurora.Gray(13, n.Value).String()
		return nil
	case *ast.Url:
		p.text += aurora.Yellow(n.Value).String()
		return nil
	case *ast.Method:
		p.text += aurora.Magenta(n.Value).String()
		return nil
	case *ast.KeyString:
		p.text += aurora.Cyan(n.Value).String()
		return nil
	case *ast.ValueString:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.JsonString:
		p.text += aurora.Green(n.Text).String()
		return nil
	case *ast.CookieValue:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.Colon:
		p.text += n.Value
		return nil
	case *ast.SectionHeader:
		p.text += aurora.Blue(n.Value).String()
		return nil
	}
	return p
}

func visualizeWhitespaces(s string) string {
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "\n", "\u21b5\n")
	s = strings.ReplaceAll(s, "\t", "\u2192   ")
	return s
}
