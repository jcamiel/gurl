package gurl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseBasicString(t *testing.T) {
	input := `So multiline string
Multiline 1
Multiline 2
Multiline 3`
	parsed := `So multiline string
Multiline 1
Multiline 2
Multiline 3`
	assert.Equal(t, input, parsed, "The two multilines should be the same.")
}
