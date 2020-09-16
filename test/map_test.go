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

func TestMapOnlyWhenAllKeysArePresent(t *testing.T) {
	data := support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	assert.Equal(t, data, data.Only("username", "password"))
}

func TestMapOnlyWhenLessKeysThanPresent(t *testing.T) {
	assert.Equal(
		t,
		support.NewMapByString(map[string]string{"username": "apple_pear"}),
		support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Only("username"),
	)
}

func TestMapOnlyWhenMoreKeysThanPresent(t *testing.T) {
	data := support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	assert.Equal(t, data, data.Only("username", "password", "age"))
}

func TestMapExceptWhenNoKeysArePresent(t *testing.T) {
	data := support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	assert.Equal(t, data, data.Except())
}

func TestMapExceptWhenLessKeysThanPresent(t *testing.T) {
	assert.Equal(t,
		support.NewMapByString(map[string]string{"username": "apple_pear"}),
		support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Except("password"),
	)
}

func TestMapExceptWhenMoreKeysThanPresent(t *testing.T) {
	assert.Equal(t,
		support.NewMapByString(map[string]string{"username": "apple_pear"}),
		support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"}).Except("password", "age"),
	)
}

func TestNoReferenceToOldStruct(t *testing.T) {
	oldStruct := support.NewMapByString(map[string]string{"username": "apple_pear"})
	_ = oldStruct.Except("username")
	assert.Equal(t, oldStruct, support.NewMapByString(map[string]string{"username": "apple_pear"}))
}

func TestMapHasWithOneKey(t *testing.T) {
	data := support.NewMapByString(map[string]string{"username": "apple_pear"})
	assert.True(t, data.Has("username"))
	assert.False(t, data.Has("age"))
}

func TestMapHasWithMultipleKeys(t *testing.T) {
	data := support.NewMapByString(map[string]string{"username": "apple_pear", "password": "34a@#dQd"})
	assert.True(t, data.Has("username", "password"))
	assert.False(t, data.Has("username", "age"))
}
