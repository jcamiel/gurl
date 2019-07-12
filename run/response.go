package run

import (
	"errors"
	"fmt"
	"gurl/ast"
	"gurl/query"
	"io/ioutil"
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

func captureVariables(captures *ast.Captures, resp *http.Response) (map[string]string, error) {
	vars := map[string]string{}

	for _, c := range captures.Captures {
		name := c.Key.Value
		val, err := evaluateQuery(c.Query, resp)
		if err != nil {
			return nil, err
		}
		switch v := val.(type) {
		case string:
			vars[name] = v
		default:
			return nil, errors.New("unsupported ")
		}
	}
	return vars, nil
}

func evaluateQuery(qry *ast.Query, resp *http.Response) (interface{}, error) {
	switch qry.Type.Value {
	case "xpath":
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return query.EvalXPathHTML(qry.Expr.Value, body)
	default:
		return nil, errors.New("unsupported query type")
	}
}
