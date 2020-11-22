package value

import (
	"github.com/lanvard/support"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_set_collection_on_empty_value(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("*", "water")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{"water"}, result.Raw())
}

func Test_set_deep_collection_on_empty_value(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("*.*", "water")
	assert.Nil(t, err)
	assert.Equal(t, []interface{}{[]interface{}{"water"}}, result.Raw())
}

func Test_set_collection_on_string_value(t *testing.T) {
	data := support.NewValue("rain")
	result, err := data.SetE("*", "water")
	assert.EqualError(t, err, "can not append value on 'string'")
	assert.Equal(t, support.NewValue("rain"), result)
}

func Test_set_on_empty_key(t *testing.T) {
	data := support.NewValue(nil)
	result, err := data.SetE("", "water")
	assert.Nil(t, err)
	assert.Equal(t, "water", result.Raw())
}
