package query

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestEvalXPathHTMLCSuite(t *testing.T) {

	html, _ := os.Open("testdata/sample1.html")
	defer html.Close()

	doc, err := ioutil.ReadAll(html)
	if err != nil {
		return
	}

	_ = EvalXPathXMLC("normalize-space(//head/title)", []byte(doc))
	//_ = EvalXPathXMLC("//head/title", []byte(doc))

}
