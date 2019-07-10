package run

import (
	"fmt"
	"gurl/ast"
	"net/http"
)

type HttpRunner struct {
	client *http.Client
	variables map[string]string
}

func NewHttpRunner() *HttpRunner {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	variables := make(map[string]string)
	return &HttpRunner{client, variables}
}

func (h *HttpRunner) Run(hurlFile *ast.HurlFile) {
	ast.Walk(h, hurlFile)
}

func (h *HttpRunner) Visit(node ast.Noder) ast.Visitor {
	switch n := node.(type) {
	case *ast.Request:
		h.doRequest(n)
		return nil
	}
	return h
}

func (h *HttpRunner)doRequest(request *ast.Request) {

	/*"GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",*/
	switch request.Method.Value {
	case "GET":
		fmt.Printf("GET %s\n", request.Url.Value)
	case "HEAD":
		fmt.Printf("HEAD %s\n", request.Url.Value)
	case "POST":
		fmt.Printf("POST %s\n", request.Url.Value)
	case "PUT":
		fmt.Printf("PUT %s\n", request.Url.Value)
	case "DELETE":
		fmt.Printf("DELETE %s\n", request.Url.Value)
	case "CONNECT":
		fmt.Printf("CONNECT %s\n", request.Url.Value)
	case "OPTIONS":
		fmt.Printf("OPTIONS %s\n", request.Url.Value)
	case "TRACE":
		fmt.Printf("TRACE %s\n", request.Url.Value)
	case "PATCH":
		fmt.Printf("PATCH %s\n", request.Url.Value)
	}
}

