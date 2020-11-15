package test

import (
	"github.com/lanvard/support/str"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_upper_first_with_empty_string(t *testing.T) {
	result := str.UpperFirst("")
	require.Equal(t, "", result)
}

func Test_upper_first_with_multiple_words(t *testing.T) {
	result := str.UpperFirst("a horse is happy")
	require.Equal(t, "A horse is happy", result)
}
