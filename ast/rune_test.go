package ast

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunesNotEqual(t *testing.T) {
	r1 := []rune("apple")
	r2 := []rune("banana")
	assert.False(t, equal(r1, r2))
}

func TestIsDigit(t *testing.T) {
	assert.True(t, isDigit('0'))
	assert.True(t, isDigit('9'))
	assert.False(t, isDigit('a'))
}
