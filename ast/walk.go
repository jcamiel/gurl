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
		if n.Whitespaces != nil {
			Walk(v, n.Whitespaces)
		}
		if n.Response != nil {
			Walk(v, n.Response)
		}
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
		if n.Body != nil {
			Walk(v, n.Body)
		}
	case *Response:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.Version)
		Walk(v, n.Spaces0)
		Walk(v, n.Status)
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
		if n.Captures != nil {
			Walk(v, n.Captures)
		}
	case *Headers:
		for _, h := range n.Headers {
			Walk(v, h)
		}
	case *Captures:
		if n.Comments != nil {
			Walk(v, n.Comments)
		}
		Walk(v, n.SectionHeader)
		if n.Spaces != nil {
			Walk(v, n.Spaces)
		}
		Walk(v, n.Eol)
		for _, c := range n.Captures {
			Walk(v, c)
		}
	case *Capture:
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
		Walk(v, n.Query)
		if n.Spaces2 != nil {
			Walk(v, n.Spaces2)
		}
		if n.Comment != nil {
			Walk(v, n.Comment)
		}
		Walk(v, n.Eol)
	case *Query:
		if n.Spaces0 != nil {
			Walk(v, n.Spaces0)
		}
		Walk(v, n.QueryType)
		if n.Spaces1 != nil {
			Walk(v, n.Spaces1)
		}
		if n.QueryExpr != nil {
			Walk(v, n.QueryExpr)
		}
		if n.Spaces2 != nil {
			Walk(v, n.Spaces2)
		}
	case *QueryExpr:
		if n.QueryString != nil {
			Walk(v, n.QueryString)
		}
		if n.JsonString != nil {
			Walk(v, n.JsonString)
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
	case *Status:
		Walk(v, n.Value)
	case *Eol, *Whitespaces, *Comment, *Spaces, *Method, *Url, *KeyString, *JsonString, *ValueString, *Colon,
		*SectionHeader, *CookieValue, *Body, *Version, *Natural, *QueryType, *QueryString:
		// do nothing
	default:
		panic(fmt.Sprintf("ast.Walk: unexpected node type %T", n))
	}

	v.Visit(nil)
}
