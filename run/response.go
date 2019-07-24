package run

import (
	"bytes"
	"errors"
	"fmt"
	"gurl/ast"
	"gurl/query"
	"io/ioutil"
	"net/http"
)

func (h *HttpRunner) checkResponse(r *ast.Response, resp *http.Response) ([]*AssertResult, error) {

	// First capture variables
	if r.Captures != nil {
		v, err := captureVariables(r.Captures, resp)
		if err != nil {
			return nil, err
		}
		h.variables = concatenateMaps(h.variables, v)
	}

	// Test status code (it's the first assert) and the remaining response asserts.
	var asserts []*AssertResult
	asserts = append(asserts, isStatusCodeValid(r, resp))

	if r.Asserts != nil {
		a, err := h.getAssertsResults(r.Asserts.Asserts, resp)
		if err != nil {
			return nil, err
		}
		asserts = append(asserts, a...)
	}

	return asserts, nil
}

func isStatusCodeValid(r *ast.Response, resp *http.Response) *AssertResult {
	ret := resp.StatusCode == r.Status.Value.Value
	if ret {
		return &AssertResult{ok:true}
	} else {
		msg := fmt.Sprintf("assert status failed expected: %d actual: %d", r.Status.Value.Value, resp.StatusCode)
		return &AssertResult{ok:false, msg:msg}
	}
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
		body, err := body(resp)
		if err != nil {
			return nil, err
		}
		// TODO: depending on the response content type, we should
		//  evaluate a HTML xpath query, or a XML xpath query.
		return query.EvalXPathHTML(qry.Expr.Value, body)
	case "jsonpath":
		body, err := body(resp)
		if err != nil {
			return nil, err
		}
		if !query.IsJSON(body) {
			return nil, errors.New("valid JSON body is expected for jsonpath query")
		}
		return query.EvalJSONPath(qry.Expr.Value, body)
	default:
		return nil, errors.New("unsupported query type")
	}
}

func (h *HttpRunner) getAssertsResults(asserts []*ast.Assert, resp *http.Response) ([]*AssertResult, error) {

	var results []*AssertResult

	for _, a := range asserts {
		var r *AssertResult
		actual, err := evaluateQuery(a.Query, resp)
		if err != nil {
			r = &AssertResult{msg:fmt.Sprintf("invalid query: %s", err)}
		} else {
			switch a.Predicate.Type.Value {
			case "equals":
				r = assertEquals(a.Predicate, actual)
			case "matches":
				r = &AssertResult{msg:"matches query unsupported"}
			case "startsWith":
				r = &AssertResult{msg:"startsWith query unsupported"}
			case "contains":
				r = &AssertResult{msg:"contains query unsupported"}
			}
		}
		results = append(results, r)
	}
	return results, nil
}

func body(resp *http.Response) ([]byte, error){
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Restore the io.ReadCloser to its original state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return body, nil
}
