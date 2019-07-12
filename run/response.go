package run

import (
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"golang.org/x/net/html"
	"gurl/ast"
	"log"
	"net/http"
)

func isStatusCodeValid(r *ast.Response, resp *http.Response) bool {
	if resp.StatusCode != r.Status.Value.Value {
		log.Print(fmt.Sprintf("Assert status failed expected: %d actual: %d", r.Status.Value.Value, resp.StatusCode))
		return false
	}
	return true
}

func captureVariables(captures *ast.Captures, resp *http.Response) {
	for _, c := range captures.Captures {
		_ = c.Key.Value
		_, _ = evaluateQuery(c.Query, resp)
	}
}

func evaluateQuery(q *ast.Query, resp *http.Response) (string, error) {
	if q.Type.Value == "xpath" {
		expr, err := xpath.Compile(q.Expr.Value)
		if err != nil {
			return "", nil
		}
		node, _ := html.Parse(resp.Body)
		root := htmlquery.FindOne(node, "//html")
		nav := htmlquery.CreateXPathNavigator(root)
		v := expr.Evaluate(nav)
		fmt.Println(v)
	}
	return "", nil
}
