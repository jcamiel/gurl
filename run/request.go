package run

import (
	"bytes"
	"fmt"
	"gurl/ast"
	"gurl/template"
	"net/http"
	"net/http/httputil"
	urlpkg "net/url"
)

func (h *HttpRunner) doRequest(client *http.Client, r *ast.Request) (*http.Response, error) {

	url, err := template.Render(r.Url.Value, h.variables)
	if err != nil {
		return nil, err
	}

	method := r.Method.Value
	var body []byte

	// Construct body from FormParams if any.
	if r.FormParams != nil {
		body, err = h.bodyFromForm(r.FormParams.Params)
		if err != nil {
			return nil, err
		}
	}

	// TODO: traiter form vs body

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	if r.Headers != nil {
		err := h.addHeaders(req, r.Headers)
		if err != nil {
			return nil, err
		}
	}
	if r.QsParams != nil {
		err := h.addQueryParams(req, r.QsParams.Params)
		if err != nil {
			return nil, err
		}
	}
	if r.FormParams != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	h.dumpRequest(req)
	return resp, nil
}

func (h *HttpRunner) addHeaders(req *http.Request, headers *ast.Headers) error {
	for _, hd := range headers.Headers {
		name, err := template.Render(hd.Key.Value, h.variables)
		if err != nil {
			return err
		}
		value, err := template.Render(hd.Value.Value, h.variables)
		if err != nil {
			return err
		}
		req.Header.Add(name, value)
	}
	return nil
}

func (h *HttpRunner) addQueryParams(req *http.Request, params []*ast.KeyValue) error {
	q := req.URL.Query()
	for _, p := range params {
		name, err := template.Render(p.Key.Value, h.variables)
		if err != nil {
			return err
		}
		value, err := template.Render(p.Value.Value, h.variables)
		if err != nil {
			return err
		}
		q.Add(name, value)
	}
	req.URL.RawQuery = q.Encode()
	return nil
}

func (h *HttpRunner) bodyFromForm(params []*ast.KeyValue) ([]byte, error) {
	form := urlpkg.Values{}
	for _, p := range params {
		name, err := template.Render(p.Key.Value, h.variables)
		if err != nil {
			return nil, err
		}
		value, err := template.Render(p.Value.Value, h.variables)
		if err != nil {
			return nil, err
		}
		form.Add(name, value)
	}
	return []byte(form.Encode()), nil
}

func (h *HttpRunner) dumpRequest(req *http.Request) {
	// Save a copy of this request for debugging.
	fmt.Println("------------------")
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(requestDump))
	fmt.Println("Captures:")
	for k, v := range h.variables {
		fmt.Printf("%s: %s\n", k, v)
	}
}