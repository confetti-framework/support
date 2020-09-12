package test

import (
	"github.com/lanvard/support/str"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpperFirstWithEmptyString(t *testing.T) {
	result := str.UpperFirst("")
	assert.Equal(t, "", result)
}

func TestUpperFirstWithMultipleWords(t *testing.T) {
	result := str.UpperFirst("a horse is happy")
	assert.Equal(t, "A horse is happy", result)
}
