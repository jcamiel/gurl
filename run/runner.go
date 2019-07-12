package run

import (
	"gurl/ast"
	"net/http"
	"net/http/cookiejar"
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

func (h *HttpRunner) Run(hurl *ast.HurlFile) error {

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, e := range hurl.Entries {
		if err := h.processEntry(client, e); err != nil {
			return err
		}
	}
	return nil
}

func (h *HttpRunner) processEntry(client *http.Client, e *ast.Entry) error {
	resp, err := h.doRequest(client, e.Request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if e.Response != nil {
		_ = isStatusCodeValid(e.Response, resp)
		if e.Response.Captures != nil {
			v, err := captureVariables(e.Response.Captures, resp)
			if err != nil {
				return err
			}
			h.variables = concatenateMaps(h.variables, v)
		}
	}
	return nil
}

func concatenateMaps(a map[string]string, b map[string]string) map[string]string {
	for k, v := range b {
		a[k] = v
	}
	return a
}
