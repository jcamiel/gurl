package run

import (
	"fmt"
	"gurl/ast"
	"gurl/template"
	"log"
	"net/http"
)

type HttpRunner struct {
	variables map[string]string
}

func NewHttpRunner() *HttpRunner {
	//variables := make(map[string]string)
	variables := map[string]string{
		"root_url": "http://localhost:8080",
	}
	return &HttpRunner{variables}
}

func (h *HttpRunner) Run(hurl *ast.HurlFile) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, e := range hurl.Entries {
		resp, err := h.doRequest(client, e.Request)
		if err != nil {
			panic(err)
		}

		if e.Response != nil {
			checkResponse(e.Response, resp)
		}

		_ = resp.Body.Close()
	}
}

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

func checkResponse(r *ast.Response, resp *http.Response) {
	if resp.StatusCode != r.Status.Value.Value {
		log.Print(fmt.Sprintf("Assert status failed expected: %d actual: %d", r.Status.Value.Value, resp.StatusCode))
	}
}
