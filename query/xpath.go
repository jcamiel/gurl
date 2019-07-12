package query

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"golang.org/x/net/html"
)

// Evaluate returns the result of the expression.
// The result type of the expression is one of the follow: bool,float64,string,[]string (representing collections)).
func EvalXPathHTML(expr string, body []byte) (interface{}, error) {
	e, err := xpath.Compile(expr)
	if err != nil {
		return "", nil
	}
	node, _ := html.Parse(bytes.NewReader(body))
	root := htmlquery.FindOne(node, "//html")
	nav := htmlquery.CreateXPathNavigator(root)
	val := e.Evaluate(nav)

	switch v := val.(type) {
	case *xpath.NodeIterator:
		items := []string{}
		for {
			if v.MoveNext() {
				items = append(items, v.Current().Value())
			} else {
				break
			}
		}
		return items, nil
	default:
		return val, nil
	}
}
