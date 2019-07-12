package run

import (
	"gurl/ast"
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
			_ = isStatusCodeValid(e.Response, resp)
			if e.Response.Captures != nil {
				captureVariables(e.Response.Captures, resp)
			}
		}
		_ = resp.Body.Close()
	}
}





