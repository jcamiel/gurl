package run

import (
	"fmt"
	"gurl/ast"
	"gurl/template"
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
	//variables := make(map[string]string)
	variables := map[string]string{
		"root_url": "https://localhost:8080",
	}
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

	url, err := template.Render(request.Url.Value, h.variables)
	if err != nil {
		return
	}

	switch request.Method.Value {
	case "GET":
		fmt.Printf("GET %s\n", url)
	case "HEAD":
		fmt.Printf("HEAD %s\n", url)
	case "POST":
		fmt.Printf("POST %s\n", url)
	case "PUT":
		fmt.Printf("PUT %s\n", url)
	case "DELETE":
		fmt.Printf("DELETE %s\n", url)
	case "CONNECT":
		fmt.Printf("CONNECT %s\n", url)
	case "OPTIONS":
		fmt.Printf("OPTIONS %s\n", url)
	case "TRACE":
		fmt.Printf("TRACE %s\n", url)
	case "PATCH":
		fmt.Printf("PATCH %s\n", url)
	}
}

