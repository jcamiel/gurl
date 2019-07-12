package run

import (
	"fmt"
	"gurl/ast"
	"gurl/template"
	"net/http"
)

func (h *HttpRunner) doRequest(client *http.Client, r *ast.Request) (*http.Response, error) {

	url, err := template.Render(r.Url.Value, h.variables)
	if err != nil {
		return nil, err
	}

	method := r.Method.Value
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if r.Headers != nil {
		addHeaders(req, r.Headers)
	}
	if r.QsParams != nil {
		addQueryParams(req, r.QsParams.Params)
	}

	fmt.Printf("%s %s\n", method, url)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func addHeaders(req *http.Request, headers *ast.Headers) {
	for _, h := range headers.Headers {
		req.Header.Add(h.Key.Value, h.Value.Value)
	}
}

func addQueryParams(req *http.Request, params []*ast.KeyValue) {
	q := req.URL.Query()
	for _, p := range params {
		q.Add(p.Key.Value, p.Value.Value)
	}
	req.URL.RawQuery = q.Encode()
}
