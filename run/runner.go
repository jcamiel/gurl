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

func NewHttpRunner() *HttpRunner {
	//variables := make(map[string]string)
	variables := map[string]string{
		"root_url":                         "http://localhost:8080",
		"orange_url":                       "https://myshop.orange.localhost:3443",
		"sosh_url":                         "https://myshop.sosh.localhost:3443",
		"https":                            "https//auth.orange.localhost:3443/r/Oid_identification",
		"wishedActivationDate":             "26/07/2019",
		"wishedActivationDatePlusOneDay":   "27/07/2019",
		"wishedActivationDatePlusFiveDays": "31/07/2019",
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
			fmt.Println(r)
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
