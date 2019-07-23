package query

import (
	"bufio"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestEvalXPathHTMLSuite(t *testing.T) {

	html, _ := os.Open("testdata/sample1.html")
	defer html.Close()

	doc, err := ioutil.ReadAll(html)
	check(err)

	fcsv, _ := os.Open("testdata/sample1.csv")
	defer fcsv.Close()

	s := bufio.NewScanner(fcsv)
	for s.Scan() {
		col := strings.Split(s.Text(), "|")
		test := strings.TrimSpace(col[0])
		expected := strings.TrimSpace(col[1])
		actual, err := EvalXPathHTML(test, []byte(doc))
		assert.Nil(t, err, test)

		if expected == "true" {
			v, ok := actual.(bool)
			assert.True(t, ok, test)
			assert.True(t, v, test)
		} else if expected == "false" {
			v, ok := actual.(bool)
			assert.True(t, ok, test)
			assert.False(t, v, test)
		} else if strings.HasPrefix(expected, "\"") {
			v, ok := actual.(string)
			assert.True(t, ok, test)
			var exp string
			err := json.Unmarshal([]byte(expected), &exp)
			check(err)
			assert.Equal(t, exp, v, test)
		} else {
			v, ok := actual.(float64)
			assert.True(t, ok, test)
			exp, err := strconv.ParseFloat(expected, 64)
			check(err)
			assert.Equal(t, exp, v, test)
		}
	}
	check(s.Err())
}