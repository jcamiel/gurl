package format

import (
	"github.com/logrusorgru/aurora"
	"gurl/ast"
	"strings"
)

type TermFormatter struct {
	text string
	whitespacesFunc func(string) string
}

func NewTermFormatter(showWhitespaces bool) *TermFormatter {
	var whitespacesFunc func(string) string
	if showWhitespaces {
		whitespacesFunc = visualizeWhitespaces
	} else {
		whitespacesFunc = func (s string) string {
			return s
		}
	}

	f := TermFormatter{"", whitespacesFunc}
	return &f
}

func (p *TermFormatter) Format(hurlFile *ast.HurlFile) string {
	p.text = ""
	ast.Walk(p, hurlFile)
	return p.text
}

func (p *TermFormatter) Visit(node ast.Noder) ast.Visitor {
	switch n := node.(type) {
	case *ast.Body:
		p.text += p.whitespacesFunc(n.Text)
		return nil
	case *ast.Eol:
		p.text += p.whitespacesFunc(n.Value)
		return nil
	case *ast.Whitespaces:
		p.text += p.whitespacesFunc(n.Value)
		return nil
	case *ast.Spaces:
		p.text += p.whitespacesFunc(n.Value)
		return nil
	case *ast.Comment:
		p.text += aurora.Gray(13, n.Value).String()
		return nil
	case *ast.Url:
		p.text += aurora.Cyan(n.Value).String()
		return nil
	case *ast.Method:
		p.text += aurora.Index(214, n.Value).String()
		return nil
	case *ast.Key:
		if n.KeyString != nil {
			p.text += n.Value
		}
		if n.JsonString != nil {
			p.text += n.Value
		}
		return nil
	case *ast.KeyString:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.ValueString:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.JsonString:
		p.text += aurora.Green(n.Text).String()
		return nil
	case *ast.QueryString:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.CookieValue:
		p.text += aurora.Green(n.Value).String()
		return nil
	case *ast.Colon:
		p.text += n.Value
		return nil
	case *ast.SectionHeader:
		p.text += aurora.Magenta(n.Value).String()
		return nil
	case *ast.Version:
		p.text += n.Value
		return nil
	case *ast.QueryType:
		p.text += aurora.Cyan(n.Value).String()
		return nil
	case *ast.PredicateType:
		p.text += aurora.Index(214, n.Value).String()
		return nil
	case *ast.Natural:
		p.text += aurora.Index(39, n.Text).String()
		return nil
	case *ast.Integer:
		p.text += aurora.Index(39, n.Text).String()
		return nil
	case *ast.Bool:
		p.text += aurora.Index(39, n.Text).String()
		return nil
	case *ast.Float:
		p.text += aurora.Index(39, n.Text).String()
		return nil

	}
	return p
}

func visualizeWhitespaces(s string) string {
	whites := map[string]string{
		" ":  "_",
		"\n": "\u21b5\n",
		"\t": "\u2192   ",
	}
	for src, dst := range whites {
		s = strings.ReplaceAll(s, src, aurora.Gray(3, dst).String())
	}
	return s
}
