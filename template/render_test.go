package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRenderOk(t *testing.T) {

	var tests = []struct {
		text      string
		variables map[string]string
		expected  string
	}{
		{
			"User id:{{id}} -> name:{{name}}, firstName:{{firstName}}",
			map[string]string{"id": "1234567", "name": "Bart", "firstName": "Simpson"},
			"User id:1234567 -> name:Bart, firstName:Simpson",
		},
		{
			"{{name}}-{{name}}-{{name}}",
			map[string]string{"name": "toto"},
			"toto-toto-toto",
		},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			result, err := Render(test.text, test.variables)
			assert.Nil(t, err)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRenderFailed(t *testing.T) {

	var tests = []struct {
		text      string
		variables map[string]string
	}{
		{
			"{{id}}:{{name}}",
			map[string]string{"id": "1234567"},
		},
		{
			"{{ name }}",
			map[string]string{"name": "toto"},
		},
		{
			"{{ Name }}",
			map[string]string{"name": "toto"},
		},
	}

	for _, test := range tests {
		t.Run(test.text, func(t *testing.T) {
			_, err := Render(test.text, test.variables)
			assert.NotNil(t, err)
		})
	}
}