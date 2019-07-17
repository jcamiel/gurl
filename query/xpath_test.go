package query

import (
	"bufio"
	"encoding/json"
	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xpath"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"strings"
	"testing"
)

func TestEvalXPathHTML(t *testing.T) {

	const html = `<!DOCTYPE html><html lang="en-US">
<head>
<title>Hello,World!</title>
</head>
<body>
<div class="container">
<header>
	<!-- Logo -->
   <h1>City Gallery</h1>
</header>  
<nav>
  <ul>
    <li><a href="/London">London</a></li>
    <li><a href="/Paris">Paris</a></li>
    <li><a href="/Tokyo">Tokyo</a></li>
  </ul>
</nav>
<article>
  <h1>London</h1>
  <img src="pic_mountain.jpg" alt="Mountain View" style="width:304px;height:228px;">
  <p>London is the capital city of England. It is the most populous city in the  United Kingdom, with a metropolitan area of over 13 million inhabitants.</p>
  <p>Standing on the River Thames, London has been a major settlement for two millennia, its history going back to its founding by the Romans, who named it Londinium.</p>
</article>
<footer>Copyright &copy; W3Schools.com</footer>
</div>
</body>
</html>`

	var tests = []struct {
		expr     string
		expected interface{}
	}{
		{`normalize-space(//div[@class="container"]/header)`, "City Gallery"},
		{`//li/a`, []string{"London", "Paris", "Tokyo"}},
		{`count(//h1)`, 2.0},
		{`boolean(count(//code))`, false},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			v, err := EvalXPathHTML(test.expr, []byte(html))
			assert.Equal(t, test.expected, v)
			assert.Nil(t, err)
		})
	}
}

func TestEvalXPathHTMLBug1(t *testing.T) {

	const html = `<html>
	<body>
		<div class="fruit">Apple</div>
		<div class="fruit">Banana</div>
		<div class="fruit">Lemon</div>
	</body>
</html>`

	doc, _ := htmlquery.Parse(strings.NewReader(html))
	test := `string((//div[@class="fruit"])[2])`
	expr, _ := xpath.Compile(test)
	v := expr.Evaluate(htmlquery.CreateXPathNavigator(doc))
	assert.Equal(t, v, "Banana")
}

func TestEvalXPathHTMLBug2(t *testing.T) {

	const html = `<html>
	<body>
			<div class="fruit">Apple</div>
			<div class="color">Red</div>
	</body>
</html>`

	doc, _ := htmlquery.Parse(strings.NewReader(html))
	test := `string((//div[@class="color"])[1])`
	expr, _ := xpath.Compile(test)
	v := expr.Evaluate(htmlquery.CreateXPathNavigator(doc))
	assert.Equal(t, v, "Red")
}

/*func TestEvalXPathHTMLBug3(t *testing.T) {

	const html = `<html>
		<body>
				<div class="fruit">
					<div class="color">Red</div>
				</div>
				<div class="fruit">
					<div class="color">Yellow</div>
				</div>
		</body>
	</html>`

	doc, _ := htmlquery.Parse(strings.NewReader(html))
	test := `count(((//div[@class="fruit"])[1])//div[@class="color"])`
	expr, _ := xpath.Compile(test)
	v := expr.Evaluate(htmlquery.CreateXPathNavigator(doc))
	// v = 2, expected v = 1
	assert.Equal(t, 1.0, v)
}*/

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestEvalXPathAll(t *testing.T) {

	fhtlml, _ := os.Open("testdata/sample1.html")
	defer fhtlml.Close()
	doc, _ := htmlquery.Parse(bufio.NewReader(fhtlml))
	nav := htmlquery.CreateXPathNavigator(doc)

	fcsv, _ := os.Open("testdata/sample1.csv")
	defer fcsv.Close()

	s := bufio.NewScanner(fcsv)
	for s.Scan() {
		col := strings.Split(s.Text(), "|")
		test := strings.TrimSpace(col[0])
		expected := strings.TrimSpace(col[1])
		expr, err := xpath.Compile(test)
		check(err)
		actual := expr.Evaluate(nav)

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