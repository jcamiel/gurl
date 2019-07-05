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
		if n.Headers != nil {
			Walk(v, n.Headers)
		}
		if n.Cookies != nil {
			Walk(v, n.Cookies)
		}
		if n.QsParams != nil {
			Walk(v, n.QsParams)
		}
		if n.FormParams != nil {
			Walk(v, n.FormParams)
		}
	case *Headers:
		for _, h := range n.Headers {
			Walk(v, h)
		}
	case *Cookies:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.SectionHeader)
		if n.Spaces != nil {
			Walk(v, n.Spaces)
		}
		Walk(v, n.Eol)
		for _, c := range n.Cookies {
			Walk(v, c)
		}
	case *QsParams:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.SectionHeader)
		if n.Spaces != nil {
			Walk(v, n.Spaces)
		}
		Walk(v, n.Eol)
		for _, p := range n.Params {
			Walk(v, p)
		}
	case *FormParams:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.SectionHeader)
		if n.Spaces != nil {
			Walk(v, n.Spaces)
		}
		Walk(v, n.Eol)
		for _, p := range n.Params {
			Walk(v, p)
		}
	case *Cookie:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.Key)
		if n.Spaces0 != nil {
			Walk(v, n.Spaces0)
		}
		Walk(v, n.Colon)
		if n.Spaces1 != nil {
			Walk(v, n.Spaces1)
		}
		Walk(v, n.CookieValue)
		if n.Spaces2 != nil {
			Walk(v, n.Spaces2)
		}
		if n.Comment != nil {
			Walk(v, n.Comment)
		}
		Walk(v, n.Eol)
	case *KeyValue:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.Key)
		if n.Spaces0 != nil {
			Walk(v, n.Spaces0)
		}
		Walk(v, n.Colon)
		if n.Spaces1 != nil {
			Walk(v, n.Spaces1)
		}
		Walk(v, n.Value)
		if n.Spaces2 != nil {
			Walk(v, n.Spaces2)
		}
		if n.Comment != nil {
			Walk(v, n.Comment)
		}
		Walk(v, n.Eol)
	case *Key:
		if n.KeyString != nil {
			Walk(v, n.KeyString)
		}
		if n.JsonString != nil {
			Walk(v, n.JsonString)
		}
	case *Value:
		if n.ValueString != nil {
			Walk(v, n.ValueString)
		}
		if n.JsonString != nil {
			Walk(v, n.JsonString)
		}
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
	case *Eol, *Whitespaces, *Comment, *Spaces, *Method, *Url, *KeyString, *JsonString, *ValueString, *Colon,
		*SectionHeader, *CookieValue:
		// do nothing
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
