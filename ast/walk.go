package ast

import "fmt"

type Visitor interface {
	Visit(node Node) (w Visitor)
}

func Walk(v Visitor, node Node) {

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
	case *Method:
		// do nothing
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
