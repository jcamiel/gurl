package ast

import "fmt"

type Visitor interface {
	Visit(node Noder) (w Visitor)
}

func Walk(v Visitor, node Noder) {

	if v = v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {

	case *HurlFile:
		if n.Whitespaces != nil {
			Walk(v, n.Whitespaces)
		}
		for _, e := range n.Entries {
			Walk(v, e)
		}
	case *Entry:
		Walk(v, n.Request)
	case *Request:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.Method)
		Walk(v, n.Spaces0)
		Walk(v, n.Url)
		if n.Spaces1 != nil {
			Walk(v, n.Spaces1)
		}
		if n.Comment != nil {
			Walk(v, n.Comment)
		}
		Walk(v, n.Eol)

	case *Comments:
		for _, c := range n.CommentLines {
			Walk(v, c)
		}
	case *CommentLine:
		Walk(v, n.Comment)
		Walk(v, n.Eol)
		if n.Whitespaces != nil {
			Walk(v, n.Whitespaces)
		}
	case *Spaces, *Method, *Url:
		// do nothing
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
