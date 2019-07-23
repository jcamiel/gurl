package run

import (
	"fmt"
	"gurl/ast"
	"net/http"
	"net/http/cookiejar"
)

type HttpRunner struct {
	variables map[string]string
}

type AssertResult struct {
	ok  bool
	msg string
}

func (a *AssertResult) String() string {
	if a.ok {
		return "success"
	} else {
		return fmt.Sprintf("failed, %s", a.msg)
	}
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
		res, err := h.checkResponse(e.Response, resp)
		if err != nil {
			return err
		}
		for _, r := range res {
			if !r.ok {
				fmt.Println(r)
			}
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
