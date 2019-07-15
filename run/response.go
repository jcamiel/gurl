package run

import (
	"bytes"
	"errors"
	"fmt"
	"gurl/ast"
	"gurl/query"
	"io/ioutil"
	"log"
	"net/http"
)

func (h *HttpRunner) checkResponse(r *ast.Response, resp *http.Response) error {

	_ = isStatusCodeValid(r, resp)

	if r.Captures != nil {
		v, err := captureVariables(r.Captures, resp)
		if err != nil {
			return err
		}
		h.variables = concatenateMaps(h.variables, v)
	}
	if r.Asserts != nil {
		res, err := h.getAssertsResults(r.Asserts.Asserts, resp)
		if err != nil {
			return err
		}

		for _, r := range res {
			fmt.Println(r)
		}
	}

	return nil
}

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
		body, err := body(resp)
		if err != nil {
			return nil, err
		}
		return query.EvalXPathHTML(qry.Expr.Value, body)
	default:
		return nil, errors.New("unsupported query type")
	}
}

func (h *HttpRunner) getAssertsResults(asserts []*ast.Assert, resp *http.Response) ([]*AssertResult, error) {

	results := []*AssertResult{}

	for _, a := range asserts {
		actual, err := evaluateQuery(a.Query, resp)
		if err != nil {
			return nil, err
		}
		switch a.Predicate.Type.Value {
		case "equals":
			r := assertEquals(a.Predicate, actual)
			results = append(results, r)
		case "matches":
			return nil, errors.New("unsupported query type")
		case "startsWith":
			return nil, errors.New("unsupported query type")
		case "contains":
			return nil, errors.New("unsupported query type")
		}
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
