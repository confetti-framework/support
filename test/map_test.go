package test

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_get_all_from_map(t *testing.T) {
	values := support.NewMapByUrlValues(map[string][]string{
		"language": {"Go"},
		"names":    {"David", "Jona"},
	})

	value := values.Get("*")
	assert.Len(
		t,
		value.Collection(),
		3,
	)
}

// func TestMapOnlyWhenAllKeysArePresent(t *testing.T) {
// 	data := support.NewMapByString(map[string]string{"username", "password"})
// 	assert.Equal(t, data, data.Only("username", "password"))
// }
