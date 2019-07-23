package query

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEvalJSONPath(t *testing.T) {
	const json = `{ 
"store": {
    "book": [ 
      { "category": "reference",
        "author": "Nigel Rees",
        "title": "Sayings of the Century",
        "price": 8.95
      },
      { "category": "fiction",
        "author": "Evelyn Waugh",
        "title": "Sword of Honour",
        "price": 12.99
      },
      { "category": "fiction",
        "author": "Herman Melville",
        "title": "Moby Dick",
        "isbn": "0-553-21311-3",
        "price": 8.99
      },
      { "category": "fiction",
        "author": "J. R. R. Tolkien",
        "title": "The Lord of the Rings",
        "isbn": "0-395-19395-8",
        "price": 22.99
      }
    ],
    "bicycle": {
      "color": "red",
      "price": 19.95,
      "new": true
    }
  }
}`
	var tests = []struct {
		expr string
		expected interface{}
	}{
		{`$.store.book[0].title`, "Sayings of the Century"},
		{`$['store']['book'][0]['title']`, "Sayings of the Century"},
		{`$["store"]["book"][0]["title"]`, "Sayings of the Century"},
		{`$.store.book[3].price`, 22.99},
		{`$.store.bicycle.new`, true},
	}
	for _, test := range tests {
		t.Run(test.expr, func(t *testing.T) {
			v, err := EvalJSONPath(test.expr, []byte(json))
			assert.Equal(t, test.expected, v)
			assert.Nil(t, err)
		})
	}
}