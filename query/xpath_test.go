package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const htmlSample = `<!DOCTYPE html><html lang="en-US">
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
</html>
`

func TestEvalXPathHTML(t *testing.T) {
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
			v, err := EvalXPathHTML(test.expr, []byte(htmlSample))
			assert.Equal(t, test.expected, v)
			assert.Nil(t, err)
		})
	}
}
